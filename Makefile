version ?= latest
go-version ?= latest
IMAGE_NAME=aircnc
PROD_IMAGE_NAME=registry-aws.neoway.com.br/datascience/$(IMAGE_NAME):$(version)
DEV_IMAGE_NAME=registry-aws.neoway.com.br/datascience/$(IMAGE_NAME):$(version)
SOURCE_DIR=/go/src/gitlab.neoway.com.br/jaws/aircnc
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

image-aws: build
	docker build -t $(PROD_IMAGE_NAME) .

deps: dev-image
	docker run $(RUN_GO_BUILD) ./hack/deps.sh
	sudo chown -R $(USER):$(USER) ./vendor .git

run: build
	docker run $(RUN_GO) ./aircnc

start-compose:
	docker build -t repository -f ./tests/repository/Dockerfile .
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

push: image-dev
	docker push $(DEV_IMAGE_NAME)

push-aws: image-aws
	docker push $(PROD_IMAGE_NAME)

logs:
	docker logs -f $(DEV_IMAGE_NAME)

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	        echo "Variable '$*' not set"; \
	        exit 1; \
	fi

release: publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: publish-dev publish-prod

publish-prod: image-aws
	docker tag -f $(DEV_IMAGE_NAME) $(PROD_IMAGE_NAME)
	docker push $(PROD_IMAGE_NAME)

publish-dev: image-aws
	docker push $(DEV_IMAGE_NAME)

scale: guard-env guard-count
	harbor $(env) scale aircnc $(count)

deploy-files:
	cat ./deploy/$(env)/aircnc-rc.yaml.tpl | sed "s/{{version}}/$(version)/g" > ./deploy/$(env)/aircnc-rc.yaml

update: guard-env deploy-files
	harbor $(env) delete rc aircnc
	harbor $(env) create ./deploy/aircnc-rc.yaml
	rm ./deploy/aircnc-rc.yaml

deploy: guard-env deploy-files
	harbor $(env) create ./deploy/$(env)/aircnc-rc.yaml
	rm ./deploy/$(env)/aircnc-rc.yaml
	harbor $(env) create ./deploy/$(env)/aircnc-svc.yaml

check-cover: dev-image
	@docker run $(RUN_GO) ./hack/cover.sh
