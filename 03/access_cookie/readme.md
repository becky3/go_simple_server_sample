# 概要
クッキーの送受信

# 前提

[run_server](../../02/run_server) を立ち上げておく

# レスポンス

## 一回目
```
HTTP/1.1 200 OK

Content-Length: 44

Content-Type: text/html; charset=utf-8

Date: Sun, 14 Jun 2020 00:31:05 GMT

Set-Cookie: VISIT=TRUE



<html><body>
<p>初訪問</p></body></html>

```

## 二回目
```
HTTP/1.1 200 OK

Content-Length: 51

Content-Type: text/html; charset=utf-8

Date: Sun, 14 Jun 2020 00:31:05 GMT

Set-Cookie: VISIT=TRUE



<html><body>
<p>二回目以降</p>
</body></html>
```

# サーバー側ログ
```
GET /cookie HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
User-Agent: Go-http-client/1.1


GET /cookie HTTP/1.1
Host: localhost:18888
Accept-Encoding: gzip
Cookie: VISIT=TRUE
User-Agent: Go-http-client/1.1
```