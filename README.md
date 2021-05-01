# go-grpc-sample
goでgRPCサーバを立てるときのサンプルです。  
vsCodeのremote-containerで構築することを前提としています。  

# develop
[remote-container](https://code.visualstudio.com/docs/remote/containers)で構築しています。  
詳しい説明は省きます。  
remote-containerで立ち上げると、同時にpostgreSQLも立ち上がります。  
また、migration用ツールの[golang-migrate](https://github.com/golang-migrate/migrate)もインストール済みの環境です。

## gRPC
9000portでgRPCサーバが立ちます。  
[evans](https://github.com/ktr0731/evans)などのgRPCクライアントを用いて接続確認してください。  
protoファイルからGoのコードは`go generate`で生成できます。

```sh
evans --proto="./api/sample.proto" --host="localhost" --port="9000" repl
```

## migrate
DBのスキーマ変更には[golang-migrate](https://github.com/golang-migrate/migrate)を使用している。

```sh
# DBのver up
migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable" -path ./migrations up

# DBのver down
migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable" -path ./migrations down
```

## sqlboiler
DBに接続してSchemaからstructを生成します。  
`go generate`を実行してください。`db/generated`以下に生成されます。
