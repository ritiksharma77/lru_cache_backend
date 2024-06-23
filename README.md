This project is based on GoLang, and it implements the LRU Cache Management System.

# How to run.

```
    go mod init lru-cache-backend
    go mod tidy
    go run main.go handler.go cache.go
```

## main.go

    This file contains the entry of application and the represents the handler.

## handler.go

    This file contains the defination of the handlers and process the requests.

## cache.go

    This file contains the all the logic and the implementation of the LRU cache.

## Curl Commands

Get API:- \
`curl -X GET 'http://localhost:8080/get?key=mykey'`

Set API:- \
`curl -X POST 'http://localhost:8080/set?key=mykey&value=myvalue&ttl=10'`

Delete API:- \
`curl -X DELETE 'http://localhost:8080/delete?key=mykey'`
