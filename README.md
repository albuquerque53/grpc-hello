# gRPC Hello

> A simple gRPC API and Client written in Go.

## :bookmark_tabs: Table of contents

- [gRPC Hello](#grpc-hello)
  * [:bookmark_tabs: Table of contents](#-bookmark-tabs--table-of-contents)
  * [:mag: What is this?](#-mag--what-is-this-)
  * [:eyes: How to use/run?](#-eyes--how-to-use-run-this-)
    + [:computer: The native way](#-computer--the-native-way)
      - [Set-up server](#set-up-server)
      - [Making gRPC calls with Client](#making-grpc-calls-with-client)
    + [:whale: The Docker way](#-whale--the-docker-way)
      - [Set-up server](#set-up-server-1)
      - [Making gRPC calls with Client](#making-grpc-calls-with-client-1)
      
## :mag: What is this?

This repository contains a simple project about gRPC, inside of it you'll find a gRPC API that receives the `name` parameter and returns a `Hello, <name>!` response. <br>

```proto
service Hello {
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string message = 1;
}
```
> Take a look into [proto/hello.proto](proto/hello.proto)

For testing purpouses, this project also contains a gRPC client that will send a request to API with `name` parameter and print the response received.

## :eyes: How to use/run?

### :computer: The native way

#### Set-up server

1 - First, install the dependencies:

```
make install
```

2 - Now, start the server:

```
make api
```

> Now you have the API running on the port 9000 (in your host and container)


#### Making gRPC calls with Client


In another terminal, you can use `make call` command passing the `name` you want to say "Hello", like:
```
make call 'John Doe'
```

And the API will return something like:
```
message:"Hello, John Doe!"
```

<hr>

### :whale: The Docker way

#### Set-up server

1 - First, start the application container:
```
make up
```

2 - Get into the container:
```
make app
```

3 - Now, install the dependencies and start the server:
```
make install && make api
```

> Now you have the API running on the port 9000 (in your host and container)


#### Making gRPC calls with Client

1 - In another terminal, get into the container
```
make app
```

2 - Now, to call the API you can use `make call` command passing the `name` you want to say "Hello", like:
```
make call 'John Doe'
```

And the API will return something like:
```
message:"Hello, John Doe!"
```
