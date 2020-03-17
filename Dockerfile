FROM scratch

COPY http-echo-request-server .

ENTRYPOINT ["./http-echo-request-server"]
