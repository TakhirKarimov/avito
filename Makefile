SHELL := /bin/bash

all: up logs

up:
	docker-compose up -d

down:
	docker-compose down

restart: down up

build:
	docker-compose build

clean:
	docker-compose down --rmi all --volumes --remove-orphans

logs:
	docker-compose logs -f

.PHONY: up down restart build clean logs