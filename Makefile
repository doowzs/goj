.PHONY: all

BIN=goj
FAKE_ROOT=${DESTDIR}
INSTALL_DIR=${FAKE_ROOT}/usr/bin

build:
	go build -o $(BIN)

install: build
	mkdir -p ${INSTALL_DIR}
	cp $(BIN) ${INSTALL_DIR}