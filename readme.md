# Skald

A personal learinig project written in golang

"A Microservice enabling updating and editing Adaptive Story Engine books."

## Building the software

```bash
rm go.mod go.sum && go mod init <your repository>
```

## Runnig the service

```bash
./cmd/skald
```

## Using the service

Create a book by using a PUT method

```http
PUT /book HTTP/1.1
Host: localhost:8080
Content-Type: application/json
User-Agent: PostmanRuntime/7.15.0
Accept: */*
Cache-Control: no-cache
Postman-Token: 92e1ec3b-dee7-4cd6-b8a3-1ad0d3e65a43,85f322b8-80e0-4b22-abe4-51ddab18475c
Host: localhost:8080
accept-encoding: gzip, deflate
content-length: 40
Connection: keep-alive
cache-control: no-cache

{
	"title":"a test",
	"creator":"Nill"
}
```
