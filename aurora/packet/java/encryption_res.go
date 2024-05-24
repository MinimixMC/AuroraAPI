package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (p EncryptionResponse) ID() int32 {
	return 0x01
}

func (p *EncryptionResponse) Decode(r *encoding.Reader) error {
	_ = r.ByteArray(&p.SharedSecret)
	return r.ByteArray(&p.VerifyToken)
}

func (p *EncryptionResponse) Encode(w *encoding.Writer) error {
	_ = w.ByteArray(p.SharedSecret)
	return w.ByteArray(p.VerifyToken)
}
