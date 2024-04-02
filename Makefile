GOFMT=find . -name *.go -exec gofmt -l -s -w {} \;
GOLINT=golangci-lint run
APP=main.go
BINARY=hex

.PHONY: all
all: vet fix fmt lint

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GOFMT)

.PHONY: lint
lint:
	@echo "go lint"
	@$(GOLINT)

.PHONY: build
build:
	go build -o $(BINARY) $(APP)

.PHONY: clean
clean:
	rm -rf $(BINARY)
