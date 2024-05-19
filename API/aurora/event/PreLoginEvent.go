package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/chat"
)

var (
	// Fired when login starts after the initial handshake
	PreLoginEvent = &Handshake{
		BaseEvent:   *aurora.NewBaseEvent(true),
		kickmessage: chat.BuildMessage("Unknown reason"),
	}
)

// Fired when login starts after the initial handshake
type PreLogin struct {
	Name       string
	UUID       [16]byte
	Protocol   int32
	RemoteAddr net.Addr

	kickmessage chat.Message

	aurora.BaseEvent
}

// Executes all the registered handle functions and returns if the event was cancelled
func (e *PreLogin) Fire() bool {
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
func (e *PreLogin) Disallow(reason chat.Message) {
	e.kickmessage = reason
	e.Cancel()
}
