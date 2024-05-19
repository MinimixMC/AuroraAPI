package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
)

var (
	LoginEvent = &Login{
		BaseEvent: *aurora.NewBaseEvent(true),
	}
)

type Login struct {
	Protocol   int32
	RemoteAddr net.Addr

	aurora.BaseEvent
}
