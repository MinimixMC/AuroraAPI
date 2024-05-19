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
	mutex       sync.Mutex           // Thread safety
}

// Register an EventHandler to the event
func (e *BaseEvent) Register(handler EventHandler) int {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	id := len(e.Handlers)
	e.Handlers[id] = handler
	return id
}

// Remove an EventHandler from the event
func (e *BaseEvent) Unregister(index int) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.Handlers, index)
}

// Executes all the registered handle functions and returns if the event was cancelled
func (e *BaseEvent) Fire() bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, handle := range e.Handlers {
		handle(e)
		if e.IsCanceled() {
			return true
		}
	}
	return false
}

// Mark the event as cancelled if it is cancellable
func (e *BaseEvent) Cancel() {
	if e.cancellable {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		e.cancelled = true
	}
}

// Checks if the event has been cancelled
func (e *BaseEvent) IsCanceled() bool {
	if e.cancellable {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		return e.cancelled
	}
	return false
}
