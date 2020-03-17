# http-echo-request-server

A really really simple server that echoes the request back as a JSON body in the response.

```console
$ curl -i localhost:8080 -H "SomeHeader: ItsValue"
HTTP/1.1 200 OK
Date: Tue, 17 Mar 2020 14:32:12 GMT
Content-Length: 86
Content-Type: text/plain; charset=utf-8

{"headers":{"Accept":["*/*"],"Someheader":["ItsValue"],"User-Agent":["curl/7.58.0"]}}
```
