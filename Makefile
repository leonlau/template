all: clean build

SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "template/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags = "-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"


build:
	@echo "build"
	@go build  -v -ldflags ${ldflags} -mod=vendor .

tidy:
	@echo "tidy"
	@go mod tidy
	@go mod vendor

clean:
	@echo "clean"
	@rm -f warder
