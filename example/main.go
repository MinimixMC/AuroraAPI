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
	Init: Init,
}

func Init(p *aurora.Plugin) {
	log.Info().Msgf("%s loaded", p.Name)
	// Plugin logic here
}

func main() {}
