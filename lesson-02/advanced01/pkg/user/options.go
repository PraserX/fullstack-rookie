package user

// Options are used for User constructor function.
type Options struct {
	Username  string
	FirstName string
	LastName  string
}

// Option definition.
type Option func(*Options)

// OptionUsername option specification.
func OptionUsername(username string) Option {
	return func(opts *Options) {
		opts.Username = username
	}
}

// OptionFirstName option specification.
func OptionFirstName(firstname string) Option {
	return func(opts *Options) {
		opts.FirstName = firstname
	}
}

// OptionLastName option specification.
func OptionLastName(lastname string) Option {
	return func(opts *Options) {
		opts.LastName = lastname
	}
}
