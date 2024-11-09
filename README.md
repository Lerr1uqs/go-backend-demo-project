
```shell
go get -u github.com/gorilla/mux
go get github.com/mattn/go-sqlite3
go get github.com/google/uuid
```

```shell
go run .
```
直接运行

curl
```shell
~ curl -X POST http://localhost:8080/users \
      -H "Content-Type: application/json" \
      -d '{"username":"testuser", "password":"testpass"}' -v
```