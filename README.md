# go echo onion architecture
## 前提
Version 1.21.4
OS macOS
shell zsh
docker

## ブランチについて
- main=local完結 hiyokoのメイン

## docs
- [quick start](./docs/markdown/quick-start.md)
- [folder struct](./docs/markdown/folder-struct.md)
- [git rule](./docs/markdown/git/rule.md)

## 問題
- sqlのdebugmodeが利用できない（https://github.com/ent/ent/issues/85）
- entの外部テンプレートについて、client.goに組み込みたい

## depend on(go)
- ent
- echo
- staticcheck
- godotenv
- air
