package event

import (
	"net"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
)

var (
	LoginEvent = &Login{
		handlers:    map[int]aurora.EventHandler{},
		cancellable: true,
		cancelled:   false,
	}
)

type Login struct {
	Protocol   int32
	RemoteAddr net.Addr

	handlers    map[int]aurora.EventHandler // All event handler functions
	cancellable bool                        // If the event can be cancelled
	cancelled   bool                        // If the event is cancelled
}

func (e *Login) Subscribe(handler aurora.EventHandler) int {
	e.handlers[len(e.handlers)] = handler
	return len(e.handlers)
}

func (e *Login) Unsubscribe(index int) {
	delete(e.handlers, index)
}

// Fires all the registered handle functions and returns if the event was cancelled
func (e *Login) Fire() bool {
	for _, handle := range e.handlers {
		event := handle(e)
		if event.IsCanceled() {
			return true
		}
	}
	return false
}

// Cancel the event
func (e *Login) Cancel() {
	e.cancelled = e.cancellable
}

// Returns if the event is cancelled
func (e *Login) IsCanceled() bool {
	return e.cancelled
}
