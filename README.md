## Go sandbox

### やったこと
- WSLにatlasコマンドをインストール
- 下記コマンドでマイグレーションファイルを作成

```
atlas migrate diff create_users --dir "file://db/migrations" --to "file://schema.sql" --dev-url "docker://mysql/8/some_db"
```

- 下記コマンドにて今のDBの状態と```schema.sql```との差分を見る

```
atlas schema diff --from mysql://root:password@127.0.0.1:3306/some_db --to "file://schema.sql" --format '{{ sql . " " }}' --dev-url "docker://mysql/8/some_db"
```

- 下記コマンドにて今のDBを```schema.sql```の状態に更新

```
atlas schema apply -u "mysql://root:password@127.0.0.1:3306/some_db" --to "file://schema.sql" --dev-url "docker://mysql/8/some_db"
```

atlas schema apply -u "mysql://root:password@127.0.0.1:3306/some_db" --to "file://db/migrations" --dev-url "docker://mysql/8/some_db"