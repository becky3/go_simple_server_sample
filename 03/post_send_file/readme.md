# 概要
POSTでファイル送信

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
Transfer-Encoding: chunked
Accept-Encoding: gzip
Content-Type: text/plain
User-Agent: Go-http-client/1.1

7c
ファイル送信テストだよ
😀 😃 😄 😁 😆 😅 😂 🤣 ☺️ 😊 
ABCDEFGHIJKLMNOPQRSTUVWXYZ
0123456789
0
```