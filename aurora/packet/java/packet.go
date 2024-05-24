package packet

import (
	"errors"

	"github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"
)

var ErrNotImplemented = errors.New("not implemented")

type Packet interface {
	ID() int32
	Decode(r *encoding.Reader) error
	Encode(w *encoding.Writer) error
}

// Unknown packet
type Unknown struct {
	Id      int32
	Payload []byte
}

func (u Unknown) ID() int32 {
	return u.Id
}

func (u Unknown) Decode(r *encoding.Reader) error {
	return nil
}

func (u Unknown) Encode(w *encoding.Writer) error {
	return w.FixedByteArray(u.Payload)
}
