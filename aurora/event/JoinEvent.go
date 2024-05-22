package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/aurora"
	"github.com/MinimixMC/AuroraAPI/aurora/data/chat"
)

var (
	JoinEvent = &Join{
		BaseEvent: *aurora.NewBaseEvent(true),
	}
)

type Join struct {
	Protocol   int32
	RemoteAddr net.Addr

	message *chat.Message

	aurora.BaseEvent
}

// Executes all the registered handle functions and returns if the event was cancelled
func (e *Join) Fire() bool {
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

// Get's the current join message
func (e *Join) JoinMessage() *chat.Message {
	return e.message
}

// Set's the join message
// if nil, no join message will be sent
func (e *Join) SetJoinMessage(message *chat.Message) {
	e.message = message
}
