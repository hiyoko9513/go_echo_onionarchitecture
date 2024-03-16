# go echo onion architecture

```text
root
├── cmd: コマンドラインツール
│   ├── cli
│   │   ├── db
│   │   └── task
│   └── public
│
├── internal:
│   ├── interactor: ユースケースを操作するロジック
│   ├── presenter: presentation layer
│   ├── infrastructure: infrastructure
│   ├── application: application layer
│   │   ├── dto: data transfer object - 外部アプリまたはレイヤー間でのデータ移送のため
│   │   └── usecase: usecase
│   │
│   ├── domain: domain layer
│   │   ├── entities: entity
│   │   ├── value-objects: 値オブジェクト(不使用)
│   │   └── services: interface
│   │
│   └── pkg: プロジェクトの共有コンポーネント
│
├── api: openapi等
├── build: パッケージングと継続的インテグレーション(dockerfile等)
├── configs: 設定ファイル
├── docs: ドキュメント(api docは除く)
└── util: 言語特有のutil
```

## 起動前提
go version 1.21.4
OS macOS
shell zsh
docker

## todo
- loggerの作成(グローバルなpackageに落とし込む)
- validate input output ユーザー入力値についてプレゼンテーション層で行う
- エラーハンドリング
- diについて
- dtoのついか
- 時間のutil
- ctxにリクエストIDを導入
- domain serviceにentityを追加
- vendorについて
- バージョンアップ（全体的に）
- todoの消化
- oapi code gen
- https://github.com/hiyoko9513/go_echo_oapi_codegen
- 全てのどうかく
- golangをdocker化
- 全てのどうかく
- linter、format確認、ent実行のgitactionの追加 
- swagger(oapiのcode genに修正)
- jwtに導入
- メール送信(パスワード忘れた機能の作成)
- 多言語化
- loggerにおいて、request idの紐づけについて、request idが存在しない場合はreqIDのコメントを排除

## ツールのインストール
go tools
```shell
$ make go/install/tools
```

## ブランチについて
- main=local完結 hiyokoのメイン

## ドキュメント
- [quick start](./docs/markdown/quick-start.md)
- [git rule](./docs/markdown/git/rule.md)

## 問題
- sqlのdebug modeが利用できない（https://github.com/ent/ent/issues/85）
- entの外部テンプレートについて、client.goに組み込みたい

## depend on(go)
- ent
- echo
- staticcheck
- godotenv
- air
