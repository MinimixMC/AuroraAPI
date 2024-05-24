package packet

import (
	"github.com/MinimixMC/AuroraAPI/aurora/data/types"
	"github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"
)

type LoginSuccess struct {
	UUID [16]byte
	Name string

	Properties          []types.Property
	StrictErrorHandling bool
}

func (p LoginSuccess) ID() int32 {
	return 0x02
}

func (p *LoginSuccess) Decode(r *encoding.Reader) error {
	_ = r.UUID(&p.UUID)
	_ = r.String(&p.Name)

	var length int32
	_ = r.VarInt(&length)
	prpty := make([]types.Property, length)

	for i := int32(0); i < length; i++ {
		a := prpty[i]
		_ = r.String(&a.Name)
		_ = r.String(&a.Value)

		var signed bool
		_ = r.Bool(&signed)
		if signed {
			_ = r.String(&a.Signature)
		}
	}

	p.Properties = prpty

	return r.Bool(&p.StrictErrorHandling)
}

func (p *LoginSuccess) Encode(w *encoding.Writer) error {
	_ = w.UUID(p.UUID)
	_ = w.String(p.Name)

	_ = w.VarInt(int32(len(p.Properties)))

	for i := 0; i < len(p.Properties); i++ {
		a := p.Properties[i]

		_ = w.String(a.Name)
		_ = w.String(a.Value)

		_ = w.Bool(a.Signature != "")

		if a.Signature != "" {
			_ = w.String(a.Signature)
		}
	}

	return w.Bool(p.StrictErrorHandling)
}
