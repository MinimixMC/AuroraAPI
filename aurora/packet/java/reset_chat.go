package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type ResetChat struct{}

func (p ResetChat) ID() int32 {
	return 0x06
}

func (p *ResetChat) Decode(r *encoding.Reader) error {
	return nil
}

func (p *ResetChat) Encode(w *encoding.Writer) error {
	return nil
}
