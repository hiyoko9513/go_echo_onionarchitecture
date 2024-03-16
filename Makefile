# SERVER option (local|develop|staging|production)
SERVER=local

GO_AIR_VERSION=latest
GO_STATICCHECK_VERSION=latest

# init
git/init: git/commit-template
docker/db/init:

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION) &&\
	go install honnef.co/go/tools/cmd/staticcheck@$(GO_STATICCHECK_VERSION)

go/lint:
	staticcheck ./...

# db
ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./internal/pkg/mypubliclib/ent/template/*.tmpl" ./internal/pkg/mypubliclib/ent/schema

db/migrate:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query migrate

db/seed:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query seed

# server
server/run:
	go run ./cmd/public/main.go -server $(SERVER)
server/run-air: # local only
	air -c ./cmd/public/air.toml

# docker
docker/build:
	docker-compose --env-file ./cmd/public/.env.$(SERVER) build
docker/up:
	docker-compose --env-file ./cmd/public/.env.$(SERVER) up -d

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
	make docker/up &&\
	make sleep &&\
	make db/migrate &&\
	make docker/db/init &&\
	make db/seed &&\
	make server/run-air

# other
sleep:
	sleep 20
