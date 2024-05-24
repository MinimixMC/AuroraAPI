package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type LoginStart struct {
	Name string
	UUID [16]byte
}

func (p LoginStart) ID() int32 {
	return 0x00

}

func (p *LoginStart) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Name)
	// If server is not in online mode (probably)
	var hasUUID bool
	if _ = r.Bool(&hasUUID); hasUUID {
		return r.UUID(&p.UUID)
	}
	return nil
}

func (p *LoginStart) Encode(w *encoding.Writer) error {
	_ = w.String(p.Name)
	// If server is not in online mode (probably)
	var hasUUID bool
	if _ = w.Bool(hasUUID); hasUUID {
		return w.UUID(p.UUID)
	}

	return nil
}
