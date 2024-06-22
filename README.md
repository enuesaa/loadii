# loadii
A CLI tool to watch file changes and execute a command

## Usage
```console
$ loadii --help
A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags] commands

FLAGS:
  --help, -h     show help
  --version, -v  print the version
  --yes, -y      Approve command execution (default: false)

  SERVE
  --port value               Serve port (default: 3000)
  --serve value              Serve dir
  --workdir value, -w value  Command execution dir (default: ".")

  WATCH
  --exclude value [ --exclude value ]  Remove path to watch
  --include value [ --include value ]  Add path to watch (default: ".")

```

## Planning Usage
```bash
loadii --serve ./dist
loadii go run .
loadii -w ./example/simple go run .
loadii --serve ./dist pnpm build
```

## モチベーション
- 開発の際にウェブサーバをぱっと立ち上げたい時があるが nginx や apache のセットアップは面倒
- SSG方式のサイトでコンテンツを更新したときに即座にビルド処理を走らせたい時がある
- 上記2つは本来無関係だが使用するタイミングが近く、一緒のコマンドになっていると便利そう

## Development Plan
- 現状、Go + frontend の開発をするとき go run . と pnpm dev の両方のコマンドを別タブで叩く必要がある
- フラグの定義は雑でもいいのでとにかく早めに解消させたい
