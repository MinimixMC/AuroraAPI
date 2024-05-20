package aurora

import (
	"sync"
	"time"
)

type Runnable struct {
	Func func()

	cancelled bool
	mu        sync.Mutex
}

func (r *Runnable) Run(delay time.Duration, period time.Duration) {
	time.Sleep(delay) // Initial delay before starting the periodic execution

	ticker := time.NewTicker(period)
	defer ticker.Stop() // Ensure the ticker is stopped to free resources

	for range ticker.C {
		r.mu.Lock()
		if r.cancelled {
			r.mu.Unlock()
			return // Stop the task if cancelled
		}
		r.mu.Unlock()
		r.Func()
	}
}

func (r *Runnable) RunAsync(delay time.Duration, period time.Duration) {
	go func() {
		time.Sleep(delay) // Initial delay before starting the periodic execution

		ticker := time.NewTicker(period)
		defer ticker.Stop() // Ensure the ticker is stopped to free resources

		for range ticker.C {
			r.mu.Lock()
			if r.cancelled {
				r.mu.Unlock()
				return // Stop the task if cancelled
			}
			r.mu.Unlock()
			r.Func()
		}
	}()
}

func (r *Runnable) Cancel() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cancelled = true
}
