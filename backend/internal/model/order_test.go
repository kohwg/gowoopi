package model

import (
	"testing"
)

func TestOrderStatus_IsValid(t *testing.T) {
	tests := []struct {
		status OrderStatus
		want   bool
	}{
		{OrderStatusPending, true},
		{OrderStatusConfirmed, true},
		{OrderStatusPreparing, true},
		{OrderStatusCompleted, true},
		{OrderStatus("INVALID"), false},
		{OrderStatus(""), false},
	}
	for _, tt := range tests {
		t.Run(string(tt.status), func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("OrderStatus(%q).IsValid() = %v, want %v", tt.status, got, tt.want)
			}
		})
	}
}

func TestOrderStatus_CanTransitionTo(t *testing.T) {
	tests := []struct {
		from OrderStatus
		to   OrderStatus
		want bool
	}{
		{OrderStatusPending, OrderStatusConfirmed, true},
		{OrderStatusConfirmed, OrderStatusPreparing, true},
		{OrderStatusPreparing, OrderStatusCompleted, true},
		// 역방향 불가
		{OrderStatusConfirmed, OrderStatusPending, false},
		{OrderStatusPreparing, OrderStatusConfirmed, false},
		{OrderStatusCompleted, OrderStatusPreparing, false},
		// 건너뛰기 불가
		{OrderStatusPending, OrderStatusPreparing, false},
		{OrderStatusPending, OrderStatusCompleted, false},
		{OrderStatusConfirmed, OrderStatusCompleted, false},
		// 최종 상태에서 전이 불가
		{OrderStatusCompleted, OrderStatusPending, false},
		// 동일 상태 전이 불가
		{OrderStatusPending, OrderStatusPending, false},
	}
	for _, tt := range tests {
		name := string(tt.from) + "->" + string(tt.to)
		t.Run(name, func(t *testing.T) {
			if got := tt.from.CanTransitionTo(tt.to); got != tt.want {
				t.Errorf("%s.CanTransitionTo(%s) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}
