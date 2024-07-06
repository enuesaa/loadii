# loadii
A CLI tool to watch file changes and execute a command

## Usage
```console
$ loadii --help
A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags]

FLAGS:
  -go:[path] [args]        Run 'go run [path] [args]'
  -pnpm:[path]             Run 'pnpm run dev [path]'
  -pnpm:[path] [script]    Run 'pnpm run [script] [path]'
  -serve                   Serve
  -help                    Show help
  -v -version              Print the version
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
