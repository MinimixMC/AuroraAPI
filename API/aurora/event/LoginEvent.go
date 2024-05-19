package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/chat"
)

var (
	// Fired before the sending the login success packet
	LoginEvent = &Login{
		BaseEvent:   *aurora.NewBaseEvent(true),
		KickMessage: chat.BuildMessage("Unknown reason"),
	}
)

// Fired before the sending the login success packet
type Login struct {
	Name       string
	UUID       [16]byte
	Protocol   int32
	RemoteAddr net.Addr

	KickMessage chat.Message

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

// Whether to allow the user to connect or not.
func (e *Login) Disallow(reason chat.Message) {
	e.KickMessage = reason
	e.Cancel()
}
