package status

import (
	"bytes"
	"encoding/json"

	"github.com/MinimixMC/AuroraAPI/API/aurora/data/chat"
)

type StatusResponse struct {
	Version             VersionInfo  `json:"version"`
	Players             PlayersInfo  `json:"players"`
	Description         chat.Message `json:"description"`
	Favicon             string       `json:"favicon,omitempty"` // Base64 ("data:image/png;base64,<data>")
	EnforcesSecureChat  bool         `json:"enforcesSecureChat"`
	PreviewsChat        bool         `json:"previewsChat"`
	PreventsChatReports bool         `json:"preventsChatReports"` // No Chat Reports support
}

type VersionInfo struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type PlayersInfo struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []PlayerSample `json:"sample"`
}
type PlayerSample struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

func (sr *StatusResponse) NewStatus() *Status {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	_ = enc.Encode(sr)

	return &Status{
		Response: sr,
		enc:      enc,
		buf:      &buf,
	}
}

type Status struct {
	Response *StatusResponse

	enc *json.Encoder
	buf *bytes.Buffer
}

func (s *Status) JsonBytes() []byte {
	return s.buf.Bytes()
}

func (s *Status) UpdateStatus(sr *StatusResponse) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	_ = enc.Encode(sr)

	s.Response = sr
	s.buf = &buf
	s.enc = enc
}
