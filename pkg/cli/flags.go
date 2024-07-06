package cli

var (
	GoFlag = Flag{
		Name: "-go",
		Help: "Run `go run`",
		MinValues: 0,
		MaxValues: 10,
		DefaultValues: []string{"."},
		DefaultWorkdir: ".",
	}
	PnpmFlag = Flag{
		Name: "-pnpm",
		Help: "Run `pnpm run`",
		MinValues: 0,
		MaxValues: 10,
		DefaultValues: []string{"dev"},
		DefaultWorkdir: ".",
	}
	ServeFlag = Flag{
		Name: "-serve",
		Help: "serve",
		MinValues: 0,
		MaxValues: 0,
		DefaultValues: []string{},
		DefaultWorkdir: ".",
	}
	HelpFlag = Flag{
		Name: "-help",
		Alias: "-h",
		Help: "Show help messages",
		MinValues: 0,
		MaxValues: 0,
		DefaultValues: []string{},
		DefaultWorkdir: "",
	}
	VersionFlag = Flag{
		Name: "-version",
		Alias: "-v",
		Help: "Print version info",
		MinValues: 0,
		MaxValues: 0,
		DefaultValues: []string{},
		DefaultWorkdir: "",
	}
)
