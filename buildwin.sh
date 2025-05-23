env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o cardindex.exe .
go-winres simply --icon images/cardindex.ico
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc \
    CXX=x86_64-w64-mingw32-g++ \
    go build -buildvcs=false \
      -ldflags="-s -w -H=windowsgui -extldflags=-static" \
      -o cardindex.exe .
strip cardindex.exe
