# SERVER option (local|develop|staging|production)

lint:
	staticcheck ./...

docker/up:
	docker-compose up -d --env-file ./cmd/public/.env.$(SERVER)

db/migrate:
	go run ./cmd/cli/db.go -server $(SERVER) -migrate

server/run:
	go run ./cmd/public/migration.go -server $(SERVER)
