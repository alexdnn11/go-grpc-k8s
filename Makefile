.PHONY: up down generate clean gen-pb gen-certs

help:
	@echo "gRPC & k8s Demo"
	@echo ""
	@echo "generate: generate Go code from *.proto files"
	@echo "          build docker containers."
	@echo "up: bring up the network"
	@echo "down: clear the network"
	@echo "clean: remove docker containers and volumes"
	@echo ""

generate:
	./scripts/gen-pb.sh
	./scripts/build-containers.sh

gen-pb:
	./scripts/gen-pb.sh

gen-certs:
	./scripts/gen-certs.sh

up:
	./network.sh -m up

down:
	./network.sh -m down

clean:
	./network.sh -m clean




















