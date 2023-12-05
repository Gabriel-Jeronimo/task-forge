.PHONY: build

build:
	sam build

deploy:
	sam deploy

buildAndDeploy:
	make build && make deploy

delete:
	sam delete