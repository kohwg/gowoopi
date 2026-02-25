package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/middleware"
	"github.com/kohwg/gowoopi/backend/internal/service"
)

type SSEHandler struct {
	sse service.SSEManager
}

func NewSSEHandler(sse service.SSEManager) *SSEHandler {
	return &SSEHandler{sse: sse}
}

func (h *SSEHandler) StreamOrders(c *gin.Context) {
	claims := middleware.GetClaims(c)
	ch := h.sse.Subscribe(claims.StoreID)
	defer h.sse.Unsubscribe(claims.StoreID, ch)

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	c.Stream(func(w io.Writer) bool {
		select {
		case event, ok := <-ch:
			if !ok {
				return false
			}
			data, _ := json.Marshal(event)
			_, _ = fmt.Fprintf(w, "data: %s\n\n", data)
			return true
		case <-ticker.C:
			_, _ = fmt.Fprintf(w, ": heartbeat\n\n")
			return true
		case <-c.Request.Context().Done():
			return false
		}
	})
}
