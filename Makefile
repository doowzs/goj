.PHONY: all

BIN=goj

build:
	go build -o $(BIN)

install: build
	mkdir -p ${exec_prefix}
	cp $(BIN) ${exec_prefix}/bin