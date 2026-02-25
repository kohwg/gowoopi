package model

import (
	"testing"

	"gorm.io/gorm"
)

func TestStore_BeforeCreate_GeneratesUUID(t *testing.T) {
	store := &Store{}
	err := store.BeforeCreate(&gorm.DB{})
	if err != nil {
		t.Fatalf("BeforeCreate() error = %v", err)
	}
	if store.ID == "" {
		t.Error("BeforeCreate() should generate UUID when ID is empty")
	}
	if len(store.ID) != 36 {
		t.Errorf("BeforeCreate() generated ID length = %d, want 36", len(store.ID))
	}
}

func TestStore_BeforeCreate_PreservesExistingID(t *testing.T) {
	existing := "existing-uuid-1234"
	store := &Store{ID: existing}
	err := store.BeforeCreate(&gorm.DB{})
	if err != nil {
		t.Fatalf("BeforeCreate() error = %v", err)
	}
	if store.ID != existing {
		t.Errorf("BeforeCreate() changed existing ID from %q to %q", existing, store.ID)
	}
}
