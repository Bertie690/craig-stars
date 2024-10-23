BINARY_NAME:=craig-stars
VERSION:=0.0.0-develop
COMMIT:=`git rev-parse HEAD`
BUILDTIME:= $$(date +'%y.%m.%d %H:%M:%S')

# detect os type and swap instructions accordingly
# Swap to powershell if on windows 
ifeq ($(OS),Windows_NT) 
detected_OS := Windows
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
else
detected_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

# conditionals needed to mimic behavior on unix-like systems
ifeq ($(detected_OS),Windows)
mkdir = if (!(Test-Path $(1) )) {New-Item -Name $(1) -ItemType Directory}
rm = if ( Test-Path $(1) ) { Remove-Item -Recurse -Force $(1) }
else
mkdir = mkdir -p $(1)
rm = rm -rf $(1)
endif 

# replaces backslashes with unix-style frontslashes to make path universally valid
goroot := $(subst \, /,$(shell go env GOROOT)) 

# always redo these
.PHONY: build test clean dev dev_backend dev_frontend

build: build_frontend tidy vendor generate build_server

build_frontend:
	cd frontend; npm install
	cd frontend; npm run build

build_server:
	$(call mkdir,dist)
	go build \
	-o dist/${BINARY_NAME} \
	-ldflags \
	"-X 'github.com/sirgwain/craig-stars/cmd.semver=${VERSION}' \
	-X 'github.com/sirgwain/craig-stars/cmd.commit=${COMMIT}' \
	-X 'github.com/sirgwain/craig-stars/cmd.buildTime=${BUILDTIME}'" \
	main.go

build_wasm:
	$(call mkdir,frontend/src/lib/wasm)
	go env -w GOOS=js GOARCH=wasm
	go build \
	-o frontend/src/lib/wasm/cs.wasm \
	wasm/main.go
	cp $(goroot)/misc/wasm/wasm_exec.js ./frontend/src/lib/wasm/wasm_exec.js

# use docker to build an amd64 image for linux deployment
build_docker:
	docker build -f builder.Dockerfile --platform linux/amd64 . -t craig-stars-builder
	docker run -f builder.Dockerfile --platform linux/amd64 -v ${CURDIR}/dist:/dist craig-stars-builder

generate:
	go generate ./...

test:
	go test ./...
	cd frontend; npm run test

clean:
	go clean
	$(call rm,dist)
	$(call rm,vendor)
	$(call rm,frontend/build)

# uninstall unused modules
tidy:
	go mod tidy -v

# get those deps local!
vendor:
	go mod vendor

dev_frontend:
	cd frontend; npm run dev

dev_backend:
	air

dev:
	make -j 2 dev_backend dev_frontend

