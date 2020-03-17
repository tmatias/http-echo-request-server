# http-echo-request-server

A really really simple server that echoes the request back as a JSON body in the response.

```console
$ curl localhost:8080 -H "SomeHeader: ItsValue" -d "hello, world" | jq
{
  "method": "POST",
  "url": "/",
  "headers": {
    "Accept": [
      "*/*"
    ],
    "Content-Length": [
      "12"
    ],
    "Content-Type": [
      "application/x-www-form-urlencoded"
    ],
    "Someheader": [
      "ItsValue"
    ],
    "User-Agent": [
      "curl/7.58.0"
    ]
  },
  "body": "hello, world",
  "remote_address": "172.17.0.1:42828"
}
```

## `remote_address`

This project uses the Go programming language standard library and the remote address is thus whatever the standard library sets as the value for the [net/http.Request](https://golang.org/pkg/net/http/#Request) `RemoteAddr` field.
