# 概要
送信するファイルに任意のMIMEタイプを設定する

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
Content-Length: 294259
Content-Type: multipart/form-data; boundary=b9c81b8ee6619c3bcd48715c1747ba7628df8947ec53c62fc1dfa9c91231
User-Agent: Go-http-client/1.1

--b9c81b8ee6619c3bcd48715c1747ba7628df8947ec53c62fc1dfa9c91231
Content-Disposition: form-data; name="name"

なんとか橋
--b9c81b8ee6619c3bcd48715c1747ba7628df8947ec53c62fc1dfa9c91231
Content-Disposition: form-data; name=="thumbnail"; filename="photo.jpg"
Content-Type: image/jpeg

 acspAPPLAPPL???-appl???%M8??????0100??????4ICC_PROFILE$applmntrRGB XYZ ?
desc?ecprtd#wtpt?rXYZ?gXYZ?bXYZ?rTRC? chad?,bTRC? gTRC? desc
                                                            Display P3textCopyright Apple Inc., 2017XYZ ?Q?XYZ ??=?????XYZ J??7
?XYZ (8
Y?     ȹparaff?
[sf32
     B????&?????????????n???"??

====== 以降画像のバイナリデータが延々続く ======

```