DOCKER_COMPOSE_FILE = build/docker-compose.yaml

.SILENT:

build:
	go build -o main .

run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up


unit-tests:
	go test -v ./tests/unit
