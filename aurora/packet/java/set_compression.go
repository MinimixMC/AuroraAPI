package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type SetCompression struct {
	Threshold int32
}

func (p SetCompression) ID() int32 {
	return 0x03
}

func (p *SetCompression) Decode(r *encoding.Reader) error {
	return r.VarInt(&p.Threshold)
}

func (p *SetCompression) Encode(w *encoding.Writer) error {
	return w.VarInt(p.Threshold)
}
