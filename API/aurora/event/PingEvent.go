package event

import "github.com/MinimixMC/AuroraAPI/API/aurora"

var (
	PingEvent = &Ping{
		Handlers:    map[int]aurora.EventHandler{},
		cancellable: true,
		cancelled:   false,
	}
)

type Ping struct {
	Handlers    map[int]aurora.EventHandler // All event handler functions
	cancellable bool                        // If the event can be cancelled
	cancelled   bool                        // If the event is cancelled
}

func (e *Ping) Subscribe(handler aurora.EventHandler) int {
	e.Handlers[len(e.Handlers)] = handler
	return len(e.Handlers)
}

func (e *Ping) Unsubscribe(index int) {
	delete(e.Handlers, index)
}

// Fires all the registered handle functions and returns if the event was cancelled
func (e *Ping) Fire() bool {
	for _, handle := range e.Handlers {
		event := handle(e)
		if event.IsCanceled() {
			return true
		}
	}
	return false
}

func (e *Ping) Cancel() {
	e.cancelled = e.cancellable
}

func (e *Ping) IsCanceled() bool {
	return e.cancelled
}
