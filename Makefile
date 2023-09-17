postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:lk515405@simple-bank.cwisktgnesvo.ap-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:lk515405@simple-bank.cwisktgnesvo.ap-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:lk515405@simple-bank.cwisktgnesvo.ap-east-1.rds.amazonaws.com:5432/simple_bank" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:lk515405@simple-bank.cwisktgnesvo.ap-east-1.rds.amazonaws.com:5432/simple_bank" -verbose down 1

# migrateup:
# 	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# migrateup1:
# 	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

# migratedown:
# 	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# migratedown1:
# 	migrate -path db/migration -database "postgresql://root:L_515405@simple-bank.c27nywywxxl2.ap-northeast-3.rds.amazonaws.com:5432/simple_bank" -verbose down 1

# awsmigrateup:
# 	migrate -path db/migration -database "postgresql://root:L_515405@simple-bank.c27nywywxxl2.ap-northeast-3.rds.amazonaws.com:5432/simple_bank" -verbose up

# awsmigrateup1:
# 	migrate -path db/migration -database "postgresql://root:L_515405@simple-bank.c27nywywxxl2.ap-northeast-3.rds.amazonaws.com:5432/simple_bank" -verbose up 1

# awsmigratedown:
# 	migrate -path db/migration -database "postgresql://root:L_515405@simple-bank.c27nywywxxl2.ap-northeast-3.rds.amazonaws.com:5432/simple_bank" -verbose down

# awsmigratedown1:
# 	migrate -path db/migration -database "postgresql://root:L_515405@simple-bank.c27nywywxxl2.ap-northeast-3.rds.amazonaws.com:5432/simple_bank" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -source=./db/sqlc/store.go -destination=./db/mock/store.go -package=mockdb

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 awsmigratedown awsmigratedown1 awsmigrateup awsmigrateup1 sqlc test server mock