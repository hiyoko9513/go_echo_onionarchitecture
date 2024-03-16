# SERVER option (local|develop|staging|production)
SERVER=local

GO_AIR_VERSION=latest
GO_STATICCHECK_VERSION=latest
GO_OAPI_CODEGEN_VERSION=latest

# init
git/init: git/commit-template
docker/db/init:

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION) &&\
	go install honnef.co/go/tools/cmd/staticcheck@$(GO_STATICCHECK_VERSION)
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@$(GO_OAPI_CODEGEN_VERSION)

# staticcheck
go/lint:
	staticcheck ./...

# ent
ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./internal/pkg/ent/template/*.tmpl" ./internal/pkg/ent/schema

# db
db/migrate:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query migrate
db/seed:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query seed

# server
server/run:
	go run ./cmd/app/main.go -server $(SERVER)
server/run-air: # local only
	air -c ./cmd/app/air.toml

# docker
docker/build:
	docker-compose --env-file ./cmd/app/.env.$(SERVER) build
docker/up:
	docker-compose --env-file ./cmd/app/.env.$(SERVER) up -d

# oapi
oapi/gen/app:
	docker run --rm -v $(PWD):/spec redocly/cli:latest bundle api/app/openapi-spec/root.yml -o api/app/openapi-spec/root.gen.yml
oapi/run/app: oapi/gen/app
	docker run -p 8081:8080 -v $(PWD)/api/app/openapi-spec:/usr/share/nginx/html/api -e API_URL=api/root.gen.yml swaggerapi/swagger-ui
oapi/codegen/app: oapi/gen/app
	oapi-codegen  --config api/app/openapi-spec/config.yml api/app/openapi-spec/root.gen.yml

# git
git/commit-template:
	cp ./.github/.gitmessage.txt.example ./.github/.gitmessage.txt &&\
    git config commit.template ./.github/.gitmessage.txt &&\
    git config --add commit.cleanup strip

# quick start
quick-start:
	go mod tidy &&\
	make go/install/tools &&\
	source ~/.zshrc &&\
	make ent/gen &&\
	make oapi/codegen/app &&\
	make docker/up &&\
	make sleep &&\
	make db/migrate &&\
	make docker/db/init &&\
	make db/seed &&\
	make server/run-air

# other
sleep:
	sleep 20
