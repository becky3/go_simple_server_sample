# 概要
x-www-form-urlencoded形式のPOSTメソッドの送信

# 前提

[run_server](../../02/run_server) を立ち上げておく

# レスポンス

```
Status: 200 OK
```

# サーバー側ログ
```
POST / HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Content-Length: 10
Content-Type: application/x-www-form-urlencoded
User-Agent: Go-http-client/1.1

test=value
```