# ============================================
#  GUFO MICROservice GENERATOR - MAKEFILE
#  Works WITH or WITHOUT Docker
# ============================================

GEN_BIN=./grpccreator
IMAGE=amyerp/gufo_grpc_microservice_generator
TAG=latest

# Detect docker availability
HAVE_DOCKER := $(shell command -v docker 2>/dev/null)

# --------------------------------------------
# Create new microservice
# Usage: make create name=orders
# --------------------------------------------
create:
ifndef name
	$(error microservice name is required: make create name=orders)
endif
ifeq ($(HAVE_DOCKER),)
	@echo "Docker not found. Using local binary generator..."
	$(GEN_BIN) $(name)
else
	@echo "Using Docker generator..."
	docker run --rm -v $(PWD):/src $(IMAGE):$(TAG) /go/bin/grpccreator $(name)
endif

# --------------------------------------------
# Build the generator binary locally
# --------------------------------------------
build:
	go build -o grpccreator main.go

# --------------------------------------------
# Run generator locally without docker
# Usage: make local name=orders
# --------------------------------------------
local:
ifndef name
	$(error microservice name is required: make local name=orders)
endif
	$(GEN_BIN) $(name)

# --------------------------------------------
# Test build of generated microservice
# Usage: make test name=orders
# --------------------------------------------
test:
ifndef name
	$(error microservice name is required: make test name=orders)
endif
	cd $(name) && go mod tidy && go build ./...

# --------------------------------------------
# Clean local binary
# --------------------------------------------
clean:
	rm -f grpccreator
