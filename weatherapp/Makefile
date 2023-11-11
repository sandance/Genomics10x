.PHONY: docs clean

COMMAND = docker-compose run --rm weatherapp /bin/bash -c

all: build test

build:
	docker-compose build

run:
	docker-compose up

clean:
	rm -rf build

test:
	echo "To run this command please run  'make run' command before"
	curl  http://localhost:8080/query?weather=rain
	curl  http://localhost:8080/query?weather=rain&limit=5
	curl  http://localhost:8080/query?date=2012-06-04

dockerclean:
	docker system prune -f
	docker system prune -f --volumes

