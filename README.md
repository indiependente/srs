# Slow Request Simulator (SRS)

## Server
`./srs 5`

```bash
    ./srs n

        n must be an integer representing the number of seconds before the server times out
```

## Client
`curl localhost:8090/hello\?t=6`

```bash
    curl localhost:8090/hello\?t=6
    
        t must be an integer representing the number of seconds the request will take to complete
```

## Example using http CLI

 - Happy path (request takes less or as much as timeout allows)
```bash
http :8090/hello\?t=5
HTTP/1.1 200 OK
Content-Length: 9
Content-Type: text/plain; charset=utf-8
Date: Thu, 01 Oct 2020 14:55:36 GMT

all good
```

 - Sad path (request takes longer than timeout allows)
```bash
http :8090/hello\?t=6
HTTP/1.1 500 Internal Server Error
Content-Length: 26
Content-Type: text/plain; charset=utf-8
Date: Thu, 01 Oct 2020 15:03:32 GMT
X-Content-Type-Options: nosniff

context deadline exceeded

```
