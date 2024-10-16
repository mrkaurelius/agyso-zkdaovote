all:
	go build -o build/zk cmd/main.go

web:
	GOOS=js GOARCH=wasm go build -o build/zk.wasm cmd/main.go