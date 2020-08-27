package rabbit

import (
	"fmt"
	"sync"
)

var (
	ErrQueueNotFound = fmt.Errorf("Queue not found in map")
)

// QueueMap protected of concurrencies
// https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c use map instead of built in sync.map
type QueueMap struct {
	sync.RWMutex
	internal map[string]Queue
}

// NewQueueMap return a new QueueMap
func NewQueueMap() *QueueMap {
	return &QueueMap{
		internal: make(map[string]Queue),
	}
}

// Load retrieve a value from the Queue map
func (queueMap *QueueMap) Load(key string) (value Queue, ok bool) {
	queueMap.RLock()
	result, ok := queueMap.internal[key]
	queueMap.RUnlock()
	return result, ok
}

// Delete value from QueueMap
func (queueMap *QueueMap) Delete(key string) {
	queueMap.Lock()
	delete(queueMap.internal, key)
	queueMap.Unlock()
}

// Add quantity to Queue
func (queueMap *QueueMap) Add(key string) {
	queueMap.Lock()
	result, ok := queueMap.internal[key]
	if ok {
		result.Add()
	}
	queueMap.Unlock()
}

// Store update or insert value in QueueMap
func (queueMap *QueueMap) Store(value Queue) {
	queueMap.Lock()
	queueMap.internal[value.Name] = value
	queueMap.Unlock()
}

// NewInternal clears the internal map
func (queueMap *QueueMap) NewInternal() {
	queueMap.internal = make(map[string]Queue)
}

// Internal returns the map internal
func (queueMap *QueueMap) Internal() map[string]Queue {
	queueMap.Lock()
	internal := queueMap.internal
	queueMap.Unlock()

	return internal
}

// Keys returns the map internal keys
func (queueMap *QueueMap) Keys() []string {
	var keys []string
	queueMap.Lock()
	for key := range queueMap.internal {
		keys = append(keys, key)
	}
	queueMap.Unlock()

	return keys
}

// StoreInteral stores the entire map inside internal
func (queueMap *QueueMap) StoreInteral(QueueMap map[string]Queue) {
	queueMap.Lock()
	queueMap.internal = QueueMap
	queueMap.Unlock()
}
