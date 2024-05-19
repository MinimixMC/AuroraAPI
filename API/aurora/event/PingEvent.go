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

// Executes all the registered handle functions and returns if the event was cancelled
func (e *Ping) Fire() bool {
	e.BaseEvent.Mutex.Lock()
	defer e.BaseEvent.Mutex.Unlock()
	for _, handle := range e.Handlers {
		event := handle(e)
		if event.IsCanceled() {
			return true
		}
	}
	return false
}
