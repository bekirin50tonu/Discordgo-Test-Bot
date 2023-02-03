package config

import (
	"time"

	"github.com/matishsiao/goInfo"
)

type Runtime struct {
	StartTime time.Time
	OSInfo    goInfo.GoInfoObject
}

func NewRuntime() *Runtime {
	gi, err := goInfo.GetInfo()
	if err != nil {
		panic(err)
	}
	return &Runtime{
		StartTime: time.Now(),
		OSInfo:    gi,
	}
}
