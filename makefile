SHELL := /bin/bash

run:
	go run main.go

#build:
#	go build -ldflags "-X main.build=local"

VERSION := 1.0

all: service

service:
	sudo docker build \
		-f zarf/docker/dockerfile \
		-t go_sales_test:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

KIND_CLUSTER := ardan-starter-cluster
kind-up:
	sudo kind create cluster \
		--image kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml

kind-down:
	sudo kind delete cluster --name $(KIND_CLUSTER)

kind-status:
	sudo kubectl get nodes -o wide
	sudo kubectl get svc -o wide
	sudo kubectl get pods -o wide --watch --all-namespaces