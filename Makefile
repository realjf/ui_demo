
VERSION="0.0.1"
BIN="bin/gui_demo"
BIN_WIN=".exe"
BIN_MACOS=".dmg"
ARCH="amd64"

DARWIN_ARCH="arm64"


WINDOWS_GCC="/usr/bin/x86_64-w64-mingw32-gcc"
WINDOWS_GXX="/usr/bin/x86_64-w64-mingw32-g++"


DARWIN_GCC="o64-clang"
DARDOW_GXX="o64-clang++"

GO_PROJECTS="/home/realjf/go/src"


XCGO_IMAGE="realjf/xcgo-sdk11.3-catalina:go1.19"


.PHONY: 



build:
	@echo 'build linux...'
	@env CGO_ENABLED=1 GOOS=linux GOARCH=${ARCH} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-linux-${ARCH}-${VERSION} main.go
	@echo 'done'

build_win:
	@echo 'build windows...'
	@env CGO_ENABLED=1 GOOS=windows GOARCH=${ARCH} CC=${WINDOWS_GCC} CXX=${WINDOWS_GXX} go build -ldflags '-s -w -X main.Version=${VERSION} -H windowsgui' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-windows-${ARCH}-${VERSION}${BIN_WIN} main.go
	@echo 'done'

build_darwin:
	@echo 'build darwin amd64...'
	@env CGO_ENABLED=1 GOOS=darwin GOARCH=${ARCH} CC=${DARWIN_GCC} CXX=${DARDOW_GXX} go build -ldflags '-s -w -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-darwin-${ARCH}-${VERSION}${BIN_MACOS} main.go
	@echo 'done'

build_darwin_arm:
	@echo 'build darwin arm64...'
	@env CGO_ENABLED=1 GOOS=darwin GOARCH=${DARWIN_ARCH} CC=${DARWIN_GCC} CXX=${DARDOW_GXX} go build -ldflags '-s -w -v -X main.Version=${VERSION}' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN}-darwin-${DARWIN_ARCH}-${VERSION}${BIN_MACOS} main.go
	@echo 'done'

launch_docker:
	@echo 'launch docker...'
	@docker run  -v ${GO_PROJECTS}:/go/src -v /var/run/docker.sock:/var/run/docker.sock -w /go/src -it ${XCGO_IMAGE} zsh 


setup:
	@echo 'install library...'
	@chmod +x ./scripts/setup.py
	@./scripts/setup.py -f ./scripts/libraries.json -d -b

push:
	@echo 'git push...'
	@git add -A && git commit -m "update" && git push origin master
	@echo 'done'

# Purging All Unused or Dangling Images, Containers, Volumes, and Networks
clean_docker_image:
	@echo 'clean docker image locally...'
	@docker system prune -a



clean:
	@rm -rf bin
	@rm -rf dist
	@echo 'done'