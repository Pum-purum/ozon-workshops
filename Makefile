.PHONY: up down build logs sh

up:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

sh:
	docker compose exec api bash
