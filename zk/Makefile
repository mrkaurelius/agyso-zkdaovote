backend:
	go build -o build/daovote-server cmd/daovote-service/main.go

proof:
	go test -v ./gnark

webasm:
	GOOS=js GOARCH=wasm go build -o build/daovote.wasm cmd/daovote-wasm/main.go
