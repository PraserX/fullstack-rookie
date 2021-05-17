# Advanced 01

## Compile

```bash
GOOS=darwin  GOARCH=386   go build -v -o build/mybin_darwin_386        cmd/main.go
GOOS=darwin  GOARCH=amd64 go build -v -o build/mybin_darwin_amd64      cmd/main.go
GOOS=linux   GOARCH=386   go build -v -o build/mybin_linux_386         cmd/main.go
GOOS=linux   GOARCH=amd64 go build -v -o build/mybin_linux_amd64       cmd/main.go
GOOS=windows GOARCH=386   go build -v -o build/mybin_windows_386.exe   cmd/main.go
GOOS=windows GOARCH=amd64 go build -v -o build/mybin_windows_amd64.exe cmd/main.go
```

## Testing

```bash
go test ./...
go test -bench=.
```