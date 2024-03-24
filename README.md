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

  [serve]
  --port value               Serve port (default: 3000)
  --serve value              Serve dir
  --workdir value, -w value  Command execution dir (default: ".")

  [watch]
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
