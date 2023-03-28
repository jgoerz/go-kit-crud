# What is this?

This is an example service that exposes gRPC as well as HTTP/JSON for a simple
CRUD resource.

# How do I use this?

## Grab the code

```
git archive --remote=https://github.com/jgoerz/go-kit-crud \
  --format=tar.gz \
  --output go-kit-crud.tar.gz \
  --prefix=go-kit-crud/ \
  main
tar -xzf go-kit-crud.tar.gz
cd go-kit-crud
```

## Build protobuf files

The generated files have been committed.  This section is optional.

```
go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
```

```
go get -u google.golang.org/grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```
make proto
```

## Get dependencies

```
go mod tidy
```

## Run it


1. Start the server
```
make run
```
1. Interact using gRPC
```
make grpc-client
./grpc-client create-contact 123 jane doe true "123 Main Street" "secret"
```
1. Interact using HTTP
```
cd cmd/client/http
./main.sh
```


# References

1. Ryer, Matt, "Go Programming Blueprints, 2nd Edition." Packt, Jan. 2015, [URL](https://github.com/PacktPublishing/Go-Programming-Blueprints/tree/master/Chapter10)
2. Travis Jeffery. "Distributed Services with Go: Your Guide to Reliable, Scalable, and Maintainable Systems" Pragmatic Bookshelf. , March 11, 2021, [URL](https://www.amazon.com/gp/product/B0923C9WB5)
3. Bourgon, Peter et al. "Go kit Frequently asked questions." Go kit A toolkit for microservices. Feb 25, 2023, [URL](https://gokit.io/faq)
