# OTUS Microservice Architecture homework

## Questions
- [ ] [server-service-monitor](k8s/app/templates/server-service-monitor.yaml) - move to prometheus namespace?
- [ ] Разделение лейблов app, db и тд - норм? 
Simple HTTP server

## Usage

### Install

```shell
make install
```

### Uninstall
```shell
make uninstall
```

### API

:point_right: See [api.json file](./api.json) for Postman collection containing API routes described below

| Method | URL                                        | Purpose                                                                                |
|--------|--------------------------------------------|----------------------------------------------------------------------------------------|
| GET    | http://arch.homework/otusapp/iabashin/(.*) | Rewrites to $1                                                                         |
| GET    | http://arch.homework/hostinfo              | Replies with `{"hostname" "<POD name>"}`                                               |
| GET    | http://arch.homework/health                | Replies with `{"status": "ok"}`                                                        |
| GET    | http://arch.homework/metrics               | Exposed prometheus metrics                                                             |
| GET    | http://arch.homework/prometheus            | Prometheus UI                                                                          |
| GET    | http://arch.homework/grafana               | Grafana UI                                                                             |
| POST   | http://arch.homework/users                 | Create a user. Requires JSON with (first_name, last_name) fields                       |
| GET    | http://arch.homework/users/{user_id}       | Get user with id=={user_id}                                                            |
| UPDATE | http://arch.homework/users/{user_id}       | Update user with id=={user_id}. <br/>Requires JSON with (first_name, last_name) fields |
| DELETE | http://arch.homework/users/{user_id}       | Delete user with id=={user_id}                                                         |                                                                                 |




