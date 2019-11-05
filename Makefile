.PHONY: all

BIN=goj

build:
	go build -o $(BIN)

install: build
	cp $(BIN) ${exec_prefix}/bin