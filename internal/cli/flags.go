package cli

var (
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
