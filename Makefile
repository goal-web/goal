DOCKER_TAG=goal

run:
	go run bootstrap/app/main.go run

build:
	go build -o ./bin_goal -v ./

test:
	go test -json ./tests

pack:
	docker build -t $(DOCKER_TAG) .
