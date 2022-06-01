package port

import "github.com/ricardojonathanromero/observer-pattern/internal/domain/models"

// Observer defines a standard interface to
// listen for a specific event./*
type Observer interface {
	// OnNotify allows to publish an event
	OnNotify(event models.Event)
}

// Notifier is the instance being observer./*
type Notifier interface {
	// Register itself to listen/observe events
	Register(observer Observer)
	// Unregister remove itself from the collection of observers/listeners
	Unregister(observer Observer)
	// Notify publishes new events to listeners
	Notify(event models.Event)
	// NotifyAsync publishes new events to listeners asynchronously
	NotifyAsync(event models.Event)
}
