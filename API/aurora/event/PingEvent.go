package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora/data/status"
)

var (
	PingEvent = &Ping{
		BaseEvent: *NewBaseEvent(true),
	}
)

type Ping struct {
	Status     *status.StatusResponse // The current status response
	RemoteAddr net.Addr

	BaseEvent
}
