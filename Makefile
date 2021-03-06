RELEASE_TAG = $(shell date -u +"%y%m%d_%H%M%S")

rebuild_app:
	docker-compose up -d --no-deps --build app

rebuild_database:
	docker-compose up -d --no-deps --build database

# Restart all project
restart:
	docker-compose stop $(c)
	docker-compose up -d $(c)

# Show log app
log_app:
	docker-compose logs -f -t app

# Show container list
ps:
	docker-compose ps

req-linter:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.31.0

lint:
	golangci-lint run --timeout=3m

deploy:
	git checkout master
	git pull
	git tag -a "${RELEASE_TAG}" -m ""
	git push origin "${RELEASE_TAG}"
