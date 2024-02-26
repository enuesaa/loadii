# loadii
Instant web server for development.

## Usage
```console
$ tryup --help
Instant web server

USAGE:
  tryup <path> [global options]

GLOBAL OPTIONS:
  --watch        run watch mode (default: false)
  --help, -h     show help
  --version, -v  print the version
```

## Development Plan
- [runapp] watch mode
- [serve] index
- [serve] custom logger

### Planning Usecase
```bash
loadii serve # this serve static content
loadii run main.go # this run `go run main.go` and also, do hot reload
loadii exec pnpm build
loadii run main.go | loadii -w admin exec pnpm build
```
