package event

import (
	"net"
)

var (
	LoginEvent = &Login{
		BaseEvent: *NewBaseEvent(true),
	}
)

type Login struct {
	Protocol   int32
	RemoteAddr net.Addr

	BaseEvent
}
