# import dev config
# You can change the default dev config with `make cnf="deploy_special.env" release`
dpl ?= dev.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

ifdef c
CLASSFOLDER:=classroom$(c)
else
CLASSFOLDER:=classroom01
endif
ifdef s
SAMPLENAME:=sample$(s)
SAMPLEFILE:=$(SAMPLENAME).go
endif
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build sample based in classroom ex.: make build c=01 s=001
ifdef s
	go build -o ./bin/$(CLASSFOLDER)-$(SAMPLENAME) ./src/$(CLASSFOLDER)/$(SAMPLEFILE)
else
	@echo "Please inform a sample file to build. see make help to see an example"
endif

run: ## Run sample based in classroom ex.: make run c=01 s=001
ifdef s
	go run ./src/$(CLASSFOLDER)/$(SAMPLEFILE)
else
	@echo "Please inform a sample file to run. see make help to see an example"
endif