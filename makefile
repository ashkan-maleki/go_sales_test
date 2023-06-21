SHELL := /bin/bash

run:
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go

#build:
#	go build -ldflags "-X main.build=local"

VERSION := 1.0.0


all: sales-api

sales-api:
	sudo docker build \
		-f zarf/docker/sales-api/Dockerfile \
		-t ashkanmaleki/go-sales-api:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

docker-push:
	docker push ashkanmaleki/go-sales-api:$(VERSION)

# ==============================================================================
# Modules support

KIND_CLUSTER := ardan-starter-cluster
kind-up:
	sudo kind create cluster \
		--image kindest/node:v1.21.1@sha256:69860bda5563ac81e3c0057d654b5253219618a22ec3a346306239bba8cfa1a6 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml
	sudo kubectl config set-context --current --namespace=sales-system

kind-down:
	sudo kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	cd zarf/k8s/kind/sales-pod; kustomize edit set image sales-api-image=ashkanmaleki/go-sales-api:$(VERSION)
	sudo kind load docker-image ashkanmaleki/go-sales-api:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	sudo kubectl kustomize zarf/k8s/kind/sales-pod/ | sudo kubectl apply -f -

kind-status:
	sudo kubectl get nodes -o wide
	sudo kubectl get svc -o wide
	sudo kubectl get pods -o wide --watch --all-namespaces

kind-status-sales:
	sudo kubectl get pods -o wide --watch

kind-logs:
	sudo kubectl logs -l app=sales --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

kind-restart:
	sudo kubectl rollout restart deployment sales-pod

kind-update: all kind-load kind-restart

kind-update-apply: all kind-load kind-apply

kind-describe:
	sudo kubectl describe pod -l app=sales

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor


#===============================
# Git commands
#===============================

git-commit:
	git add .
	git commit -m "$(CMSG)"
	git push