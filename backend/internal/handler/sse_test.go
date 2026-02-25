package handler

import (
	"testing"

	"github.com/gowoopi/backend/internal/service"
)

func TestSSEHandler_NewSSEHandler(t *testing.T) {
	mgr := service.NewSSEManager()
	h := NewSSEHandler(mgr)
	if h == nil {
		t.Fatal("handler should not be nil")
	}
	if h.sse != mgr {
		t.Error("sse manager mismatch")
	}
}
