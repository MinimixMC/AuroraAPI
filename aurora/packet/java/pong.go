package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type Pong struct {
	Payload int64
}

func (p Pong) ID() int32 {
	return 0x01
}

func (p *Pong) Decode(r *encoding.Reader) error {
	return r.Int64(&p.Payload)
}

func (p Pong) Encode(w *encoding.Writer) error {
	return w.Int64(p.Payload)
}
