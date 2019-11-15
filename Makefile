GOOS = linux
GOARC = amd64
VERSION ?= latest

.PHONY: dep
dep:
	dep ensure
	@echo "Dependicies downloaded!"

.PHONY: build
build:
	go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo \
	  -o dist/komutan main.go
	@echo "You can find executable under dist directory."

.PHONY: install
install: build
	@sudo mv dist/komutan /usr/local/bin/komutan
	@sudo chmod +x /usr/local/bin/komutan
	@echo "Installation finished!"
