RUN=docker-compose run --service-ports --rm --workdir="/go/src/github.com/toyohashi6140/eenen-api" eenen-api


up:
	docker compose up -d eenen-api

down:
	docker compose rm -fsv eenen-api

start:  
	docker compose start eenen-api

stop: 
	docker compose stop eenen-api

gqlgen:
	${RUN} gqlgen

gql-run:
	make stop
	${RUN} gqlgen
	make start