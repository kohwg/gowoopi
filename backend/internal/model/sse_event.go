package model

type SSEEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

const (
	SSEOrderCreated       = "ORDER_CREATED"
	SSEOrderStatusChanged = "ORDER_STATUS_CHANGED"
	SSEOrderDeleted       = "ORDER_DELETED"
	SSETableReset         = "TABLE_RESET"
)
