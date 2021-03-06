GOOS = linux
GOARC = amd64
VERSION ?= latest

.PHONY: dep
dep:
	dep ensure
	@echo "Dependicies downloaded!"

.PHONY: build
build:
	@go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo \
	  -o dist/komutan main.go
	@echo "You can find executable under dist directory."

.PHONY: package
package: build
	@cd dist && tar -czf komutan.tar.gz komutan ../LICENSE ../README.md

.PHONY: install
install: build
	@sudo mv dist/komutan /usr/local/bin/komutan
	@sudo chmod +x /usr/local/bin/komutan
	@echo "Installation finished!"

.PHONY: distclean
distclean:
	@sudo rm -f dist/komutan /usr/local/bin/komutan
	@echo "Komutan deleted from executables."

.PHONY: upgrade
upgrade: distclean install
	@echo "Upgrade finished!"

.PHONY: test
test:
	@go test -cover -coverprofile /tmp/komutan.test ./... \
	| tee /tmp/komutan.covarage

.PHONY: uninstall
uninstall: distclean
	@echo "Komutan uninstalled successfully!"
