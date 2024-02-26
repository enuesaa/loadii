# loadii
Instant web server for development.

## Development Plan
- [runapp] watch mode
- [serve] index
- [serve] custom logger

### Planning Usecase
```bash
loadii serve # this serve static content
loadii run main.go # this run `go run main.go` and also, do hot reload
loadii exec pnpm build
loadii serve | loadii exec pnpm build
loadii run main.go | loadii -w admin exec pnpm build
```

## モチベーション
- 開発の際にウェブサーバをぱっと立ち上げたい時があるが nginx や apache のセットアップは面倒
- SSG方式のサイトでコンテンツを更新したときに即座にビルド処理を走らせたい時がある (hot reload)
- 上記2つは本来無関係だが使用するタイミングが近く、一緒のコマンドになっていると便利そう
