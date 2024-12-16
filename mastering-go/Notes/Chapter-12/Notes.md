# Chapter 12: working with gRPC
gRPC Remote Procedure Calss, is an alternative to RESTful services that was developed by Google. The main advantage of gRPC is that it is faster than working with REST and JSON messages. gRPC uses the binary data format, which is lighter than RESTful services that work with JSON format.

## Introduction to gRPC
gRPC is an open source remote procedure call (RPC) system that was developed at Google back in 2015, is built on top of HTTP/2, allows you to create services easily, and used protocol buffers as the IDL which specifies the format of the interchanged messages and the service interface.

The advantages of gRPC include the following:

* THe use of binary format for data exchange makes gRPC faster than services that work with data in plain text format.

* gRPC can be used for streaming services and is built on top of HTTP/2.

However, the list of advantages should not make you think that gRPC is a panacea that does not have any flaws. always use the best tool or technology for the job.

## Protocol buffers

A protocol buffer (protobuf) is basically a method for serializing structured data. A part of protobuf is the IDL. As protobuf uses binary format for data exchange, it takes up less space than plain text serialization formats. However, data needs to be encoded and decoded to be machine-usable and human-readable, respectively. Protobuf has its own data types that are translated to natively supported data types of the programming language used.

Generally speaking, the IDL file is the center of each gRPC service because it defines the format of the data that is exchanged as well as the service interface. You cannot have a gRPC service without having a protobuf file at hand. It inclues the definition of services, methods of services, and the format of the messsages that are going to be exchanged.

## Defining an interface definition language file

```
syntax = "proto3"
```

If you do not specify that you want to use proto3, then the protocol buffer compiler assumes you want to use proto2.

```
option go_package = "protoapi";
```

The gRPC tools are going to generate Go code form .proto file. The previous line specifies that the name of the Go package that is going to ve created is protoapi.

After completing the code of proto file, the next important step is converting that file into a format that can be used by Go. The name of the protocol buffer compiler binary is protoc. On a linux machine, you need to install protobuf using your favorite package manager and protoc-gen-go using the go install github.com/golang/protobuf/protoc-gen-go@latest command. Similarly, you should install the protoc-gen-go-grpc executable by running the same command. 

So, the conversion process requires the next step.
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.
--go-grpc_opt=paths=source_relative protoapi.proto
```

\* Starting with Go 1.16, go install is the recommended way of building and installing packages in module mode. The use of go get is deprecated.

## Developing a gRPC server

```
type RandomServer struct {
    protoapi.UnimplementedEandomServer
}
```

This structure is named after the name of the gRPC service. This structure is going to implement the interface required by the gRPC server. The use of protoapi.UnimplementedRandomServer is required for the implementation of the interface.
```
server := grpc.NewServer()
```

Th4e previous statement creates a new gRPC server that is not attached to any specific gRPC service.
```
var randomServer RandomServer
protoapi.RegisterRandomServer(server, randomServer)
```

The previous statements call protoapi.RegisterRandomServer() to create a gRPC server for our specific service.



