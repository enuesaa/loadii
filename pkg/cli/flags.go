package cli

var (
	GoFlag = Flag{
		Name: "-go",
		Help: "Run `go run`",
		MinValues: 0,
		MaxValues: 10,
		DefaultValues: []string{"."},
		Workdir: ".",
		ReceiveWorkdir: true,
	}
	PnpmFlag = Flag{
		Name: "-pnpm",
		Help: "Run `pnpm run`",
		MinValues: 0,
		MaxValues: 10,
		DefaultValues: []string{"dev"},
		Workdir: ".",
		ReceiveWorkdir: true,
	}
	HelpFlag = Flag{
		Name: "-help",
		Help: "Show help messages",
		MinValues: 0,
		MaxValues: 0,
		DefaultValues: []string{},
		Workdir: "",
		ReceiveWorkdir: false,
	}
	VersionFlag = Flag{
		Name: "-version",
		Help: "Print version info",
		MinValues: 0,
		MaxValues: 0,
		DefaultValues: []string{},
		Workdir: "",
		ReceiveWorkdir: false,
	}
)
