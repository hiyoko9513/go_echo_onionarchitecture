# SERVER option (local|develop|staging|production)
SERVER=local

lint:
	staticcheck ./...

ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template ./ent/template --template glob="./ent/template/*.tmpl" ./ent/schema

docker/up:
	docker-compose --env-file ./cmd/public/.env.$(SERVER) up -d

db/migrate:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query migrate

db/seed:
	go run ./cmd/cli/db/main.go -server $(SERVER) -query seed

server/run:
	go run ./cmd/public/migration.go -server $(SERVER)

server/run-air:
	air -c ./cmd/public/air.toml
