package aurora

func RegisterEvent(e Event, handler EventHandler) {
	e.Subscribe(handler)
}

type Event interface {
	Fire() bool                 // Fire the event
	Subscribe(EventHandler) int // Subscribe to the event
	Unsubscribe(int)            // Unsubscribe from the event
	IsCanceled() bool           // If the event has been cancelled
	Cancel()                    // Cancels the event
}

type EventHandler func(e Event) Event

// Example
type UnknownEvent struct {
	handlers    map[int]EventHandler // All event handler functions
	cancellable bool                 // If the event can be cancelled
	cancelled   bool                 // If the event is cancelled
}

func (e *UnknownEvent) Subscribe(handler EventHandler) int {
	e.handlers[len(e.handlers)] = handler
	return len(e.handlers)
}

func (e *UnknownEvent) Unsubscribe(index int) {
	delete(e.handlers, index)
}

// Fires all the registered handle functions and returns if the event was cancelled
func (e *UnknownEvent) Fire() bool {
	for _, handle := range e.handlers {
		event := handle(e)
		if event.IsCanceled() {
			return true
		}
	}
	return false
}

func (e *UnknownEvent) Cancel() {
	e.cancelled = e.cancellable
}

func (e *UnknownEvent) IsCanceled() bool {
	return e.cancelled
}
