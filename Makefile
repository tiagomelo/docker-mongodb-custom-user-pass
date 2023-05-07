SHELL = /bin/bash

DOCKER_MONGODB=docker exec -it mongodb-sample mongosh -u $(ADMIN_USER) -p $(ADMIN_PASSWORD) --authenticationDatabase admin

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: setup-db
## sets up MongoDB
setup-db: export ADMIN_USER=admin
setup-db: export ADMIN_PASSWORD=f3MdBEcz
setup-db:
	@ echo "Setting up MongoDB..."
	@ docker-compose up -d mongodb
	@ until $(DOCKER_MONGODB) --eval 'db.getUsers()' >/dev/null 2>&1 && exit 0; do \
	  >&2 echo "MongoDB not ready, sleeping for 5 secs..."; \
	  sleep 5 ; \
	done
	@ echo "... MongoDB is up and running!"


.PHONY: run
## run: runs the application
run: setup-db
	@ go run cmd/main.go

.PHONY: cleanup
## cleanup: removes MongoDB and associated volumes
cleanup:
	@ docker-compose down
	@ docker volume rm $$(docker volume ls -q)

.PHONY: test
## test: runs unit tests
test:
	@ go test -v ./...