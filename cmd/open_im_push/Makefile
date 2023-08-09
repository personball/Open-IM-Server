.PHONY: all build run gotool install clean help image-push

NAME=open_im_push
BIN_DIR=../../bin/

OS:= $(or $(os),linux)
ARCH:=$(or $(arch),amd64)
all: gotool build

ifeq ($(OS),windows)

BINARY_NAME=${NAME}.exe

else

BINARY_NAME=${NAME}

endif

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

install:build
	mv ${BINARY_NAME} ${BIN_DIR}

clean:
	@if [ -f ${BINARY_NAME} ] ; then rm ${BINARY_NAME} ; fi

# 打包 open_im_push
image-push:
	docker build -t registry.cn-hangzhou.aliyuncs.com/vzero/docker.io.openim.push:v2.3.9-beta13 -f ./push.Dockerfile .

