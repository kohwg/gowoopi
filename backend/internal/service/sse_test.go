package service

import (
	"testing"
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
)

func TestSSEManager_SubscribeAndBroadcast(t *testing.T) {
	mgr := NewSSEManager()
	ch := mgr.Subscribe("store1")

	event := model.SSEEvent{Type: model.SSEOrderCreated, Data: "test"}
	mgr.Broadcast("store1", event)

	select {
	case received := <-ch:
		if received.Type != model.SSEOrderCreated {
			t.Errorf("type = %q, want ORDER_CREATED", received.Type)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for event")
	}
}

func TestSSEManager_Unsubscribe(t *testing.T) {
	mgr := NewSSEManager()
	ch := mgr.Subscribe("store1")
	mgr.Unsubscribe("store1", ch)

	// channel should be closed
	_, ok := <-ch
	if ok {
		t.Error("channel should be closed after unsubscribe")
	}
}

func TestSSEManager_BroadcastToMultipleClients(t *testing.T) {
	mgr := NewSSEManager()
	ch1 := mgr.Subscribe("store1")
	ch2 := mgr.Subscribe("store1")

	event := model.SSEEvent{Type: model.SSEOrderCreated, Data: "test"}
	mgr.Broadcast("store1", event)

	for _, ch := range []chan model.SSEEvent{ch1, ch2} {
		select {
		case received := <-ch:
			if received.Type != model.SSEOrderCreated {
				t.Errorf("type = %q, want ORDER_CREATED", received.Type)
			}
		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestSSEManager_BroadcastIsolation(t *testing.T) {
	mgr := NewSSEManager()
	ch1 := mgr.Subscribe("store1")
	ch2 := mgr.Subscribe("store2")

	mgr.Broadcast("store1", model.SSEEvent{Type: "test", Data: nil})

	select {
	case <-ch1:
		// expected
	case <-time.After(time.Second):
		t.Fatal("store1 should receive event")
	}

	select {
	case <-ch2:
		t.Fatal("store2 should not receive store1 event")
	case <-time.After(100 * time.Millisecond):
		// expected
	}
}
