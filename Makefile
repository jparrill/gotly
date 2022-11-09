.ONESHELL:
.PHONY: run redis tests doc clean 
.EXPORT_ALL_VARIABLES:

run:
	go run main.go

redis:
	docker run --name gotlydb -p 6379:6379 -d redis redis-server --save 60 1 --loglevel warning

tests:
	go test -v -cover ./...

build: bin/gotly
	go build ./...

coverage: c.out
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out

doc:
	godoc


clean:
	go clean
