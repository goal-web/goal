DOCKER_TAG=goal

run:
	go run main.go run

build:
	go build -o ./bin -v ./

test:
	go test -json ./tests

image:
	docker build -t $(DOCKER_TAG) .
