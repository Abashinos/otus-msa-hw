# OTUS Microservice Architecture homework

Simple HTTP server

## Usage

### Install
TODO: chart version substitution
```shell
kubectl create namespace otus-hw
helm package server-k8s
helm install --namespace otus-hw otus-hw-server ./otus-hw-server-0.0.5.tgz
```

### Uninstall
```shell
helm uninstall --namespace otus-hw otus-hw-server
```

### API

:point_right: See [api.json file](./api.json) for Postman collection containing API routes described below

| Method | URL                                        | Purpose                                                                                |
|--------|--------------------------------------------|----------------------------------------------------------------------------------------|
| GET    | http://arch.homework/health                | Replies with `{"status": "ok"}`                                                        |
| GET    | http://arch.homework/hostinfo              | Replies with `{"hostname" "<POD name>"}`                                               |
| GET    | http://arch.homework/otusapp/iabashin/(.*) | Rewrites to $1                                                                         |
| POST   | http://arch.homework/users                 | Create a user. Requires JSON with (first_name, last_name) fields                       |
| GET    | http://arch.homework/users/{user_id}       | Get user with id=={user_id}                                                            |
| UPDATE | http://arch.homework/users/{user_id}       | Update user with id=={user_id}. <br/>Requires JSON with (first_name, last_name) fields |
| DELETE | http://arch.homework/users/{user_id}       | Delete user with id=={user_id}                                                         |                                                                                 |




