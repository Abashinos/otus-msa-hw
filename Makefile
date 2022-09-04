# Docker

.PHONY: build push latest tag

DOCKER_TAG=latest

build:
	@docker build --build-arg BUILD_PACKAGE=server --platform linux/amd64 --tag abashin/otus-hw-server:$(DOCKER_TAG) -f app/Dockerfile app
	@docker build --build-arg BUILD_PACKAGE=migrator --platform linux/amd64 --tag abashin/otus-hw-migrator:$(DOCKER_TAG) -f app/Dockerfile app

push:
	@docker push abashin/otus-hw-server:$(DOCKER_TAG)
	@docker push abashin/otus-hw-migrator:$(DOCKER_TAG)

latest: build push

tag: DOCKER_TAG=$(VERSION)
tag: build push

# K8s

.PHONY: install_prometheus uninstall_prometheus install_ingress uninstall_ingress install_app uninstall_app install uninstall clean

install_prometheus:
	@helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	@helm repo update
	@helm upgrade --install -n prometheus prometheus-operator prometheus-community/kube-prometheus-stack -f k8s/prom/values.yaml --create-namespace

uninstall_prometheus:
	@helm status -n prometheus prometheus-operator >/dev/null 2>&1 && helm uninstall -n prometheus prometheus-operator || echo "prometheus operator not installed"
	@kubectl get ns prometheus >/dev/null 2>&1 && kubectl delete ns prometheus || true

install_ingress:
	@helm upgrade --install -n ingress-nginx ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx -f k8s/ingress/values.yaml --create-namespace

uninstall_ingress:
	@helm status -n ingress-nginx ingress-nginx >/dev/null 2>&1 && helm uninstall -n ingress-nginx ingress-nginx || echo "NGINX ingress controller not installed"
	@kubectl get ns ingress-nginx >/dev/null 2>&1 && kubectl delete ns ingress-nginx || true

install_app:
	@helm upgrade --install -n otus-hw otus-hw k8s/app --create-namespace

uninstall_app:
	@helm status -n otus-hw otus-hw >/dev/null 2>&1 && helm uninstall -n otus-hw otus-hw || echo "otus-hw app not installed"
	@kubectl get ns otus-hw >/dev/null 2>&1 && kubectl delete ns otus-hw || true

install: install_ingress install_prometheus install_app

uninstall: uninstall_ingress uninstall_prometheus uninstall_app

clean: uninstall