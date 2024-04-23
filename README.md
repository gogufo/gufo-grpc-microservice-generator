# GUFO Microservices Generator

## How to bild Docker image

```
docker build --no-cache -t amyerp/gufo_grpc_microservice_generator:latest -f Dockerfile .
```

## How to Create GRPC Microservice

1. Run Docker image to create GRPC microservice structure
```
docker run --rm -v $(pwd):/src amyerp/gufo_grpc_microservice_generator:latest /go/bin/grpccreator {microservice_name}
```

2. Before build MicroServiceName

```
go mod tidy
```
