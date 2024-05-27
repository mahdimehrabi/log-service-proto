include .env
MIGRATE=migrate -path=entities/migrations -database "" -verbose


db-migrate-up:
		$(MIGRATE) up
db-migrate-down:
		$(MIGRATE) down
db-force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

db-goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

db-drop:
		$(MIGRATE) drop

db-create-migration:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir entities/migrations  $$NAME
test-all:
	${DOCKER_COMMAND} exec web go test ./tests/tests/...

