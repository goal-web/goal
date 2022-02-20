# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=go test
GOGET=$(GOCMD) get
DOCKER_TAG=goal

run:
	$(GOCMD) run main.go run

migrate:
	$(GOCMD) run main.go migrate

migrate-rollback:
	$(GOCMD) run main.go migrate:rollback

migrate-refresh:
	$(GOCMD) run main.go migrate:refresh

migrate-reset:
	$(GOCMD) run main.go migrate:reset

migrate-status:
	$(GOCMD) run main.go migrate:status

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin_linux -v ./

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin_windows.exe -v ./

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./bin_mac -v ./

win-build-linux:
	SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 $(GOBUILD) -o ./bin_linux -v ./

win-build-mac:
	SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 $(GOBUILD) -o ./bin_mac -v ./

win-build:
	$(GOBUILD) -o ./bin_goal.exe -v ./

build:
	$(GOBUILD) -o ./bin_goal -v ./

test:
	$(GOTEST) -json ./tests

pack: build-linux
	docker build -t $(DOCKER_TAG) .

test-and-pack: test pack
