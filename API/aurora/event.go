package aurora

type Event interface {
	Fire() bool
	Subscribe(EventHandler) int
	Unsubscribe(int)
	IsCanceled() bool
	Cancel()
}

type EventHandler func(e Event) Event

// Example
type UnknownEvent struct {
	Handlers    map[int]EventHandler // All event handler functions
	cancellable bool                 // If the event can be cancelled
	cancelled   bool                 // If the event is cancelled
}

func (e *UnknownEvent) Subscribe(handler EventHandler) int {
	e.Handlers[len(e.Handlers)] = handler
	return len(e.Handlers)
}

func (e *UnknownEvent) Unsubscribe(index int) {
	delete(e.Handlers, index)
}

// Fires all the registered handle functions and returns if the event was cancelled
func (e *UnknownEvent) Fire() bool {
	for _, handle := range e.Handlers {
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
