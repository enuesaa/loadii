# tryserve
Instant web server for development.

## Usage
```console
$ tryserve --help
Instant web server

USAGE:
  tryserve <path> [global options]

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
tryserve .       # this serve static content
tryserve main.go # this run `go run main.go` internally
tryserve main.go --watch . # this run `go run main.go` and also, do hot reload

# or 
tryserve go run main.go
tryserve pnpm dev
```
