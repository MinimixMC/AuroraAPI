package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/status"
)

var (
	PingEvent = &Ping{
		BaseEvent: *aurora.NewBaseEvent(true),
	}
)

type Ping struct {
	Status     *status.StatusResponse // The current status response
	RemoteAddr net.Addr

	aurora.BaseEvent
}
