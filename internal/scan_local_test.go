package internal

import (
	"golang.org/x/mod/modfile"
	"testing"
)

func TestInvalid(t *testing.T) {
	var withLocal = `module livingit.de/code/go-dlr/demofiles-linux

replace github.com/sascha-andres/doesnotexist => ../test
replace github.com.sascha-andres/doest-also-not-exist => /absolute/path`

	m, err := modfile.Parse("go.mod", []byte(withLocal), nil)
	if err != nil {
		t.Fatalf("error parsing go.mod: %s", err)
	}
	result  := scanForLocals(m)
	if !result {
		t.Logf("no locals detected though it shoud have")
		t.FailNow()
	}
}

func TestNoLocals(t *testing.T) {
	var withLocal = `module livingit.de/code/go-dlr/demofiles-linux`

	m, err := modfile.Parse("go.mod", []byte(withLocal), nil)
	if err != nil {
		t.Fatalf("error parsing go.mod: %s", err)
	}
	result  := scanForLocals(m)
	if result {
		t.Logf("no locals detected though it shoud have")
		t.FailNow()
	}
}