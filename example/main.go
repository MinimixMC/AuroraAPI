package main

import (
	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/rs/zerolog/log"
)

var Plugin aurora.Plugin = aurora.Plugin{
	Name: "Example Plugin",
	Author: []string{
		"I'm Carsen",
	},
	Version: aurora.PluginVersion{
		Version: "v0.0.1",
		VersionSupport: map[int]bool{
			1: true,
		},
	},
	MainFunc: MainFunc,
}

func MainFunc(p *aurora.Plugin) {
	log.Info().Msgf("%s loaded: %s", p.Name, p.Version.Version)
}

func GetPlugin() *aurora.Plugin {
	return &Plugin
}

func main() {}
