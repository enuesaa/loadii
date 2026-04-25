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

## feature flags
- verbose request header
- otel
