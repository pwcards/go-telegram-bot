RELEASE_TAG = $(shell date -u +"%y%m%d_%H%M%S")

req-linter:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.31.0

lint: build-mocks
	golangci-lint run --timeout=3m

deploy:
	git checkout master
	git pull
	git tag -a "${RELEASE_TAG}" -m ""
	git push origin "${RELEASE_TAG}"
