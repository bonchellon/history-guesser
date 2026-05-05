.PHONY: dev build migrate seed test

dev:
	docker compose up --build

build:
	cd server && go build ./cmd/api
	cd client && npm run build

migrate:
	cd server && goose -dir migrations postgres "$$DATABASE_URL" up

seed:
	psql "$$DATABASE_URL" -f server/sql/seed.sql

test:
	cd server && go test ./...
	cd client && npm run test -- --runInBand
