package main

import (
	"github.com/MinimixMC/AuroraAPI/API/aurora"
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
}

func main() {}
