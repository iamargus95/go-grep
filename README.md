# go-grep
Grep clone using Golang

The clone supports the flags 

1. `-i` for case-insensitive search.

2. `-c` for the number of matches within the destination.

3. `-A [n]` for 'n' number of lines after a match.

    - Usage: `gogrep -A 3 pattern destination`

4. `-B [n]` for 'n' number of matches before a match.
    - Usage: `gogrep -A 3 pattern destination`

