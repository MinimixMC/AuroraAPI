package event

import (
	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/status"
)

var (
	PingEvent = &Ping{
		handlers:    map[int]aurora.EventHandler{},
		cancellable: true,
		cancelled:   false,
	}
)

type Ping struct {
	Status *status.StatusResponse // The current status response

	handlers    map[int]aurora.EventHandler // All event handler functions
	cancellable bool                        // If the event can be cancelled
	cancelled   bool                        // If the event is cancelled
}

func (e *Ping) Subscribe(handler aurora.EventHandler) int {
	e.handlers[len(e.handlers)] = handler
	return len(e.handlers)
}

func (e *Ping) Unsubscribe(index int) {
	delete(e.handlers, index)
}

// Fires all the registered handle functions and returns if the event was cancelled
func (e *Ping) Fire() bool {
	for _, handle := range e.handlers {
		event := handle(e)
		if event.IsCanceled() {
			return true
		}
	}
	return false
}

// Cancel the event
func (e *Ping) Cancel() {
	e.cancelled = e.cancellable
}

// Returns if the event is cancelled
func (e *Ping) IsCanceled() bool {
	return e.cancelled
}
