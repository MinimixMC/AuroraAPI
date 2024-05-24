package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

type EncryptionRequest struct {
	ServerID           string
	PublicKey          []byte
	VerifyToken        []byte
	ShouldAuthenticate bool
}

func (p EncryptionRequest) ID() int32 {
	return 0x01
}

func (p *EncryptionRequest) Decode(r *encoding.Reader) error {
	_ = r.String(&p.ServerID)
	_ = r.ByteArray(&p.PublicKey)
	_ = r.ByteArray(&p.VerifyToken)
	return r.Bool(&p.ShouldAuthenticate)
}

func (p *EncryptionRequest) Encode(w *encoding.Writer) error {
	_ = w.String(p.ServerID)
	_ = w.ByteArray(p.PublicKey)
	_ = w.ByteArray(p.VerifyToken)
	return w.Bool(p.ShouldAuthenticate)
}
