package player

import "github.com/MinimixMC/AuroraAPI/aurora/data/types"

type Player struct {
	Name       string
	UUID       [16]byte
	Properties []types.Property

	Settings *ClientSettings
}
