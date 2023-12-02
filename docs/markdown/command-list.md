# Command List
## docker
### mysql
up
```shell
$ make docker/up -SERVER=local
```

## ent
install
```shell
$ go install entgo.io/ent/cmd/ent@latest
```

schema generate
```shell
$ ent init <任意(User)>
```

code generate
```shell
# original template
$ go run -mod=mod entgo.io/ent/cmd/ent generate --template ./ent/template --template glob="./ent/template/*.tmpl" ./ent/schema
# normal
$ go generate ./ent
```

db migrate(local)
```shell
$ make db/migrate -SERVER=local
```

# air
install
```shell
$ go get -u github.com/cosmtrek/air
```

up
```shell
$ air -c ./cmd/public/air.toml
```
