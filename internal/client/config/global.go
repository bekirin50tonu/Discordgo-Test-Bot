package config

import "sun_bot/pkg/config"

var (
	Config  *config.Config
	Runtime *config.Runtime
)

func New(file_name string) {
	Config = config.New(file_name)
	Runtime = config.NewRuntime()
}
