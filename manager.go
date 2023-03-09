package internalcmd

import "github.com/sandwich-go/boost/singleflight"

var GlobalManager *Manager

type Manager struct {
	cc    *Options
	fight *singleflight.Group
}

func New(opts ...Option) *Manager {
	cfg := NewOptions(opts...)
	GlobalManager = &Manager{cc: cfg, fight: &singleflight.Group{}}
	return GlobalManager
}

func GetOptions() *Options {
	return GlobalManager.cc
}

func GetCheckUpFlight() *singleflight.Group {
	return GlobalManager.fight
}
