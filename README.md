# go echo onion architecture

```text
root
├── cmd: コマンドラインツール
│   ├── cli
│   │   ├── db
│   │   └── task
│   └── app
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
│   └── pkg: プロジェクトの共有コンポーネント(このプロジェクト固有)
│
├── api: openapi等
├── build: パッケージングと継続的インテグレーション(dockerfile等)
├── configs: 設定ファイル
├── docs: ドキュメント(api docは除く)
├── pkg: プロジェクトの共有コンポーネント(他のプロジェクトでも利用可)
└── util: 言語特有のutil
```

## 起動前提
go version 1.21.4
OS macOS
shell zsh
docker

## todo
- responseのファイル名の修正（modelつかわないから修正いらないかも）とentityの作成
- diについて
- entityを追加
- vendorについて
- バージョンアップ（全体的に）
- testコード
- jwtに導入→dtoが必要そうなら導入
- メール送信(パスワード忘れた機能の作成)
- 多言語化
- linter、format確認、ent実行、testのgitactionの追加
- golangをdocker化

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
- oapi-codegen
