# path to docker compose file
DCOMPOSE:=docker-compose.yml

# improve build time
DOCKER_BUILD_KIT:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

down:
	docker-compose -f ${DCOMPOSE} down --remove-orphans

build:
	${DOCKER_BUILD_KIT} docker-compose build

up:
	docker-compose up --build -d --remove-orphans

# Vendoring is useful for local debugging since you don't have to
# reinstall all packages again and again in docker
mod:
	go mod tidy -compat=1.21 && go mod vendor && go install ./...

tests:
	go test ./internal/... -cover -coverprofile=cover.out -coverpkg=./internal/...
	cat cover.out | fgrep -v "http" | fgrep -v "config" | fgrep -v "mocks" > cover1.out
	go tool cover -func=cover1.out

mock:
	mockgen -source=internal/http-server/handler/actor.go -destination=mocks/service/mock_actor.go
	mockgen -source=internal/http-server/handler/film.go -destination=mocks/service/mock_film.go
	mockgen -source=internal/http-server/handler/auth.go -destination=mocks/service/mock_auth.go
	mockgen -source=internal/service/actor.go -destination=mocks/db/mock_actor.go
	mockgen -source=internal/service/film.go -destination=mocks/db/mock_film.go
	mockgen -source=internal/service/auth.go -destination=mocks/db/mock_auth.go

swag:
	swag init -g cmd/app/main.go

lint:
	golangci-lint run