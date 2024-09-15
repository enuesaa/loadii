package cli

//TODO: build help text from flag command.
var helpText = `A CLI tool to watch file changes and execute a command

USAGE:
  loadii [flags]

FLAGS:
  -go:[path] [args]        Run 'go run [path] [args]'
  -pnpm:[path]             Run 'pnpm run dev [path]'
  -pnpm:[path] [script]    Run 'pnpm run [script] [path]'
  -serve:[path]            Serve
  -h -help                 Show help
  -v -version              Print the version`
