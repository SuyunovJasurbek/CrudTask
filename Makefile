include .env

go:
	go run cmd/main.go

watch:
	google-chrome 'http://${HTTP_HOST}:${HTTP_PORT}/swagger/index.html'
	make go

swag:	
	swag init -g ./cmd/main.go -o ./docs

login-psql:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} psql ${POSTGRES_DB} ${POSTGRES_USER}

createdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB}

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_DB=${POSTGRES_DB} -p ${POSTGRES_PORT}:5432 -d postgres:15-alpine3.16

migration-up:
	migrate -path ./migration/ -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' up

migration-down:
	migrate -path ./migration/ -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' down

stop-psql:
	docker stop ${DOCKER_POSTGRES_CONTAINER_NAME}