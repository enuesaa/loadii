# tryup
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
tryup serve # this serve static content
tryup watch -e go run main.go # this run `go run main.go` and also, do hot reload
tryup watch -r main.go     # this run `go run main.go` and also, do hot reload
tryup watch -e pnpm build
tryup watch -r main.go | tryup watch -w admin -e pnpm build
```

- あまり pnpm dev で hot reload したいケースを思いつかない. そもそも next dev とかはホットリロードされるから
