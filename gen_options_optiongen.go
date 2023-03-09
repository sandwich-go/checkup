// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package internalcmd

import (
	"context"

	"github.com/sandwich-go/internalcmd/protocol/netutils"
)

// Options should use NewOptions to initialize it
type Options struct {
	DevopsCheckup func(ctx context.Context) *netutils.CmdCheckup
	// Stream
	IStream     IStream
	ISteamCache IStreamCache
}

// NewOptions new Options
func NewOptions(opts ...Option) *Options {
	cc := newDefaultOptions()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogOptions != nil {
		watchDogOptions(cc)
	}
	return cc
}

// ApplyOption apply multiple new option
func (cc *Options) ApplyOption(opts ...Option) {
	for _, opt := range opts {
		opt(cc)
	}
}

// Option option func
type Option func(cc *Options)

// WithDevopsCheckup option func for filed DevopsCheckup
func WithDevopsCheckup(v func(ctx context.Context) *netutils.CmdCheckup) Option {
	return func(cc *Options) {
		cc.DevopsCheckup = v
	}
}

// WithIStream option func for filed IStream
func WithIStream(v IStream) Option {
	return func(cc *Options) {
		cc.IStream = v
	}
}

// WithISteamCache option func for filed ISteamCache
func WithISteamCache(v IStreamCache) Option {
	return func(cc *Options) {
		cc.ISteamCache = v
	}
}

// InstallOptionsWatchDog the installed func will called when NewOptions  called
func InstallOptionsWatchDog(dog func(cc *Options)) { watchDogOptions = dog }

// watchDogOptions global watch dog
var watchDogOptions func(cc *Options)

// newDefaultOptions new default Options
func newDefaultOptions() *Options {
	cc := &Options{}

	for _, opt := range [...]Option{
		WithDevopsCheckup(func(ctx context.Context) *netutils.CmdCheckup {
			return &netutils.CmdCheckup{Code: netutils.ErrorCode_OK.NumberInt32(), Message: "default ok"}
		}),
		WithIStream(nil),
		WithISteamCache(nil),
	} {
		opt(cc)
	}

	return cc
}

// all getter func
func (cc *Options) GetDevopsCheckup() func(ctx context.Context) *netutils.CmdCheckup {
	return cc.DevopsCheckup
}
func (cc *Options) GetIStream() IStream          { return cc.IStream }
func (cc *Options) GetISteamCache() IStreamCache { return cc.ISteamCache }

// OptionsVisitor visitor interface for Options
type OptionsVisitor interface {
	GetDevopsCheckup() func(ctx context.Context) *netutils.CmdCheckup
	GetIStream() IStream
	GetISteamCache() IStreamCache
}

// OptionsInterface visitor + ApplyOption interface for Options
type OptionsInterface interface {
	OptionsVisitor
	ApplyOption(...Option)
}
