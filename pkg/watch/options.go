package watch

type Options struct {
	Includes  []string
	Excludes  []string
	Callbacks []func()
}

type Option = func(*Options)

func WithIncludes(paths []string) Option {
	return func(opts *Options) {
		opts.Includes = paths
	}
}

func WithExcludes(paths []string) Option {
	return func(opts *Options) {
		opts.Excludes = paths
	}
}

func WithCallback(fn func()) Option {
	return func(opts *Options) {
		opts.Callbacks = []func(){fn}
	}
}
