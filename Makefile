ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "universe-single"
DOCKER_NAME = "universe-single"
TIMESTAMP=$(shell date +%Y%m%d%H%M%S)
include .env
export
include ./hack/hack.mk

# This target is used to catch all arguments
%:
	@: