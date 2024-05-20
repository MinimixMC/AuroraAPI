package aurora

func RegisterEvent(e Event, handler EventHandler) {
	e.Register(handler)
}

func Fire(e Event) {
	e.Fire()
}

type Event interface {
	Fire() bool                // Fire the event
	Register(EventHandler) int // Subscribe to the event
	Unregister(int)            // Unsubscribe from the event
	IsCanceled() bool          // If the event has been cancelled
	Cancel()                   // Cancels the event
}

type EventHandler func(e Event) Event
