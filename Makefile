TOOLS ?= autonomy/tools:96cb1b9

KUBECTL_VERSION ?= v1.16.0

ifeq ($(UNAME_S),Linux)
KUBECTL_ARCHIVE := https://storage.googleapis.com/kubernetes-release/release/$(KUBECTL_VERSION)/bin/linux/amd64/kubectl
endif
ifeq ($(UNAME_S),Darwin)
KUBECTL_ARCHIVE := https://storage.googleapis.com/kubernetes-release/release/$(KUBECTL_VERSION)/bin/darwin/amd64/kubectl
endif

ifeq ($(UNAME_S),Linux)
GITMETA := https://github.com/talos-systems/gitmeta/releases/download/v0.1.0-alpha.2/gitmeta-linux-amd64
endif
ifeq ($(UNAME_S),Darwin)
GITMETA := https://github.com/talos-systems/gitmeta/releases/download/v0.1.0-alpha.2/gitmeta-darwin-amd64
endif

BINDIR ?= ./bin

SHA ?= $(shell $(BINDIR)/gitmeta git sha)
TAG ?= $(shell $(BINDIR)/gitmeta image tag)
BRANCH ?= $(shell $(BINDIR)/gitmeta git branch)

DOCKER_BUILD = docker buildx build

COMMON_ARGS = --progress=plain
# COMMON_ARGS += --allow security.insecure
COMMON_ARGS += --build-arg=TOOLS=$(TOOLS)
COMMON_ARGS += --build-arg=SHA=$(SHA)
COMMON_ARGS += --build-arg=TAG=$(TAG)
COMMON_ARGS += --build-arg=GO_VERSION=$(GO_VERSION)
COMMON_ARGS += .

DOCKER_ARGS ?=

GO_VERSION ?= 1.13
TESTPKGS ?= ./...

all: deps rootfs initramfs kernel osctl-linux osctl-darwin installer container

.PHONY: deps
deps: gitmeta

gitmeta: $(BINDIR)/gitmeta

$(BINDIR)/gitmeta:
	@mkdir -p $(BINDIR)
	@curl -L $(GITMETA) -o $(BINDIR)/gitmeta
	@chmod +x $(BINDIR)/gitmeta

kubectl: $(BINDIR)/kubectl

$(BINDIR)/kubectl:
	@mkdir -p $(BINDIR)
	@curl -L -o $(BINDIR)/kubectl $(KUBECTL_ARCHIVE)
	@chmod +x $(BINDIR)/kubectl

.PHONY: generate
generate:
	@$(DOCKER_BUILD) \
		--output type=local,dest=./ \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: kernel
kernel:
	@$(DOCKER_BUILD) \
		--output type=local,dest=build \
		--target=$@ \
		$(COMMON_ARGS)
	@-rm -rf ./build/modules

.PHONY: initramfs
initramfs:
	@$(DOCKER_BUILD) \
		--output type=local,dest=build \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: rootfs
rootfs: osd trustd proxyd ntpd networkd
	@$(DOCKER_BUILD) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: installer
installer:
	@mkdir -p build
	@$(DOCKER_BUILD) \
		--output type=docker,dest=build/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)
	@docker load < build/$@.tar

.PHONY: image-aws
image-aws:
	@docker run --rm -v /dev:/dev -v $(PWD)/build:/out \
		--privileged $(DOCKER_ARGS) \
		autonomy/installer:$(TAG) \
		install \
		-n aws \
		-r \
		-p aws \
		-u none \
		-e console=tty1 \
		-e console=ttyS0
	@tar -C $(PWD)/build -czf $(PWD)/build/aws.tar.gz aws.raw
	@rm -rf $(PWD)/build/aws.raw

.PHONY: image-azure
image-azure:
	@docker run --rm -v /dev:/dev -v $(PWD)/build:/out \
		--privileged $(DOCKER_ARGS) \
		autonomy/installer:$(TAG) \
		install \
		-n azure \
		-r \
		-p azure \
		-u none \
		-e console=ttyS0,115200n8 \
		-e earlyprintk=ttyS0,115200 \
		-e rootdelay=300
	@docker run --rm -v $(PWD)/build:/out $(DOCKER_ARGS) \
		--entrypoint qemu-img \
		autonomy/installer:$(TAG) \
		convert \
		-f raw \
		-o subformat=fixed,force_size \
		-O vpc /out/azure.raw /out/azure.vhd
	@tar -C $(PWD)/build -czf $(PWD)/build/azure.tar.gz azure.vhd
	@rm -rf $(PWD)/build/azure.raw $(PWD)/build/azure.vhd

.PHONY: image-gcp
image-gcp:
	@docker run --rm -v /dev:/dev -v $(PWD)/build:/out \
		--privileged $(DOCKER_ARGS) \
		autonomy/installer:$(TAG) \
		install \
		-n disk \
		-r \
		-p gcp \
		-u none \
		-e console=ttyS0
	@tar -C $(PWD)/build -czf $(PWD)/build/gcp.tar.gz disk.raw
	@rm -rf $(PWD)/build/disk.raw

.PHONY: push-image-aws
push-image-aws:
	@TAG=$(TAG) ./hack/test/aws-setup.sh

.PHONY: push-image-azure
push-image-azure:
	@TAG=$(TAG) ./hack/test/azure-setup.sh

.PHONY: push-image-gcp
push-image-gcp:
	@TAG=$(TAG) ./hack/test/gcp-setup.sh

.PHONY: image-test
image-test:
	@docker run --rm -v /dev:/dev -v /tmp:/out --privileged $(DOCKER_ARGS) autonomy/installer:$(TAG) install -n test -r -p test -u none

.PHONY: iso
iso:
	@docker run --rm -i -v $(PWD)/build:/out autonomy/installer:$(TAG) iso

.PHONY: container
container:
	@$(DOCKER_BUILD) \
		--output type=docker,dest=build/$@.tar,name=docker.io/autonomy/talos:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)
	@docker load < build/$@.tar

.PHONY: basic-integration
basic-integration:
	@TAG=$(TAG) ./hack/test/$@.sh

.PHONY: capi
capi:
	@TAG=$(TAG) ./hack/test/$@.sh

.PHONY: e2e-integration
e2e-integration:
	@TAG=$(TAG) ./hack/test/$@.sh

.PHONY: unit-tests
unit-tests:
	@$(BINDIR)/buildctl --addr $(BUILDKIT_HOST) \
		build \
		--opt target=$@ \
		--output type=local,dest=./ \
		--opt build-arg:TESTPKGS=$(TESTPKGS) \
		$(COMMON_ARGS)

.PHONY: unit-tests-race
unit-tests-race:
	@$(BINDIR)/buildctl --addr $(BUILDKIT_HOST) \
		build \
		--opt target=$@ \
		--opt build-arg:TESTPKGS=$(TESTPKGS) \
		$(COMMON_ARGS)

.PHONY: fmt
fmt:
	@docker run --rm -it -v $(PWD):/src -w /src golang:$(GO_VERSION) bash -c "export GO111MODULE=on; export GOPROXY=https://proxy.golang.org; cd /tmp && go mod init tmp && go get mvdan.cc/gofumpt/gofumports && cd - && gofumports -w -local github.com/talos-systems/talos ."

.PHONY: lint
lint:
	@$(BINDIR)/buildctl --addr $(BUILDKIT_HOST) \
		build \
		--opt target=$@ \
		$(COMMON_ARGS)

.PHONY: protolint
protolint:
	@$(BINDIR)/buildctl --addr $(BUILDKIT_HOST) \
		build \
		--opt target=$@ \
		$(COMMON_ARGS)

.PHONY: markdownlint
markdownlint:
	@$(BINDIR)/buildctl --addr $(BUILDKIT_HOST) \
		build \
		--opt target=$@ \
		$(COMMON_ARGS)

.PHONY: osctl-linux
osctl-linux:
	@$(DOCKER_BUILD) \
		--output type=local,dest=build \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: osctl-darwin
osctl-darwin:
	@$(DOCKER_BUILD) \
		--output type=local,dest=build \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: machined
machined: images
	@$(DOCKER_BUILD) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: osd
osd: images
	@$(DOCKER_BUILD) \
		--output type=docker,dest=images/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: trustd
trustd: images
	@$(DOCKER_BUILD) \
		--output type=docker,dest=images/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: proxyd
proxyd: images
	@$(DOCKER_BUILD) \
		--output type=docker,dest=images/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: ntpd
ntpd: images
	@$(DOCKER_BUILD) \
		--output type=docker,dest=images/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

.PHONY: networkd
networkd: images
	@$(DOCKER_BUILD) \
		--output type=docker,dest=images/$@.tar,name=docker.io/autonomy/$@:$(TAG) \
		--target=$@ \
		$(COMMON_ARGS)

images:
	@mkdir -p images

.PHONY: login
login:
	@docker login --username "$(DOCKER_USERNAME)" --password "$(DOCKER_PASSWORD)"

.PHONY: push
push: gitmeta login
	@docker push autonomy/installer:$(TAG)
	@docker push autonomy/talos:$(TAG)
ifeq ($(BRANCH),master)
	@docker tag autonomy/installer:$(TAG) autonomy/installer:latest
	@docker tag autonomy/talos:$(TAG) autonomy/talos:latest
	@docker push autonomy/installer:latest
	@docker push autonomy/talos:latest
endif

.PHONY: clean
clean:
	@-rm -rf build images vendor
