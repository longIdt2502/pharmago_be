postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Hoanglong2502 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root pharmago

dropdb:
	docker exec -it postgres dropdb pharmago

migrateup:
	migrate -path db/migration -database "postgresql://root:Hoanglong2502@localhost:5432/pharmago?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:Hoanglong2502@localhost:5432/pharmago?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:Hoanglong2502@localhost:5432/pharmago?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:Hoanglong2502@localhost:5432/pharmago?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	rm -f docs/swagger/*.swagger.json
	protoc --proto_path=proto --proto_path=proto/entities  --proto_path=proto/payloads --proto_path=proto/rpc \
		--proto_path=proto/rpc/address --proto_path=proto/rpc/company --proto_path=proto/rpc/order \
		--proto_path=proto/rpc/warehouse --proto_path=proto/rpc/product --proto_path=proto/rpc/customer \
		--proto_path=proto/rpc/supplier \
	 	--go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=pharmago \
        proto/*.proto proto/entities/*.proto proto/payloads/*.proto proto/rpc/*.proto proto/rpc/address/*.proto proto/rpc/company/*.proto \
        proto/rpc/product/*.proto proto/rpc/warehouse/*.proto proto/rpc/order/*.proto proto/rpc/customer/*.proto proto/rpc/supplier/*.proto
	cp pb/entities/*pb.go pb
	cp pb/payloads/*pb.go pb
	cp pb/rpc/*pb.go pb
	cp pb/rpc/address/*pb.go pb
	cp pb/rpc/company/*pb.go pb
	cp pb/rpc/product/*pb.go pb
	cp pb/rpc/warehouse/*pb.go pb
	cp pb/rpc/order/*pb.go pb
	cp pb/rpc/customer/*pb.go pb
	cp pb/rpc/supplier/*pb.go pb
	rm -r pb/entities
	rm -r pb/payloads
	rm -r pb/rpc
	statik -src=./docs/swagger -dest=./docs

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

new_service:
	mkdir $(name)
	cd $(name)
	mkdir proto
	mkdir pb
	mkdir db
	mkdir docs
	mkdir gapi
	touch main.go



.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration server sqlc proto evans new_service