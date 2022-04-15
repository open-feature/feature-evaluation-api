# docker info
SWAGGER_VERSION ?= v0.29.0
IS_SWAGGER_INSTALLED = $(shell which swagger >> /dev/null 2>&1; echo $$?)
IS_DOCKER_INSTALLED = $(shell which docker >> /dev/null 2>&1; echo $$?)
IS_TRIVY_INSTALLED = $(shell which trivy >> /dev/null 2>&1; echo $$?)
CONTAINER_REGISTRY?=ghcr.io
CONTAINER_REGISTRY_ORG?=open-feature
CONTAINER_IMAGE?=openfeature-evaluation-api
VERSION?=dev

.PHONY: all
all: api test build

.PHONY: clean
clean:
	@echo ""
	@echo "Removing Swagger generated folders"
	@rm -rf models

.PHONY: swaggerdeps
swaggerdeps:
	@echo ""
	@echo "Installing dependencies for Swagger"
	@go get golang.org/x/tools/cmd/goimports
	@go install github.com/go-swagger/go-swagger/cmd/swagger@$(SWAGGER_VERSION)

.PHONY: builddeps
builddeps:
	@echo ""
	@echo "Installing Build dependencies"
	@go get ./...

.PHONY: api
api: clean swaggerdeps _build_check_swagger
	@echo ""
	@echo "Building API"
	@swagger generate server -f api/swagger.yaml --default-scheme http
	@go get -u -f ./...

.PHONY: test
test: unused-package-check gofmt-check go-test

.PHONY: go-test
go-test:
	@echo "------------------"
	@echo "--> Run Go Test"
	@echo "------------------"
	@go test ./... -coverprofile=coverage.txt -covermode=atomic -v

.PHONY: unused-package-check
unused-package-check:
	@echo "------------------"
	@echo "--> Check unused packages for the feature-evaluation-api"
	@echo "------------------"

	@tidy=$$(go mod tidy); \
	if [ -n "$${tidy}" ]; then \
		echo "go mod tidy checking failed!"; echo "$${tidy}"; echo; \
	fi

.PHONY: gofmt-check
gofmt-check:
	@echo "------------------"
	@echo "--> Check unused packages for the feature-evaluation-api"
	@echo "------------------"
	@gfmt=$$(gofmt -s -l . | wc -l); \
	if [ "$${gfmt}" -ne 0 ]; then \
		echo "The following files were found to be not go formatted:"; \
   		gofmt -s -l .; \
   		exit 1; \
  	fi

.PHONY: build
build: api test
	@echo ""
	@echo "Building Go Binary"
	@mkdir -p bin
	@go build -o bin/feature-evaluation-api -ldflags="-extldflags=-static -X github.com/open-feature/feature-evaluation-api/handlers.Version=${VERSION}" cmd/openfeature-evaluation-api-server/main.go

_build_check_swagger:
	@if [ $(IS_SWAGGER_INSTALLED) -eq 1 ]; \
		then echo "" \
		&& echo "ERROR:\tswagger is not installed. Please install it before build." \
		&& echo "" \
		&& exit 1; \
		fi;

_build_check_docker:
	@if [ $(IS_DOCKER_INSTALLED) -eq 1 ]; \
		then echo "" \
		&& echo "ERROR:\tdocker is not installed. Please install it before build." \
		&& echo "" \
		&& exit 1; \
		fi;


container: container-build container-push
container-build: _build_check_docker
	@echo ""
	@echo "Building Container"
	@docker build -f build/docker/Dockerfile -t ${CONTAINER_REGISTRY}/${CONTAINER_REGISTRY_ORG}/${CONTAINER_IMAGE}:${VERSION} .

container-push: _build_check_docker
	@echo ""
	@echo "Pushing Container"
	@docker push ${CONTAINER_REGISTRY}/${CONTAINER_REGISTRY_ORG}/${CONTAINER_IMAGE}:${VERSION}

run: build
	@./bin/feature-evaluation-api --port 8080

run-container: container-build
	@docker run -p 8080:8080 ${CONTAINER_REGISTRY}/${CONTAINER_REGISTRY_ORG}/${CONTAINER_IMAGE}:${VERSION} .
