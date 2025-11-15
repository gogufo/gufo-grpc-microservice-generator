 🚀 **GUFO gRPC Microservice Generator**

[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](https://opensource.org/licenses/Apache-2.0)

A production-grade generator for creating GUFO-compatible microservices.

This tool creates a full, ready-to-run microservice following the **GUFO API Gateway** contract:

✔ unified middleware (logging, metrics, security, recovery, request-id)
✔ secure config loader
✔ gRPC + streaming support
✔ Prometheus metrics
✔ graceful shutdown
✔ fully structured project layout
✔ Docker-ready
✔ Makefile, documentation templates, examples

---

# 📦 **Build Docker Image**

```bash
docker build --no-cache -t amyerp/gufo_grpc_microservice_generator:latest -f Dockerfile .
```

Or the cached version:

```bash
docker build -t amyerp/gufo_grpc_microservice_generator:latest -f Dockerfile .
```

---

# 🛠 **Generate a New Microservice**

```bash
docker run --rm \
   -v $(pwd):/src \
   amyerp/gufo_grpc_microservice_generator:latest \
   /go/bin/grpccreator {microservice_name}
```

Example:

```bash
docker run --rm -v $(pwd):/src amyerp/gufo_grpc_microservice_generator:latest /go/bin/grpccreator orders
```

This will create:

```
orders/
  main.go
  router.go
  middleware/
  config/
  cron/
  entrypoint/
  Dockerfile
  Makefile
  README.md
```

---

# ⚙️ **Before Building the Microservice**

```bash
cd {microservice_name}
go mod tidy
```

---

# 🧰 **Using the Makefile (for the generator)**

The generator now includes a Makefile that works **with or without Docker**.

### ▶️ Create a new microservice (auto-detects Docker)

```bash
make create name=orders
```

If Docker is installed → the generator runs inside Docker.
If Docker is not installed → it uses the local binary (`./grpccreator`).

---

### ▶️ Build the generator binary (no Docker required)

```bash
make build
```

This produces:

```
./grpccreator
```

Run locally:

```bash
make local name=orders
```

or directly:

```bash
./grpccreator orders
```

---

### ▶️ Test the generated microservice

```bash
make test name=orders
```

This executes:

* `go mod tidy`
* `go build ./...`

---

### ▶️ Clean the generator binary

```bash
make clean
```

---

# 📚 **Learn More About GUFO**

Official GUFO API Gateway:
👉 [https://github.com/gogufo/gufo-api-gateway](https://github.com/gogufo/gufo-api-gateway)

Documentation includes:

* Gateway architecture
* Masterservice integration
* HMAC / SIGN / mTLS request validation
* Cron coordination
* Reverse gRPC protocol

---

# 🎯 **What the Generator Provides**

### ✔ Production-Ready Structure

* Clean, readable code
* Unified logging standard
* Security validation compatible with GUFO
* Plug-and-play streaming (Reverse.Stream)
* Heartbeat & health endpoints
* Metrics (requests + duration + panics)
* Graceful shutdown
* Config-driven behavior
* Panic-safe routing

### ✔ Developer Tools

* Makefile (create, build, local, test, clean)
* Multi-stage Dockerfile
* docker-compose example
* Documentation template
* FAQ & diagrams

---

# 📄 **LICENSE**

Apache 2.0


