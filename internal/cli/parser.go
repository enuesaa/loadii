package cli

var Args = make([]string, 0)

func Parse(args []string) {
	Args = args
}
