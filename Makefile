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
LDFLAGS="-w -s"

clean:
	rm -rf ${BINARIES_DIRECTORY}

build-docker: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=${LDFLAGS} -o ${BINARIES_DIRECTORY}/${PROJECT_NAME}_docker ${MAIN_FILE_PATH}


CONTAINER_RUNNER=docker
COMPOSE_FILE_PATH=docker-compose.d/docker-compose.yaml

compose-up:
	$(CONTAINER_RUNNER) compose -f $(COMPOSE_FILE_PATH) up
