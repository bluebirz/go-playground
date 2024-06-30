# test-http

## cmd

- run in background

```sh
go run server.go &
```

- sample curl

```sh
curl localhost:8010/test
```

- clear port

```sh
sudo lsof -iTCP -sTCP:LISTEN -n -P
kill -9 <pid>
```
