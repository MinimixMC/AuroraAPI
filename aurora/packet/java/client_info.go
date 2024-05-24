package packet

import (
	"github.com/MinimixMC/AuroraAPI/aurora/player"
	"github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"
	"github.com/rs/zerolog/log"
)

type ClientInfo struct {
	player.ClientSettings
}

/*
Displayed Skin Parts flags:

Bit 0 (0x01): Cape enabled
Bit 1 (0x02): Jacket enabled
Bit 2 (0x04): Left Sleeve enabled
Bit 3 (0x08): Right Sleeve enabled
Bit 4 (0x10): Left Pants Leg enabled
Bit 5 (0x20): Right Pants Leg enabled
Bit 6 (0x40): Hat enabled
The most significant bit (bit 7, 0x80) appears to be unused.
*/

func SkinParts(client ClientInfo) {
	skinParts := map[uint8]string{
		0x01: "Cape",
		0x02: "Jacket",
		0x04: "Left Sleeve",
		0x08: "Right Sleeve",
		0x10: "Left Pants Leg",
		0x20: "Right Pants Leg",
		0x40: "Hat",
	}

	log.Info().Msg("Displayed Skin Parts:")
	for flag, part := range skinParts {
		if client.DisplayedSkinParts&flag != 0 {
			log.Info().Msgf("%s is enabled", part)
		} else {
			log.Info().Msgf("%s is disabled", part)
		}
	}
}

func (p ClientInfo) ID() int32 {
	return 0x00
}

func (p *ClientInfo) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Locale)
	_ = r.Int8(&p.ViewDistance)
	_ = r.VarInt(&p.ChatMode)
	_ = r.Bool(&p.ChatColors)
	_ = r.Uint8(&p.DisplayedSkinParts)
	_ = r.VarInt(&p.MainHand)
	_ = r.Bool(&p.TextFilteringEnabled)
	return r.Bool(&p.AllowServerListing)
}

func (p *ClientInfo) Encode(w *encoding.Writer) error {
	_ = w.String(p.Locale)
	_ = w.Int8(p.ViewDistance)
	_ = w.VarInt(p.ChatMode)
	_ = w.Bool(p.ChatColors)
	_ = w.Uint8(p.DisplayedSkinParts)
	_ = w.VarInt(p.MainHand)
	_ = w.Bool(p.TextFilteringEnabled)
	return w.Bool(p.AllowServerListing)
}
