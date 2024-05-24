package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type LoginAcknowledged struct {
}

func (p LoginAcknowledged) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledged) Decode(r *encoding.Reader) error {
	return nil
}

func (p *LoginAcknowledged) Encode(r *encoding.Writer) error {
	return nil
}
