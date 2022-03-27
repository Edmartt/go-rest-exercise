# Simple web service example

A simple exercise for learning about http, rest api in Go language.

## Steps for start this project

```
git clone https://github.com/edmartt/go-rest-excercise
```

```
go get github.com/gorilla/mux
```

```
go run main.go
```

## Request examples

- Get all articles

```
curl -i -H "Content-Type: application/json" http://localhost:8081/articles 
```

- Get a single article

```
curl -i -H "Content-Type: application/json" http://localhost:8081/article/1
```

- Create new article

```
curl -i -X POST -H "Content-Type: application/json" -d '{"id":"3", "title": "My new title", "desc": "My new desc", "content": "My new content"}' http://localhost:8081/article
```

- Update an article

```
 curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"alice", "desc": "fantasy book", "content": "ilustrations"}' http://localhost:8081/article/2
```

- Delete an article

```
curl -i -X DELETE -H "Content-Type: application/json" http://localhost:8081/article/1
```