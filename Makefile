dev:
	go get ./...

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/dbmigrater-linux *.go
	env GOOS=darwin go build -ldflags="-s -w" -o bin/dbmigrater-darwin *.go
