include ~/go-utils/local.env


APP=go_utils
APP_EXECUTABLE="./out/$(APP)"
ENVIRONMENT="development"

clean: ##@build remove executable
	rm -f $(APP_EXECUTABLE)

compile: ##@build build the executable
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)


db-create:  ##@database create db
	createdb -h $(DB_HOST) -U $(DB_USER) -O$(DB_USER) -Eutf8 $(DB_NAME)

db-apply_extension:  ##@database create uuid extension
	psql -h $(DB_HOST) -d $(DB_NAME) -U $(DB_USER) -c 'CREATE EXTENSION if not exists "uuid-ossp"'

db-migrate: db-apply_extension  ##@database run migrations
	$(APP_EXECUTABLE) migrate:run

db-rollback: db-apply_extension  ##@database rollback migrations
	$(APP_EXECUTABLE) migrate:rollback

db-drop:  ##@database drop db
	dropdb -h $(DB_HOST) -U $(DB_USER) --if-exists $(DB_NAME)

db-reset: db-drop db-create db-migrate  ##@database resets local db to fresh db

set-dirty-false:
	psql -h $(DB_HOST) -d $(DB_NAME) -U $(DB_USER) -c 'UPDATE schema_migrations SET dirty=false;'

drop-schema-migrations:
	psql -h $(DB_HOST) -d $(DB_NAME) -U $(DB_USER) -c 'DROP TABLE schema_migrations;'

#Server related makes
build:
	go build -o $(APP_EXECUTABLE)

start-server: build
	$(APP_EXECUTABLE)