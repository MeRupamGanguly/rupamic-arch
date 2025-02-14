# rupamic-arch

![Image](https://github.com/user-attachments/assets/0579375e-e9ac-442a-90fd-2a515db575f4)

## First create Proto file for ByDirectional Stream for Ticker in Domain:

```proto
syntax="proto3";

package domain;

option go_package = "/domain";

message TickerRequest {
    string symbol = 1;
}

message TickerResponse {
    string symbol = 1;
    double ltp = 2;
    string timestamp = 3; 
}

service TickerStreamService{
    rpc TickerStream (stream TickerRequest)returns(stream TickerResponse);
}
```
option go_package = "github.com/MeRupamGanguly/rupamic"  this line decides where generated file will store.

## Steps to gen grpc-go Codes:
```bash
docker run -it golang:1.23.6 /bin/bash
apt update
apt install protobuf-compiler
go version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
docker cp ticker/ e3a5b0647829:/ticker/ #New Terminal from Host
protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto
```
We can ran these commands with automation using dockerfile.

## Second Create Dockerfile for generating grpc-go codes:

```dockerfile
FROM golang:1.23.6
RUN apt update
COPY /ticker /ticker
RUN apt install protobuf-compiler -y
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN mkdir -p /ticker/domain/gogen
RUN protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto

```
## Build from Dockerfile
```bash
docker build -t grpcgen -f dockerfile.ticker .
docker run -it grpcgen
docker cp 52509fb095ed:/ticker/github.com/MeRupamGanguly/rupamic/ .
```
