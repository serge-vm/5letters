SWAG_VERSION := $(shell swag -v 2>/dev/null)

all: clean docs api

.PHONY: api
api:
	go build -o bin/5letters cmd/main.go

.PHONY: docs
docs:	
ifdef SWAG_VERSION
	@echo "Found $(SWAG_VERSION)"
	swag init -g ./cmd/main.go
else
	@echo "Error: swag not found. Documentation was not regenerated."
endif

clean:
	rm -f bin/*