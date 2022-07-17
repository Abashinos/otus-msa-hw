# OTUS Microservice Architecture homework

Simple HTTP server

| URL                                        | Purpose                                  |
|--------------------------------------------|------------------------------------------|
| http://arch.homework/health                | Replies with `{"status": "ok"}`          |
| http://arch.homework/hostinfo              | Replies with `{"hostname" "<POD name>"}` |
| http://arch.homework/otusapp/iabashin/(.*) | Rewrites to $1                           |

## Usage

### Install
TODO: chart version substitution
```shell
kubectl create namespace otus-hw
helm package server-k8s
helm install --namespace otus-hw otus-hw-server ./otus-hw-server-0.0.1.tgz
```

### Uninstall
```shell
helm uninstall --namespace otus-hw otus-hw-server
```