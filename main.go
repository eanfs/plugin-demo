package demo

import (
	_ "embed"
	"fmt"
	"sync"

	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/config"
)

type DemoConfig struct {
	config.Subscribe
	recordings sync.Map
}

var DemoPluginConfig = &DemoConfig{}

var plugin = InstallPlugin(DemoPluginConfig)

func (conf *DemoConfig) OnEvent(event any) {
	switch v := event.(type) {
	case FirstConfig, config.Config:
		fmt.Printf(config.Config.Global)
	case SEpublish:
		streamPath := v.Target.Path
		fmt.Printf(streamPath)
	}
}
