# Great resource:

https://github.com/grpc-ecosystem/awesome-grpc

# Grpc tips to install protoc:

https://grpc.io/docs/protoc-installation/

# Grpc tips for golang:

https://grpc.io/docs/languages/go/quickstart/

# Google docs for compiling:

https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

# IMPORTANT:

Run protoc at the root of my project

If I don't run it a root, make sure to have this in my proto file:

```
option go_package = "./";
```

# How to test:

https://stackoverflow.com/a/52080545

^ what does dial do?

https://www.educative.io/edpresso/what-is-the-netdial-function-in-golang

Also, take a look at distributed services

# Plugin Documentation

## protoc-gen-go-grpc:

Creates the Go bindings for the service definition

https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc

## protoc-gen-go:

This tool generates Go language bindings of services in protobuf definition files for gRPC

https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go

## Commands I used:

Using only protoc-gen-go:

```
protoc -I ch2/pokemon \
         ch2/pokemon/pokemon.proto \
         --go_out=plugins=grpc:./ch2/pokemon
```

Using only protoc-gen-go and protoc-gen-go-grpc:

```
protoc ch2/pokemon/*.proto \
                    --go_out=. \
                    --go-grpc_out=. \
                    --go_opt=paths=source_relative \
                    --go-grpc_opt=paths=source_relative \
                    --proto_path=.
```

What's the difference? protoc-gen-go generates the client and gprc server, as well as the protobuf
from the service and message definition in the pokemon.proto file, thanks to the grpc plugin.

If I specify the --go-grpc_out flag and --go_out flag together, I will get a result that reflects
separation of concerns: the protobuf will go in one file (pokemon.pb.go) and the grpc service skeleton and client
stubs will go in another (pokemon_grpc.pb.go).

Essentially, I'll get the same results but protoc-gen-go will centralize everything in one file while the protoc-gen-grpc-go
generates two files due to separation of concerns.
