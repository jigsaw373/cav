test: 
	go test -v -cover ./...

server: 
	go run ./cmd/api/ 

mock: 
	mockgen -destination db/mock/store.go -package mockdb  github.com/jigsaw373/sb/db/sqlc Store

proto: 
	rm -rf pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: createdb migrateup migratedown postgres sqlc server proto
