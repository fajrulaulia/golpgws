export APP_NAME=golpgws
export JWT_SECRET=secret
export SERVER_PORT=8081
export DB_URL=postgres://golpgwsuser:golpgwspass@localhost:5432/golpgwsdb?sslmode=disable

image-build:
	@docker build  \
	--build-arg SERVER_PORT=${SERVER_PORT} \
	-t ${APP_NAME} \
	-f ./deployments/Dockerfile .

image-run:
	@docker run \
	--env SERVER_PORT=${SERVER_PORT} \
	--env JWT_SECRET=${JWT_SECRET} \
	--env DB_URL=${DB_URL} \
	-p ${SERVER_PORT}:${SERVER_PORT} \
	--net host \
	${APP_NAME}

local-run:
	@go run cmd/main.go

local-migrate:
	@STRCONN=${DB_URL} ./migrations/migration 