

```bash
docker build -t grpcgen -f dockerfile.ticker .
docker run -it grpcgen
docker cp 5d7352636391:/ticker/domain/gogen ./ticker/domain
```