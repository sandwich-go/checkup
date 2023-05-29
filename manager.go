package checkup

import (
	"github.com/sandwich-go/boost/singleflight"
	"sync"
)

var ManagerMap sync.Map

var GlobalManager *Manager
var GlobalManagerOnce sync.Once

type Manager struct {
	cc    *Options
	fight *singleflight.Group
}

func New(opts ...Option) *Manager {
	GlobalManagerOnce.Do(func() {
		cfg := NewOptions(opts...)
		GlobalManager = &Manager{cc: cfg, fight: &singleflight.Group{}}
	})
	return GlobalManager
}

func ApplyOption(opts ...Option) {
	GlobalManager.cc.ApplyOption(opts...)
}

func GetOptions() *Options {
	return GlobalManager.cc
}

func GetCheckUpFlight() *singleflight.Group {
	return GlobalManager.fight
}
