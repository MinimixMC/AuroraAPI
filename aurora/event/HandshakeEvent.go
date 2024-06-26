package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/aurora"
	"github.com/MinimixMC/AuroraAPI/aurora/data/chat"
)

var (
	// Fired on the initial connection request (handshake with next state 0x02)
	HandshakeEvent = &Handshake{
		BaseEvent:   *aurora.NewBaseEvent(true),
		KickMessage: chat.BuildMessage("Unknown reason"),
	}
)

// Fired on the initial connection request (handshake with next state 0x02)
type Handshake struct {
	Protocol   int32
	RemoteAddr net.Addr

	KickMessage chat.Message

	aurora.BaseEvent
}

// Executes all the registered handle functions and returns if the event was cancelled
func (e *Handshake) Fire() bool {
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
func (e *Handshake) Disallow(reason chat.Message) {
	e.KickMessage = reason
	e.Cancel()
}
