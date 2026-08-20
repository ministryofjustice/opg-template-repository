build:
	docker compose build --no-cache {{APP_NAME}}

build-dev:
	docker compose -f docker-compose.yml -f docker/docker-compose.dev.yml build --no-cache --parallel {{APP_NAME}} yarn

clean:
	docker compose -f docker-compose.yml -f docker/docker-compose.dev.yml down --remove-orphans

compile-assets:
	docker compose run --rm yarn build

dev-up: clean build-dev
	docker compose -f docker-compose.yml -f docker/docker-compose.dev.yml up {{APP_NAME}} yarn

up: clean compile-assets build
	docker compose up -d --wait {{APP_NAME}}

down:
	docker compose down --remove-orphans