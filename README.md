# bot
an MRKL implementation using gofiber


## how to add swagger
```bash
$ go get github.com/gofiber/contrib/swagger  
```

```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
```

```bash
$ swag init -g cmd/server/main.go --parseInternal
```
