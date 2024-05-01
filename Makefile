.SILENT:
.DEFAULT_GOAL := fast-run

include ./.env

export ALPINE_VER
export DB_CONNECTION_LATENCY_MILLISECONDS
export DB_CONNECTION_SHOW_STATUS
export DB_CONNECTION_TIMEOUT_SECONDS
export DB_DATABASE
export DB_DRIVER
export DB_HOST
export DB_HOSTNAME
export DB_LOCAL_DIR
export DB_MIGRATION_CONTAINER_NAME
export DB_MIGRATION_PORT
export DB_NAME
export DBO_CONNECTION_LATENCY_MILLISECONDS
export DBO_CONNECTION_SHOW_STATUS
export DBO_CONNECTION_TIMEOUT_SECONDS
export DBO_DRIVER
export DBO_HOST
export DBO_NAME
export DBO_PASSWORD
export DBO_PORT
export DBO_SSL_ENABLE
export DBO_USERNAME
export DB_PASSWORD
export DB_PORT
export DB_SSL_ENABLE
export DB_USERNAME
export GO_VER
export HASHING_PASSWORD_SALT
export HASHING_TOKEN_SIGNING_KEY
export HTTP_SERVER_LISTEN_PORT
export NGINX_VER
export POSTGRES_VER
export SERVER_APP_NAME
export SERVER_HOST
export SERVER_PORT
export SERVER_READ_TIMEOUT_SECONDS
export SERVER_WRITE_TIMEOUT_SECONDS
export WEB_LOCAL_DIR
export WEB_PORT

.PHONY: migrate
migrate: run-dbc migrate-up stop-dbc

.PHONY: run-dbc
run-dbc:
	docker run \
	--rm \
	--name ${DB_MIGRATION_CONTAINER_NAME} \
	-d \
	-e PGDATA=/data \
	-e POSTGRES_USER=${DB_USERNAME} \
	-e POSTGRES_PASSWORD=${DB_PASSWORD} \
	-e POSTGRES_DB=${DB_DATABASE} \
	-v ${DB_LOCAL_DIR}:/data \
	-p ${DB_MIGRATION_PORT}:5432 \
	postgres:${POSTGRES_VER}-alpine${ALPINE_VER}

.PHONY: migrate-up
migrate-up:
	./wait-for-postgres.sh \
	${DB_PASSWORD} \
	${DB_HOSTNAME} \
	${DB_MIGRATION_PORT} \
	${DB_USERNAME} \
	migrate \
	-path ./schema \
	-database ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOSTNAME}:${DB_MIGRATION_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE} \
	up

.PHONY: stop-dbc
stop-dbc:
	docker stop ${DB_MIGRATION_CONTAINER_NAME}

.PHONY: build
build:
	go build -o ./build/${SERVER_APP_NAME} ./cmd/${SERVER_APP_NAME}/main.go

.PHONY: run
run:
	./build/${SERVER_APP_NAME}

.PHONY: fast-run
fast-run:
	go run ./cmd/${SERVER_APP_NAME}/main.go

.PHONY swag:
swag:
	swag init -g ./cmd/${SERVER_APP_NAME}/main.go -o ./docs
