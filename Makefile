

proto:
	@ clear && printf "protoc pb/*.proto \n --go_out=. \n --go-grpc_out=. \n --go_opt=paths=source_relative \n --go-grpc_opt=paths=source_relative \n --proto_path=.\n" && \
		protoc pb/*.proto \
			--go_out=. \
			--go-grpc_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--proto_path=.

build:
	@ clear && echo "go build ./..." && go build ./...

run:
	@ clear && echo "go run cmd/address-book-server/main.go" && go run cmd/address-book-server/main.go

test-service:
	@ clear && echo "go test -v -race ./internal/addressbook" && go test -v -race ./internal/addressbook

test-service-cover:
	@ clear && echo "go test -v -race -covermode=atomic -coverprofile=coverage.out ./internal/addressbook" && go test -v -race -covermode=atomic -coverprofile=coverage.out ./internal/addressbook 
	@ go tool cover -func=coverage.out

test-service-cover-html:
	@ clear && echo "go test -v -race -covermode=atomic -coverprofile=coverage.out ./internal/addressbook" && go test -v -race -covermode=atomic -coverprofile=coverage.out ./internal/addressbook
	@ go tool cover -html=coverage.out -o=coverage.html
	@ open ./coverage.html

grpc-client:
	@ clear && echo "go build -o grpc-client cmd/client/grpc/main.go" && go build -o grpc-client cmd/client/grpc/main.go 
