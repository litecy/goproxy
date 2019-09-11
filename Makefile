SHELL=/bin/bash

env_pwd=$(shell pwd)

UNAME_S := $(shell uname -s)

sedprefix := sed -i -e
ifeq ($(UNAME_S),Darwin)
	sedprefix := sed -i ''
endif

md5sum := md5sum
ifeq ($(UNAME_S),Darwin)
	md5sum := md5
endif

# only change to app_name to your project main execu
app_name=goproxy

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# for statically link
CGO_ENABLED?=0

output_bin=bin/$(GOOS)_$(GOARCH)/$(app_name)

default: all

all: ${app_name} package

build: ${app_name}

check-hash:
	@$(md5sum) ${output_bin}

run-dev:
	${output_bin}

goproxy:
	go build -o ${output_bin} cmd/$(subst -,_,$(app_name))/main.go

package:
	cp ${output_bin} ./${app_name}
	tar -czf bin/$(GOOS)_$(GOARCH)/${app_name}.tgz ${app_name}
	rm -rf ./${app_name}

clean:
	rm ${output_bin} test/${app_name}_test

debug:
	echo ${app_name} ${output_bin}
	echo go build -o ${output_bin} cmd/$(subst -,_,$(app_name))/main.go