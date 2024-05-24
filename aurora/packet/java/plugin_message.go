package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

// Serverbound
type ServerboundPluginMessage struct {
	Channel string
	Data    []byte
}

func (p ServerboundPluginMessage) ID() int32 {
	return 0x01
}

func (p *ServerboundPluginMessage) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Channel)
	return r.ByteArray(&p.Data)
}

func (p *ServerboundPluginMessage) Encode(w *encoding.Writer) error {
	_ = w.String(p.Channel)
	return w.ByteArray(p.Data)
}

//Clientbound
type ClientboundPluginMessage struct {
	Channel string
	Data    []byte
}

func (p ClientboundPluginMessage) ID() int32 {
	return 0x00
}

func (p *ClientboundPluginMessage) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Channel)
	return r.ByteArray(&p.Data)
}

func (p *ClientboundPluginMessage) Encode(w *encoding.Writer) error {
	_ = w.String(p.Channel)
	return w.ByteArray(p.Data)
}
