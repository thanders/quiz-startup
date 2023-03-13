# quiz-startup
A quiz microservice project written in GO and GRPC

## Description

`quiz-startup` was built to help me learn about Go, GRPC and Protocol Buffers.

### What is it?

It is a simple go module with multiple packages. The module architecture is as follows:
```
github.com/thanders/quiz-startup (module)
- - broker
- - - proto
- - - client (package)
- - - server (package)
```

What is Protocol Buffers?

## Usage

### scripts

#### build

Builds the binaries for all microservices:
```
./createBinaries.sh
```

#### proto

Run the protocol buffer compiler to create the `pb.go` and `grpc.pb.go` files. The generated files help go to understand the proto contract.
```
./createProtoFiles.sh
``` 

### Binaries

Each project will contain a `/bin` folder that contains a compiled binary. On a unix system the binary can be executed from a shell like any other executable.

For example:
```
./broker/server/bin/server
```

```
./broker/client/bin/client
```


