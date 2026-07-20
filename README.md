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
  --help, -h             show help
  --version, -v          print the version
```

## モチベーション
- 開発の際にウェブサーバをぱっと立ち上げたい時がある
- コマンド一つで立ち上げられないか

## planning features
```bash
loadii         # serve current dir. track file changes
loadii dist    # serve dist dir. track file changes
loadii --index index.html,index.md
loadii --header Cache-Control:no-store
loadii --verbose
loadii --port 3000
loadii --proxy https://example.com/
loadii -e 'path=/ proxy=https://example.com/' -e 'path=/aa header=Cache-Control:no-store'
```

例えばある特定のパスの時だけヘッダーを返すとか、柔軟な制御を入れられるようにするか。
