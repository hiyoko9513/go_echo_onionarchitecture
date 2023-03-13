# go echo onion architecture

## ブランチについて
- main=minimum
- hiyoko/local=local hiyokoのメイン
- hiyoko/aws=aws hiyokoのメイン

# quick start
go 1.19.7(goinstall goroot gopathについては各自で、、、)  
todo 動作確認
```shell
$ go mod tidy
$ source ~/.zshrc
$ make docker/up -SERVER=local
$ make db/migrate -SERVER=local
$ air -c ./cmd/public/air.toml
```
http://localhost:8000/api/v1/users

# Commit Prefix
## ルール
| type                     | emoji |
|--------------------------|:-----:|
| 初めてのコミット（Initial Commit） |  🎉   |
| バージョンタグ（Version Tag）     |  🔖   |
| 新機能（New Feature）         |   ✨   |
| 機能改善                     |  🔧   |
| バグ修正（Bugfix）             |  🐛   |
| リファクタリング(Refactoring)    |  ♻️   |
| ドキュメント（Documentation）    |  📚   |
| デザインUI/UX(Accessibility) |  🎨   |
| パフォーマンス（Performance）     |  🐎   |
| 注意                       |  🚨   |
| 非推奨追加（Deprecation）       |  💩   |
| 削除（Removal）              |  🗑️  |
| WIP(Work In Progress)    |  🚧   |

## コミットテンプレートの設定方法
```shell
$ cp ./.github/.gitmessage.txt.example ./.github/.gitmessage.txt
$ git config commit.template ./.github/.gitmessage.txt
$ git config --add commit.cleanup strip
```

# todo
- アノテーションをREADMEにまとめておいておきたい https://qiita.com/taka-kawa/items/673716d77795c937d422
- seederの追加（クイックスタートにいれたい）、db truncateを追加 // あまりコードに入れるのは良くないと思うけど、開発環境で機能として欲しい
- txを追加
- エラーを返す位置について考える
- エラーメッセージの考察,slogを導入,エラーハンドラーの作成(https://zenn.dev/mizutani/articles/golang-exp-slog)
- taskコマンドを作成
- input modelとoutput modelを追加、（limitの変数化）バリデーション追加
- バージョン管理型マイグレーションを利用したい(https://entgo.io/ja/docs/versioned-migrations/)
- swaggerの追加（UIとモックの追加）
- テストコード、DBモック
- entvizの追加 https://entgo.io/ja/blog/2021/08/26/visualizing-your-data-graph-using-entviz/
- ドキュメントテンプレート作成(https://system-kanji.com/posts/system-deliverable)（自動生成出来る分については）
- gitaction https://entgo.io/ja/docs/ci
- .envの扱いについてどうしようか、、、cmd毎にわけるのは、、、プロジェクトによってかわるとは思うけれど、、、

## 問題
- sqlのdebugmodeが利用できない（https://github.com/ent/ent/issues/85）
- entの外部テンプレートについて、、、client.goを統一したい
- entの扱いについてどうしよう、、、？

## 別ブランチで追加したい(hiyoko)
- auth(jwt、originalID)機能の追加(https://iketechblog.com/go-jwt/#index_id2)
- local画像保存機能の追加
- cognito
- cloudsearchでlog、CPU使用率などの収集も
- s3ユーザー画像保存
- mail送信、ses、sns、sqs（について）

# depend on
- ent
- echo
- staticcheck
- godotenv
- air

# docker
## mysql
up
```shell
$ make docker/up -SERVER=local
```

# ent
install
```shell
$ go install entgo.io/ent/cmd/ent@latest
```
schema generate
```shell
$ ent init <任意(User)>
```
code generate
```shell
# original template
$ go run -mod=mod entgo.io/ent/cmd/ent generate --template ./ent/template --template glob="./ent/template/*.tmpl" ./ent/schema
# normal
$ go generate ./ent
```
db migrate(local)
```shell
$ make db/migrate -SERVER=local
```

# air
install
```shell
$ go get -u github.com/cosmtrek/air
```
up
```shell
$ air -c ./cmd/public/air.toml
```
