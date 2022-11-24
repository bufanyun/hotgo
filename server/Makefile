ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

build:
	@rm -rf ./resource/public/admin/*
	@cd ../web && yarn build && \cp -rf ./dist/*  ../server/resource/public/admin/
	@echo "y" | gf build

push:
	@cd $(ROOT_DIR) && cd .. && ./push.sh

all:
	gf run main.go --args "all"

web:
	@cd ../web && yarn serve

# 清理gf调试进程
killmain:
	@kill -9 $(ps -ef|grep main|grep -v grep|awk '{print $2}')

# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf


# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	fi;


# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@gf gen service

# Build image, deploy image and yaml to current kubectl environment and make port forward to local machine.
.PHONY: start
start:
	@set -e; \
	make image; \
	make deploy; \
	make port;

# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git log -1 --format="%cd.%h" --date=format:"%Y%m%d%H%M%S"))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@gf docker -p -b "-a amd64 -s linux -p temp" -t $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push:
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_TAG};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}";


