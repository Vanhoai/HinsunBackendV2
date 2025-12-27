package events

import (
	"context"
	"sync"
)

type AsyncEventBus struct {
	handlers []EventHandler
	mutex    sync.RWMutex
	wg       sync.WaitGroup
}

// NewAsyncEventBus creates a new asynchronous event bus
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{handlers: make([]EventHandler, 0)}
}

// Publish publishes an event to all interested handlers asynchronously
func (b *AsyncEventBus) Publish(ctx context.Context, event Event) error {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	for _, handler := range b.handlers {
		if handler.InterestedIn(event.EventName()) {
			// Increase the wait group counter
			b.wg.Add(1)

			// Execute handler asynchronously
			go func(h EventHandler, e Event) {
				defer b.wg.Done()
				// Handle errors in background, you might want to log them
				_ = h.HandleEvent(ctx, e)
			}(handler, event)
		}
	}

	return nil
}

// Subscribe registers an event handler
func (b *AsyncEventBus) Subscribe(handler EventHandler) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.handlers = append(b.handlers, handler)
}

// Unsubscribe removes an event handler
func (b *AsyncEventBus) Unsubscribe(handler EventHandler) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for i, h := range b.handlers {
		if h == handler {
			b.handlers = append(b.handlers[:i], b.handlers[i+1:]...)
			break
		}
	}
}

// Wait waits for all async handlers to complete (useful for testing or graceful shutdown)
func (b *AsyncEventBus) Wait() {
	b.wg.Wait()
}

// PublishAndWait publishes an event and waits for all handlers to complete
func (b *AsyncEventBus) PublishAndWait(ctx context.Context, event Event) error {
	if err := b.Publish(ctx, event); err != nil {
		return err
	}
	b.Wait()
	return nil
}

// AsyncEventBusWithErrorHandling includes error logging
type AsyncEventBusWithErrorHandling struct {
	*AsyncEventBus
	errorHandler func(event Event, err error)
}

// NewAsyncEventBusWithErrorHandling creates a new async event bus with error handling
func NewAsyncEventBusWithErrorHandling(errorHandler func(Event, error)) *AsyncEventBusWithErrorHandling {
	return &AsyncEventBusWithErrorHandling{
		AsyncEventBus: NewAsyncEventBus(),
		errorHandler:  errorHandler,
	}
}
