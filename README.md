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

## 前提
go version 1.22.1  
OS macOS  
shell zsh  
docker  

## todo
- jwtに導入→dtoが必要そうなら導入
- oapiのvalidate makefileを追加
- entityの作成+dto作成
- トランザクションの考慮
- postmanのドキュメント化について
- vendorについて
- メール送信(パスワード忘れた機能の作成)
- 多言語化
- testコード
- linter、format確認、ent実行、testのgitactionの追加
- golangをdocker化

## ツールのインストール
go tools
```shell
$ make go/install/tools
```

## ドキュメント
- [quick start](./docs/markdown/quick-start.md)
- [git rule](./docs/markdown/git/rule.md)

## depend on(go)
- ent
- echo
- staticcheck
- godotenv
- air

## ポートが起動したままになってしまった場合
8000 port kill
```shell
$ lsof -i :8000  | tail -n 1 | awk '{print $2}' | xargs kill -9
```
