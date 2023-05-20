BUILDDIR:=./bin
WEBBUILDDIR:=./web/bin
GOROOT:=$(shell go env GOROOT)
BIN:=typerace

.PHONY: run
run:
	go run .

.PHONY: build
build: clean
	mkdir -p ${BUILDDIR}
	go build \
		-o ${BUILDDIR}/${BIN} \
		.

.PHONY: build/wasm
build/wasm: clean
	mkdir -p ${BUILDDIR}
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ${BUILDDIR}
	GOOS=js GOARCH=wasm go build \
		-o ${BUILDDIR}/${BIN}.wasm \
		.

.PHONY: clean
clean:
	rm -rf *.wasm
	rm -rf ${BUILDDIR}
	rm -rf ${WEBBUILDDIR}

.PHONY: docker/build
docker/build: build/wasm
	mkdir -p ${WEBBUILDDIR}
	cp ${BUILDDIR}/* ${WEBBUILDDIR}
	docker build -f web/Dockerfile.nginx -t typerace-nginx:latest web

.PHONY: docker/run
docker/run:
	docker run --rm --name typerace -p 8080:80 typerace-nginx:latest
