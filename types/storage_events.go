package types

// ----------------------------------------------------------------------------
// Event Manager
// ----------------------------------------------------------------------------

// StorageEventManager implements a simple wrapper around a slice of Event objects that
// can be emitted from.
type StorageEventManager struct {
	events StorageEvents
}

func NewStorageEventManager() *StorageEventManager {
	return &StorageEventManager{EmptyStorageEvents()}
}

func (em *StorageEventManager) StorageEvents() StorageEvents { return em.events }

// EmitStorageEvent stores a single StorageEvent object.
func (em *StorageEventManager) EmitStorageEvent(event StorageEvent) {
	em.events = em.events.AppendEvent(event)
}

// EmitStorageEvents stores a series of StorageEvent objects.
func (em *StorageEventManager) EmitStorageEvents(events StorageEvents) {
	em.events = em.events.AppendEvents(events)
}

type (
	// StorageEvents defines a slice of StorageEvent objects
	StorageEvents []StorageEvent

	// StorageEvents defines a slice of StorageEvent objects
	Attributes []Attribute
)

// EmptyStorageEvents returns an empty slice of events.
func EmptyStorageEvents() StorageEvents {
	return make(StorageEvents, 0)
}

// AppendEvent adds an Event to a slice of events.
func (e StorageEvents) AppendEvent(event StorageEvent) StorageEvents {
	return append(e, event)
}

// AppendEvents adds a slice of Event objects to an exist slice of Event objects.
func (e StorageEvents) AppendEvents(events StorageEvents) StorageEvents {
	return append(e, events...)
}

// StorageEvent for storage update event tracking
type StorageEvent struct {
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

// NewStorageEvent creates a new StorageEvent object with a given type and slice of one or more
// attributes.
func NewStorageEvent(ty string, attrs ...Attribute) StorageEvent {
	e := StorageEvent{Type: ty}

	for _, attr := range attrs {
		e.Attributes = append(e.Attributes, NewAttribute(attr.Key, attr.Value))
	}

	return e
}

// AppendAttributes adds one or more attributes to an Event.
func (e StorageEvent) AppendAttributes(attrs ...Attribute) StorageEvent {
	for _, attr := range attrs {
		e.Attributes = append(e.Attributes, attr)
	}
	return e
}

// AppendStorageEvent adds an StorageEvent to a slice of StorageEvents.
func (e StorageEvents) AppendStorageEvent(storageEvent StorageEvent) StorageEvents {
	return append(e, storageEvent)
}

// AppendStorageEvents adds a slice of StorageEvent objects to an exist slice of StorageEvent objects.
func (e StorageEvents) AppendStorageEvents(storageEvents StorageEvents) StorageEvents {
	return append(e, storageEvents...)
}
