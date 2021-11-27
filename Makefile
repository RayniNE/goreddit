.PHONY: postgres adminer migrate


postgres: 
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=admin postgres

adminer: 
	docker run --rm -ti --network host -p 8080:8080 adminer

migrate:
	migrate -source file://migrations \
			-database postgresql://postgres:admin@localhost/postgres?sslmode=disable up