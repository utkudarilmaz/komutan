VERSION ?= latest

.PHONY: dep
dep:
	dep ensure
	@echo "Dependicies downloaded!"

.PHONY: install
install: dep
	env GOOS=linux GOARC=amd64 \
	  go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo \
	  -o dist/komutan main.go
	@echo "Installation finished! You can find executable under dist directory."
