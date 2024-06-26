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

## Usage
```bash
loadii --serve ./dist
loadii go run .
loadii -w ./example/simple go run .
loadii --serve ./dist pnpm build
```

## モチベーション
次のような課題があり、コマンド一つで解決できないか探っている
- ウェブサーバをぱっと立ち上げたい時があるが nginx や apache のセットアップは面倒
- ファイル編集をトリガーにビルドを走らせたい時がある (HMRではない)
- Go + フロントエンドの開発をするとき `go run .` と `pnpm dev` をそれぞれ実行する必要があり面倒

## Planning Usage [Refactor]
```bash
loadii -i # interactive. scan which language is used and predict dev command
loadii -i -d 4 # scan sub-dirs with 4 depth. defalut depth is 2
loadii -go # run `go run .`
loadii -go -pnpm:ui # run `go run .` and `cd ./ui && pnpm dev`
loadii -go sub subsub -pnpm:ui
loadii -go 'sub subsub' -pnpm:ui
loadii -go 'sub subsub --flag' -pnpm:ui -i
loadii -serve # serve .
```
