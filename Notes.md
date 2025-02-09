
`option go_package = "/gogen";` Ea proto file ka line dockerfile ke `RUN protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto` `_out` path se append ho jata hai and `/ticker/domain/gogen` pai apne generated code ko rakhta hai.

```bash
docker build -t grpcgen -f dockerfile.ticker .
docker run -it grpcgen
docker cp 5d7352636391:/ticker/domain/gogen ./ticker/domain
```