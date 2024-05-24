package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

// Login
type LoginCookieResponse struct {
	Key     string
	Payload []byte
}

func (p LoginCookieResponse) ID() int32 {
	return 0x04
}

func (p *LoginCookieResponse) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Key)
	var hasPayload bool
	_ = r.Bool(&hasPayload)
	if hasPayload {
		return r.ByteArray(&p.Payload)
	}
	return nil
}

func (p *LoginCookieResponse) Encode(w *encoding.Writer) error {
	_ = w.String(p.Key)
	var hasPayload bool
	_ = w.Bool(hasPayload)
	if hasPayload {
		return w.ByteArray(p.Payload)
	}
	return nil
}

// Config
type ConfigCookieResponse struct {
	Key     string
	Payload []byte
}

func (p ConfigCookieResponse) ID() int32 {
	return 0x01
}

func (p *ConfigCookieResponse) Decode(r *encoding.Reader) error {
	_ = r.String(&p.Key)
	var hasPayload bool
	_ = r.Bool(&hasPayload)
	if hasPayload {
		return r.ByteArray(&p.Payload)
	}
	return nil
}

func (p *ConfigCookieResponse) Encode(w *encoding.Writer) error {
	_ = w.String(p.Key)
	var hasPayload bool
	_ = w.Bool(hasPayload)
	if hasPayload {
		return w.ByteArray(p.Payload)
	}
	return nil
}
