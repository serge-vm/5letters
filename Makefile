SWAG_VERSION := $(shell swag -v 2>/dev/null)

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
ifdef SWAG_VERSION
	@echo "Found $(SWAG_VERSION)"
	swag init -g ./api/main.go
else
	@echo "Error: swag not found. Documentation was not regenerated."
endif

clean:
	rm -f bin/*