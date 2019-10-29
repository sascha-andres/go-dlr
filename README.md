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

## History

|Version|Description|
|---|---|
||Initial version|