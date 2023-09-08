DOCKER_TAG=goal

run:
	go run main.go run

build:
	go build -o ./bin -v ./

test:
	go test -json ./tests

image:
	docker build -t $(DOCKER_TAG) .

migrate:
	go run bootstrap/console/main.go migrate

migrate-rollback:
	go run bootstrap/console/main.go migrate:rollback

migrate-refresh:
	go run bootstrap/console/main.go migrate:refresh

migrate-reset:
	go run bootstrap/console/main.go migrate:reset

migrate-status:
	go run bootstrap/console/main.go migrate:status

make-migration:
	go run bootstrap/console/main.go make:migration $(NAME)
