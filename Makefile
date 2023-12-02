# SERVER option (local|develop|staging|production)
SERVER=local
GO_AIR_VERSION=latest

# init
git/init: git/commit-template
docker/db/init: docker/up sleep db/migrate

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION)

go/lint:
	staticcheck ./...

# db
ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template ./ent/template --template glob="./ent/template/*.tmpl" ./ent/schema

db/migrate:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query migrate

db/seed:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query seed

# server
server/run:
	go run ./cmd/public/migration.go -server $(SERVER)

server/run-air:
	air -c ./cmd/public/air.toml

# docker
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
	make docker/db/init &&\
	make db/seed &&\
	make server/run-air

# other
sleep:
	sleep 20
