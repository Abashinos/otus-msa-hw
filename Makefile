# Docker

.PHONY: build push latest tag

DOCKER_TAG=latest

build:
	docker build --build-arg BUILD_PACKAGE=server --platform linux/amd64 --tag abashin/otus-hw-server:$(DOCKER_TAG) -f app/Dockerfile app
	docker build --build-arg BUILD_PACKAGE=migrator --platform linux/amd64 --tag abashin/otus-hw-migrator:$(DOCKER_TAG) -f app/Dockerfile app

push:
	docker push abashin/otus-hw-server:$(DOCKER_TAG)
	docker push abashin/otus-hw-migrator:$(DOCKER_TAG)

latest: build push

tag: DOCKER_TAG=$(VERSION)
tag: build push

# K8s

.PHONY: prometheus install uninstall

prometheus:
	kubectl create ns prometheus
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	helm install -n prometheus prometheus-operator prometheus-community/kube-prometheus-stack

install:
	kubectl get ns otus-hw >/dev/null 2>&1 || kubectl create ns otus-hw
	helm upgrade --install -n otus-hw otus-hw server-k8s

uninstall:
	helm status -n otus-hw otus-hw >/dev/null 2>&1 && helm uninstall -n otus-hw otus-hw || true
	kubectl get ns otus-hw >/dev/null 2>&1 && kubectl delete ns otus-hw || true
	helm status -n prometheus-hw prometheus-operator >/dev/null 2>&1 && helm uninstall -n prometheus prometheus-operator || true
	kubectl get ns prometheus >/dev/null 2>&1 && kubectl delete ns prometheus || true