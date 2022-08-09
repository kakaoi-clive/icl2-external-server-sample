# iCL 2.0 External Server Sample (Golang)

## HOW TO RUN (default)
```bash
go run cmd/external-server-sample/main.go
````

## HOW TO RUN (docker)
```bash
docker build -t icl2-external-server-sample .

docker run -p 8080:8080 icl2-external-server-sample
```

## HOW TO USE

```bash
# generate token from server

$ curl -v localhost:8080/api/generate
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /api/generate HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 14 Jul 2022 08:08:27 GMT
< Content-Length: 203
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
{"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3ODk3MDcsImlhdCI6MTY1Nzc4NjEwNywidXVpZCI6IjRjOThiY2M2LTIwZGItNDQ0Yi1hNjc0LTQ3OTBhZGIwMWZkZiJ9.6g_8YdHyYJsEDs0NgscLCgg1idAZTz0dd7xwH6WSGoY"}%

# validate token from server

$ curl localhost:8080/api/validate -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3ODk3MDcsImlhdCI6MTY1Nzc4NjEwNywidXVpZCI6IjRjOThiY2M2LTIwZGItNDQ0Yi1hNjc0LTQ3OTBhZGIwMWZkZiJ9.6g_8YdHyYJsEDs0NgscLCgg1idAZTz0dd7xwH6WSGoY"

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /api/validate HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3ODk3MDcsImlhdCI6MTY1Nzc4NjEwNywidXVpZCI6IjRjOThiY2M2LTIwZGItNDQ0Yi1hNjc0LTQ3OTBhZGIwMWZkZiJ9.6g_8YdHyYJsEDs0NgscLCgg1idAZTz0dd7xwH6WSGoY
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 14 Jul 2022 08:11:54 GMT
< Content-Length: 17
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host localhost left intact
{"admin":false}

# example error case

$ curl -v localhost:8080/api/generate -X POST
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /api/generate HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 405 Method Not Allowed
< Date: Thu, 14 Jul 2022 08:20:23 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```
