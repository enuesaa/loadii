package cli

var HelpText = `A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags]

FLAGS:
  -go:[path] [args]        Run 'go run [path] [args]'
  -pnpm:[path]             Run 'pnpm run dev [path]'
  -pnpm:[path] [script]    Run 'pnpm run [script] [path]'
  -serve                   Serve
  -help                    Show help
  -v -version              Print the version`
