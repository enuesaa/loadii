# loadii
Instant web server for development purposes.

## Usage
```bash
$ loadii --help
Instant web server for development

USAGE:
  loadii [global options] [arguments...]

FLAGS:
  --port value           Serve port (default: 3000)
  --dir value, -d value  Serve dir (default: ".")
  --help, -h             show help
  --version, -v          print the version
```

## モチベーション
- 開発の際にウェブサーバをぱっと立ち上げたい時がある
- コマンド一つで立ち上げられないか

## planning features
```bash
loadii            # serve current dir. track file changes
loadii dist       # serve dist dir. track file changes
loadii --no-track # serve current dir. no track
loadii --index index.html,index.md
loadii --no-index # serve current dir. no index
loadii --header Cache-Control:no-store
loadii --verbose
loadii --port 3000
loadii --demo     # serve demo app
loadii --proxy /a:https://example.com/
```

例えばある特定のパスの時だけヘッダーを返すとか、柔軟な制御を入れられるようにするか。

## planning tech stacks
- go uber fx
