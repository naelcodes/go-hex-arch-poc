package pubsub

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Event struct {
	Type          string
	Data          interface{}
	CorrelationID string
}

type Subscriber struct {
	EventType    string
	Handler      func(eventData interface{}) error
	ErrorHandler func(err error, broker *Broker)
}

type Broker struct {
	subscribers map[string][]*Subscriber
	store       map[string]*Event // Event store with correlation ID as key
	mutex       sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*Subscriber),
		store:       make(map[string]*Event),
	}
}

func (b *Broker) Subscribe(eventType string, handler func(eventData interface{}) error, errorHandler func(err error, broker *Broker)) error {
	if handler == nil {
		return errors.New("handler function cannot be nil")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	subscriber := &Subscriber{
		EventType:    eventType,
		Handler:      handler,
		ErrorHandler: errorHandler,
	}

	b.subscribers[eventType] = append(b.subscribers[eventType], subscriber)

	return nil
}

func (b *Broker) UnsubscribeTopic(eventType string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	delete(b.subscribers, eventType)
}

func (b *Broker) UnsubscribeAll() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.subscribers = make(map[string][]*Subscriber)
}

func (b *Broker) Publish(eventType string, eventData interface{}) {
	correlationID := generateCorrelationID()

	event := &Event{
		Type:          eventType,
		Data:          eventData,
		CorrelationID: correlationID,
	}

	b.mutex.RLock()

	subscribers := b.subscribers[eventType]

	b.mutex.RUnlock()

	b.mutex.Lock()
	b.store[correlationID] = event // Store event with correlation ID
	b.mutex.Unlock()

	for _, subscriber := range subscribers {
		go func(subscriber *Subscriber) {
			b.notifySubscriber(subscriber, event)
		}(subscriber)
	}
}

func (b *Broker) PublishBatch(events []*Event) {
	for _, event := range events {
		if subscribers, ok := b.subscribers[event.Type]; ok {
			for _, subscriber := range subscribers {
				go func(subscriber *Subscriber) {
					b.notifySubscriber(subscriber, event)
				}(subscriber)
			}
		}
	}
}

func (b *Broker) RetrieveEvent(correlationID string) (*Event, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	event, ok := b.store[correlationID]
	return event, ok
}

func (b *Broker) ClearEvent(correlationID string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	delete(b.store, correlationID)
}

func (b *Broker) ClearAllEvents() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.store = make(map[string]*Event)
}

func (b *Broker) notifySubscriber(subscriber *Subscriber, event *Event) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && subscriber.ErrorHandler != nil {
				subscriber.ErrorHandler(err, b)
			} else {
				panic(r)
			}
		}
	}()

	if err := subscriber.Handler(event.Data); err != nil {
		panic(err)
	}
}

func generateCorrelationID() string {
	return uuid.New().String()
}
