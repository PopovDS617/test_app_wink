LOCAL_BIN:=$(CURDIR)/bin
BINARY_NAME=app
include .env

## install-deps: installs dependencies
install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=${LOCAL_BIN} go install github.com/gojuno/minimock/v3/cmd/minimock@latest
## get-deps: getting dependencies
get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/${BINARY_NAME} ./cmd/grpc_server
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env GRPC_HOST=${GRPC_HOST} GRPC_PORT=${GRPC_PORT} ORIGINAL_HOST=${ORIGINAL_HOST} CDN_HOST=${CDN_HOST} ./bin/${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ./bin/${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## race: checks for data race
race:
	cd ./cmd && \
	go run -race .

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...

## testrace: checks tests for data race
testrace:
	go test -race -v ./...

## test-coverage: collects test coverage
test-coverage:
	go test -coverprofile=coverage.out -v ./...

## testout: displays test coverage as html in browser
testout:
	go tool cover -html=coverage.out

## db: builds docker images
db:
	docker-compose build

## du: starts all the docker-compose containers in detached mode
du:
	docker-compose up -d

## dd: stops all docker-compose containers
dd :
	docker-compose down

## lint: runs linter for all files in the project
lint:
	cd ./bin && \
	golangci-lint run ../...

## generate-grpc: generates grpc files for all api
generate-grpc:
	make generate-lb-api

## generate-lb-api: generates gprc files for lb api
generate-lb-api:
	mkdir -p pkg/lb_v1
	protoc --proto_path api/lb_v1 \
	--go_out=pkg/lb_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/lb_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/lb_v1/lb.proto