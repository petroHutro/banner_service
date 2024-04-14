.PHONY: run
run:
	sudo docker compose --env-file ./docker_run.env up

.PHONY: stop
stop:
	sudo docker compose --env-file ./docker_run.env stop

.PHONY: down
down:
	sudo docker compose --env-file ./docker_run.env down

.PHONY: build-banner
build-banner:
	go build -o server -v ./cmd/banner

.PHONY: build-docker-banner
build-docker-banner:
	sudo docker build --no-cache --network host -f ./docker/banner.Dockerfile . --tag banner
