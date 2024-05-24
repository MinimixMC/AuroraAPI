package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type FinishConfiguration struct{}

func (p FinishConfiguration) ID() int32 {
	return 0x03
}

func (p *FinishConfiguration) Decode(r *encoding.Reader) error {
	return nil
}

func (p *FinishConfiguration) Encode(w *encoding.Writer) error {
	return nil
}
