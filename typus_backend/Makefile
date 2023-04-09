# Makefile for typus_backend


# Variables 

## Including environment variables from the .env file
source .env


# Targets

## build: Build the executable file
.PHONY: build
build:
	go build -v ./cmd/api


## test: Run all tests (tests are not added yet)
.PHONY: test
test:
	go test -v -race -timeout 30s ./...


## makemigration: Create up and down migration files in /migrations directory.
## 				  Migration name needs to be passed in args.
## Usage: 
## 		make makemigration name="..."
.PHONY: makemigration
	migrate create -ext sql -dir migrations $(name)


## migrate_up: Run all the up migration files that were not applied to the database yet.
.PHONY: migrate_up
	migrate -path migrations -database \
	"postgres://$(POSTGRES_HOST)/$(POSTGRES_NAME)?sslmode=$(POSTGRES_SSLMODE)&user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)" \
	-verbose up


## migrate_down: Run the latest down migration file. 
.PHONY: migrate_down
	migrate -path migrations -database \
	"postgres://$(POSTGRES_HOST)/$(POSTGRES_NAME)?sslmode=$(POSTGRES_SSLMODE)&user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)" \
	-verbose down


## makemigration: Apply a certain migration version.
##				  Version needs to be passed in args.
## Usage: 
## 		make migrate_fix version="..."
.PHONY: migrate_fix
	migrate -path migrations -database \
	"postgres://$(POSTGRES_HOST)/$(POSTGRES_NAME)?sslmode=$(POSTGRES_SSLMODE)&user=$(POSTGRES_USER)&password=$(POSTGRES_PASSWORD)" \
	force $(version)


# Setting default make target to build.
.DEFAULT_GOAL := build

# End of Makefile