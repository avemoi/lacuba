mysql:
	docker run --name lacubamydbcontainer -e MYSQL_ROOT_PASSWORD=mypassword -p 3306:3306 -d mysql:latest
createdb:
	docker exec -it lacubamydbcontainer mysql -u root -pmypassword -e "CREATE DATABASE lacubadb;"

dropdb:
	docker exec -it lacubacontainer dropdb --username=postgres  lacubadb

migrateup:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3306)/lacubadb" --verbose up

migratedown:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3306)/lacubadb" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

build:
	go build -ldflags="-I cmd/api -I db/sqlc"


.PHONY: postgres createdb dropdb migrateup migratedown