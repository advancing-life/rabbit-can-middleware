# Middleware

## aha!
1. jQueryがEchoにRequest
2. Echoが開いてるIDを探しIDをRedixに保存
3. EchoがDockerコンテナを作成しコンテナIDをRedixに保存
4. EchoがIDでWS通信をデキるようにしてjQueryにレスポンス
5. jQueryがコードを送信
6. EchoがIDに紐付いてるDokcerコンテナでコードを実行
7. Echoがjqueryに返却


![image1](https://raw.githubusercontent.com/advancing-life/rabbit-can-middleware/develop/.images/Middleware-1.jpg)
![image2](https://raw.githubusercontent.com/advancing-life/rabbit-can-middleware/develop/.images/Middleware-2.jpg)

## /api/v1/ (GET)
[SampleURL](http://localhost:1234/api/v1)
> サーバーとの接続確認に使う

### Responce

~~~json
{
    "status": 200,
    "message": "OK"
}
~~~

|Key|Model|Value|
|:--|:--|:--|
|status|integer|200|
|message|string|OK|

## /api/v1/connection (GET)
[SampleURL](http://localhost:1234/api/v1/connection)
> WebSocet用のコネクションURLを要求及び発行

### Responce

~~~json
~~~
