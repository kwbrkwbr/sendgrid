# sendgrid

## 概要
sendgridでユーザーにメール送信する

だれかが用意したpubsubにメッセージを発行

発行されたメッセージをトリガーにして自分にpushさせる

pushしてきたものでsendgridAPIを叩く

うまくいったらuserにメールが飛ぶ

## 構成
     ┌────┐          ┌──────┐          ┌──┐          ┌────────┐          ┌────┐
     │some│          │pubsub│          │ME│          │sendgrid│          │user│
     └─┬──┘          └──┬───┘          └┬─┘          └───┬────┘          └─┬──┘
       │    mail req    │               │                │                 │   
       │ ───────────────>               │                │                 │   
       │                │               │                │                 │   
       │                │  triger push  │                │                 │   
       │                │ ──────────────>                │                 │   
       │                │               │                │                 │   
       │                │               │  sendgrid API  │                 │   
       │                │               │ ───────────────>                 │   
       │                │               │                │                 │   
       │                │               │                │       mail      │   
       │                │               │                │ ────────────────>   
     ┌─┴──┐          ┌──┴───┐          ┌┴─┐          ┌───┴────┐          ┌─┴──┐
     │some│          │pubsub│          │ME│          │sendgrid│          │user│
     └────┘          └──────┘          └──┘          └────────┘          └────┘
     
AA created by [planttext.com](https://www.planttext.com/)

## 環境設定
### 環境変数
- SENDGRID_API_KEY

sendgridのAPI KEY

- SENDGRID_HOST

sendgridのAPI Endpointなホスト、初期値はコードに埋め込んでます。

- SENDGRID_ENDPOINT

sendgridのAPI Endpointなパス、初期値はコードに埋め込んでます。

- MAILER_ENV

環境変数による動作切り替え

- PORT

port指定。主にcloud runの設定に依存しているもの。

## 開発向けツール
基本的に導入は環境によって読み替えてください。

### direnv
mailer独自の環境変数の設定用。

cloud runの環境変数と同じものを用意するのに利用。

`brew direnv`

### realize
ホットデプロイツール。

`make install-realize`

### delve
デバッグツール。

`make install-delve`

### postman & newman
httpなテストツール。

postmanで作ったテストケースをnewmanでコマンド実行すると楽。

テストの実行はpostmanでしてもOK。

`brew install postman`

`brew install newman`



## ローカル開発

### 起動
realizeで起動。

起動設定は `.realize.yaml` を見てください。

### デバッグ
golandの場合、realizeで起動している `./main` にattachして使ってください。

## API リクエスト例
### /mail

```
curl --location --request POST 'http://localhost:3333/mail' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "--title--",
    "from": "postman@post.man",
    "template": "--name--様\n\nありがとうございました。\n",
    "params": "{\"name\":\"bind-name\",\"title\":\"送信成功\"}",
    "firebaseID": "secret"
}'
```
