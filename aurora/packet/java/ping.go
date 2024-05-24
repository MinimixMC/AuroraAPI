package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type Ping struct {
	Payload int64
}

func (p Ping) ID() int32 {
	return 0x01
}

func (p *Ping) Decode(r *encoding.Reader) error {
	return r.Int64(&p.Payload)
}

func (p Ping) Encode(w *encoding.Writer) error {
	return w.Int64(p.Payload)
}
