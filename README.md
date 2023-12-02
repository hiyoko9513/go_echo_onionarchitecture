# go echo onion architecture
go 1.21.4

## ブランチについて
- main=local完結 hiyokoのメイン

# docs
[quick start](/docs/markdown/quick-start.md)
[git rule](/docs/markdown/git/rule.md)
[command list](/docs/markdown/command-list.md)

## 問題
- sqlのdebugmodeが利用できない（https://github.com/ent/ent/issues/85）
- entの外部テンプレートについて、client.goに組み込みたい

# depend on
- ent
- echo
- staticcheck
- godotenv
- air
