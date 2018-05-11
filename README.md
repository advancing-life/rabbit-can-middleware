# Middleware

## aha!
1. jQueryがEchoにRequest
2. Echoが開いてるIDを探しIDをRedixに保存
3. EchoがDockerコンテナを作成しコンテナIDをRedixに保存
4. EchoがIDでWS通信をデキるようにしてjQueryにレスポンス
5. jQueryがコードを送信
6. EchoがIDに紐付いてるDokcerコンテナでコードを実行
7. Echoがjqueryに返却


![image1](https://github.com/advancing-life/rabbit-can-middleware/blob/master/.images/Middleware-1.jpg?raw=true)
![image2](https://github.com/advancing-life/rabbit-can-middleware/blob/master/.images/Middleware-2.jpg?raw=true)

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
|status|integer|http status code|
|message|string|サーバーからの歓迎の言葉|

## /api/v1/connection/:lang (GET)
[SampleURL](http://localhost:1234/api/v1/connection/rb)
> WebSocet用のコネクションURLを要求及び発行

### Shot cut for parameter meter

|lang  |short |
|:-----|:-----|
|Ruby  | rb   |
|Java  | java |
|Clang | c    |
|Python| py   |

### Responce

~~~json
{
    "url": "ws://localhost:1234/api/v1/execution_environment/ce16824e6180167ef65b1803c6b21b5d",
    "container_id": "ce16824e6180167ef65b1803c6b21b5d",
    "result": "99f7f325eea9e42d7c494e6fc9a69e778b5a071dbba5504ebc6072147a8a9323 "
}
~~~

|Key|Model|Value|
|:--|:--|:--|
|url|string|WS通信用のURL|
|container_id|string|作成されたDockerContainerのName|
|result|string|作成されたDockerContainerのUuid|


## /api/v1/execution_environment/:name (WebSocket)
> WebSocketでDokcerContainer内を操作

### Request

~~~json
{
    "container_id": "11aa1d97274d794a90fe10c32ba828de",
    "command": "ls",
}
~~~

|Key|Model|Value|
|:--|:--|:--|
|container_id|string|DockerContainerのName|
|command|string|実行したいコマンド|

### Responce

~~~json
{
    "container_id": "11aa1d97274d794a90fe10c32ba828de",
    "command": "ls",
    "result":"bin/nboot\ndev\netc\nhome\nlib\nlib64\nmedia\nmnt\nopt\nproc\nroot\nrun\nsbin\nsrv\nsys\ntmp\nusr\nvar",
    "exit_status": 0,
}
~~~

|Key|Model|Value|
|:--|:--|:--|
|container_id|string|DockerContainerのName|
|command|string|実行したコマンド|
|result|string|実行結果|
|exit_status|int|終了時のExitStatus|


