all: clean docs console api

.PHONY: console
console:
	go build -o bin/5lettersc console/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/5lettersc.exe console/main.go

.PHONY: api
api:
	go build -o bin/5letters api/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/5letters.exe api/main.go

.PHONY: docs
docs:
	swag init -g ./api/main.go

clean:
	rm -f bin/*