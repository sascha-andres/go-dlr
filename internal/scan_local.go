package internal

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io/ioutil"
)

// HasLocals returns true if the go.mod contains local redirects
func HasLocals(fileToParse string) (bool, error) {
	data, err := ioutil.ReadFile(fileToParse)
	if err != nil {
		return false, err
	}

	m, err := modfile.Parse(fileToParse, data, nil)
	if err != nil {
		return false, err
	}

	if scanForLocals(m) {
		return true, nil
	}
	return false, nil
}


func scanForLocals(m *modfile.File) bool {
	foundLocal := false
	for _, r := range m.Replace {
		if modfile.IsDirectoryPath(r.New.Path) {
			fmt.Printf("%s@%s is replaced to a local path\n", r.Old.Path, r.Old.Version)
			foundLocal = true
		}
	}
	return foundLocal
}