dev:
	go get ./...

test:
	go run ./... create
	go run ./... migrate up
	go run ./... migrate down
	go run ./... drop

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/dbmigrater-linux *.go
	env GOOS=darwin go build -ldflags="-s -w" -o bin/dbmigrater-darwin *.go
