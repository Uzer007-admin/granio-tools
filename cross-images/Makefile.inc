# VERSION is the Fyne Cross Images version
VERSION := $(shell git tag --points-at HEAD | awk -v def="dev" '{print} END { if(NR==0) {print def} }')
# REPOSITORY is the docker repository
REPOSITORY ?= ghcr.io/uzer007-admin/granio-cross

version:
	@echo $(VERSION)
