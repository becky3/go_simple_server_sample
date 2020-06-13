# 概要
multipart/form-dataでファイルの送信

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
Content-Length: 294272
Content-Type: multipart/form-data; boundary=ed15f496ac5a05bda8c03b7342d9f4118f905573e22bae7de61b221cd869
User-Agent: Go-http-client/1.1

--ed15f496ac5a05bda8c03b7342d9f4118f905573e22bae7de61b221cd869
Content-Disposition: form-data; name="name"

なんとか橋
--ed15f496ac5a05bda8c03b7342d9f4118f905573e22bae7de61b221cd869
Content-Disposition: form-data; name="thumbnail"; filename="photo.jpg"
Content-Type: application/octet-stream

 acspAPPLAPPL???-appl???%M8??????0100??????4ICC_PROFILE$applmntrRGB XYZ ?
desc?ecprtd#wtpt?rXYZ?gXYZ?bXYZ?rTRC? chad?,bTRC? gTRC? desc
                                                            Display P3textCopyright Apple Inc., 2017XYZ ?Q?XYZ ??=?????XYZ J??7
?XYZ (8
Y?     ȹparaff?
[sf32
     B????&?????????????n???"??

====== 以降画像のバイナリデータが延々続く ======

```