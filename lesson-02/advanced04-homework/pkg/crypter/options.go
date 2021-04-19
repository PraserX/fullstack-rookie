package crypter

// Options are used for Crypter constructor function.
type Options struct {
	File string
}

// Option definition.
type Option func(*Options)

// OptionFile option specification.
func OptionFile(file string) Option {
	return func(opts *Options) {
		opts.File = file
	}
}
