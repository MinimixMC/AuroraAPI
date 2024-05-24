package packet

import "github.com/MinimixMC/AuroraAPI/aurora/protocol/encoding"

// Login
type LoginCookieRequest struct {
	Key string
}

func (p LoginCookieRequest) ID() int32 {
	return 0x05
}

func (p *LoginCookieRequest) Decode(r *encoding.Reader) error {
	return r.String(&p.Key)
}

func (p *LoginCookieRequest) Encode(w *encoding.Writer) error {
	return w.String(p.Key)
}

// Config
type ConfigCookieRequest struct {
	Key string
}

func (p ConfigCookieRequest) ID() int32 {
	return 0x00
}

func (p *ConfigCookieRequest) Decode(r *encoding.Reader) error {
	return r.String(&p.Key)
}

func (p *ConfigCookieRequest) Encode(w *encoding.Writer) error {
	return w.String(p.Key)
}
