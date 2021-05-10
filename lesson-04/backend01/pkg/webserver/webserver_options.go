package webserver

import (
	"github.com/PraserX/fullstack-rookie/pkg/database"
)

// Options are used for User constructor function.
type Options struct {
	Host     string
	Port     string
	Content  string
	DevMode  bool
	Database *database.Database
}

// Option definition.
type Option func(*Options)

// OptionHost option specification.
func OptionHost(host string) Option {
	return func(opts *Options) {
		opts.Host = host
	}
}

// OptionPort option specification.
func OptionPort(port string) Option {
	return func(opts *Options) {
		opts.Port = port
	}
}

// OptionContent option specification.
func OptionContent(content string) Option {
	return func(opts *Options) {
		opts.Content = content
	}
}

// OptionDevMode option specification.
func OptionDevMode(devmode bool) Option {
	return func(opts *Options) {
		opts.DevMode = devmode
	}
}
