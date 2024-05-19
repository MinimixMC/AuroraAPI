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

// Executes all the registered handle functions and returns if the event was cancelled
func (e *Login) Fire() bool {
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
