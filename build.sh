go build -ldflags "-s -w" -o cardindex .
strip cardindex
upx -q --best --lzma cardindex
