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
  -serve:[path]            Serve
  -h -help                 Show help
  -v -version              Print the version
```

## モチベーション
次のような課題があり、コマンド一つで解決できないか探っている
- ウェブサーバをぱっと立ち上げたい時があるが nginx や apache のセットアップは面倒
- ファイル編集をトリガーにビルドを走らせたい時がある (HMRではない)
- Go + フロントエンドの開発をするとき `go run .` と `pnpm dev` をそれぞれ実行する必要がある

## Planning Usage [Experimental]
現実問題としてWebサーバをぱっと立ち上げたいので、それに集中する
