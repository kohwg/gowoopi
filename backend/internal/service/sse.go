package service

import (
	"sync"

	"github.com/gowoopi/backend/internal/model"
)

type sseManager struct {
	mu      sync.RWMutex
	clients map[string][]chan model.SSEEvent
}

func NewSSEManager() SSEManager {
	return &sseManager{clients: make(map[string][]chan model.SSEEvent)}
}

func (m *sseManager) Subscribe(storeID string) chan model.SSEEvent {
	m.mu.Lock()
	defer m.mu.Unlock()
	ch := make(chan model.SSEEvent, 10)
	m.clients[storeID] = append(m.clients[storeID], ch)
	return ch
}

func (m *sseManager) Unsubscribe(storeID string, ch chan model.SSEEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	clients := m.clients[storeID]
	for i, c := range clients {
		if c == ch {
			m.clients[storeID] = append(clients[:i], clients[i+1:]...)
			close(ch)
			return
		}
	}
}

func (m *sseManager) Broadcast(storeID string, event model.SSEEvent) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, ch := range m.clients[storeID] {
		select {
		case ch <- event:
		default:
		}
	}
}
