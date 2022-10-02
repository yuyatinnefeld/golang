postgres:
	docker run --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PWD) -d $(IMG_NAME)
createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)
dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb $(POSTGRES_DB)
migrateup:
	migrate -path simplebank/db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PWD)@127.0.0.1:5432/$(POSTGRES_DB)?sslmode=disable" -verbose up 
migratedown:
	migrate -path simplebank/db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PWD)@127.0.0.1:5432/$(POSTGRES_DB)?sslmode=disable" -verbose down 

.PHONY: postgres createdb dropdb migrateup migratedown 