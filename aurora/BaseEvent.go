package aurora

import (
	"sync"
)

// Initializes a new BaseEvent with default values
func NewBaseEvent(cancellable bool) *BaseEvent {
	return &BaseEvent{
		Handlers: map[int]EventHandler{},
	}
}

// Base structure for all events
type BaseEvent struct {
	Handlers    map[int]EventHandler // All event handler functions
	cancellable bool                 // If the event can be cancelled
	cancelled   bool                 // If the event is cancelled
	Mutex       sync.Mutex           // Thread safety
}

// Register an EventHandler to the event
func (e *BaseEvent) Register(handler EventHandler) int {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()
	id := len(e.Handlers)
	e.Handlers[id] = handler
	return id
}

// Remove an EventHandler from the event
func (e *BaseEvent) Unregister(index int) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()
	delete(e.Handlers, index)
}

// Mark the event as cancelled if it is cancellable
func (e *BaseEvent) Cancel() {
	if e.cancellable {
		e.Mutex.Lock()
		defer e.Mutex.Unlock()
		e.cancelled = true
	}
}

// Checks if the event has been cancelled
func (e *BaseEvent) IsCanceled() bool {
	if e.cancellable {
		e.Mutex.Lock()
		defer e.Mutex.Unlock()
		return e.cancelled
	}
	return false
}

func (e *BaseEvent) GetBaseEvent() *BaseEvent {
	return e
}
