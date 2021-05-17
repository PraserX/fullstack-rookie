package database

// Options are used for User constructor function.
type Options struct {
	Path string
}

// Option definition.
type Option func(*Options)

// OptionPath option specification.
func OptionPath(path string) Option {
	return func(opts *Options) {
		opts.Path = path
	}
}
