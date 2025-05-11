PROJECT_NAME=webhook

tidy:
	go mod tidy

VERSION=`cat VERSION`

tag:
	git tag -a ${VERSION}
	git push --tags

MAIN_FILE_PATH=cmd/main.go
CONFIG_FILE_PATH=examples/config.yml

run:
	CONFIG_FILE_PATH=${CONFIG_FILE_PATH} go run $(MAIN_FILE_PATH)

BINARIES_DIRECTORY=bin
LDFLAGS="-w -s -X main.version=${VERSION}"

clean:
	rm -rf ${BINARIES_DIRECTORY}

build: clean
	go build -ldflags=${LDFLAGS} -o ${BINARIES_DIRECTORY}/${PROJECT_NAME} ${MAIN_FILE_PATH}

build-docker: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o ${BINARIES_DIRECTORY}/${PROJECT_NAME}_docker ${MAIN_FILE_PATH}


CONTAINER_RUNNER=docker
COMPOSE_FILE_PATH=docker-compose.d/docker-compose.yaml

compose-up:
	$(CONTAINER_RUNNER) compose -f $(COMPOSE_FILE_PATH) up

LOCAL_BIN := $(CURDIR)/bin
GOOSE_VERSION := v3.24.2
LINT_VERSION := 2.1.6

.prep_bin:
	mkdir -p ${LOCAL_BIN}

.install-lint:
	curl -Ls https://github.com/golangci/golangci-lint/releases/download/v${LINT_VERSION}/golangci-lint-${LINT_VERSION}-linux-amd64.tar.gz | tar xvz --strip-components=1 -C ${LOCAL_BIN} golangci-lint-${LINT_VERSION}-linux-amd64/golangci-lint

install-deps: \
	.prep_bin \
	.install-lint

lint: $(LINT_BIN)
	$(LOCAL_BIN)/golangci-lint run
