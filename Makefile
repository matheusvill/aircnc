version ?= latest
go-version ?= latest
IMAGE_NAME=aircnc
DEV_IMAGE_NAME=$(IMAGE_NAME):$(version)
SOURCE_DIR=/go/src/github.com/matheusvill/aircnc
RUN_GO=--rm --net="host" -e SEVERINO_LOGGER='debug' -v `pwd`:$(SOURCE_DIR) -w $(SOURCE_DIR) $(IMAGE_NAME)
RUN_GO_BUILD=--rm -v `pwd`:$(SOURCE_DIR) -w $(SOURCE_DIR) $(IMAGE_NAME)

dev-image:
	docker build -t $(IMAGE_NAME) -f ./hack/Dockerfile .

build: dev-image
	docker run $(RUN_GO_BUILD) ./hack/build.sh

image: build
	docker build -t $(IMAGE_NAME) .

image-dev: build
	docker build -t $(DEV_IMAGE_NAME) .

deps: dev-image
	docker run $(RUN_GO_BUILD) ./hack/deps.sh
	sudo chown -R $(USER):$(USER) ./vendor .git

run: build
	docker run $(RUN_GO) ./aircnc

start-compose:
	docker-compose up -d

stop-compose:
	docker-compose kill
	docker-compose rm -f

check: dev-image
	@docker run $(RUN_GO) ./hack/check.sh

check-integration: dev-image stop-compose start-compose
	@docker run $(RUN_GO) ./hack/check-integration.sh
	docker-compose kill
	docker-compose rm -f

stop:
	docker stop $(DEV_IMAGE_NAME)

shell: dev-image
	docker run -ti $(RUN_GO) bash

logs:
	docker logs -f $(DEV_IMAGE_NAME)

release: publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

check-cover: dev-image
	@docker run $(RUN_GO) ./hack/cover.sh
