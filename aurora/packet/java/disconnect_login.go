package packet

import (
	"github.com/MinimixMC/AuroraAPI/aurora/data/chat"
	"github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"
)

type DisconnectLogin struct {
	Reason chat.Message
}

func (l DisconnectLogin) ID() int32 {
	return 0x00
}

func (l *DisconnectLogin) Decode(r *encoding.Reader) error {
	return ErrNotImplemented
}

func (l DisconnectLogin) Encode(w *encoding.Writer) error {
	return w.String(l.Reason.JsonString())
}
