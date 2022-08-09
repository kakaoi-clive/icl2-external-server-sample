# iCL 2.0 External Server Sample (NodeJs)

## HOW TO RUN (default)
```bash
npm install

npm run server
````

## HOW TO RUN (docker)
```bash
docker build . -t icl2-external-server-sample-node

docker run -p 8080:8080 -d icl2-external-server-sample-node
```

## HOW TO USE

```bash
# generate token from server
curl -v localhost:8080/api/generate

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /api/generate HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200
< Content-Type: application/json
< Transfer-Encoding: chunked
< Date: Fri, 15 Jul 2022 01:21:49 GMT
<
* Connection #0 to host localhost left intact
{"token":"GENERATED_TOKEN"}

# validate token from server
curl -v localhost:8080/api/validate -H "Authorization: Bearer {TOKEN}"

curl -v localhost:8080/api/validate -H "Authorization: Bearer {GENERATED_TOKEN}"

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /api/validate HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Authorization: Bearer {GENERATED_TOKEN}
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200
< Content-Type: application/json
< Transfer-Encoding: chunked
< Date: Fri, 15 Jul 2022 01:22:06 GMT
<
* Connection #0 to host localhost left intact
{"admin":false}
```
