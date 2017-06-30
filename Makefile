build:
	go build -o bin/wsc -ldflags="-s -w" wsc.go

compress:
	upx --brute bin/wsc
