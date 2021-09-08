# go-grep
Grep clone using Golang

In the root of the project run the command `go install ./...`

Then use the program like mentioned below.

The clone supports the flags 

1. `-i` for case-insensitive search.
    - USAGE: `gogrep -i pattern destination`.

2. `-c` for the number of matches within the destination.
    - USAGE: `gogrep -c pattern destination`.

3. `-A [n]` for 'n' number of lines after a match.

    - USAGE: `gogrep -A 3 pattern destination`.

4. `-B [n]` for 'n' number of matches before a match.
    - USAGE: `gogrep -B 3 pattern destination`.

