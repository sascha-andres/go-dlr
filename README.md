# go-dlr

    Detect
    Local
    Replacements

This is a simple tool that detects local replacements in a go.mod file. It was created
to ensure not to push a local replacement to a remote with CI in place which (and thus
failing a build).

## Usage

    go-dlr [-f path-to-go-mod]

With -f defaulting to go.mod

Within [git-commit](https://github.com/sascha-andres/git-commit) you can integrate the tool like this:

    external-tools:
      - name: detect local replacements
        command:
          - go-dlr
        severity: error

This calls go-dlr and would prohibit committing unless all replacements are eliminated.

## History

|Version|Description|
|---|---|
|0.1.0|Initial version|
