# quick start
(goinstall goroot gopath dockerについては省略)
```shell
$ go mod tidy
$ source ~/.zshrc
$ make docker/up
$ make db/migrate
$ make db/seed
$ make server/run-air
```
http://localhost:8000/api/v1/users
