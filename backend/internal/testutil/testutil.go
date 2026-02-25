package testutil

import (
	"fmt"
	"os"
	"testing"

	"github.com/gowoopi/backend/internal/database"
	"gorm.io/gorm"
)

// SetupTestDB - 테스트용 DB 연결 (환경변수 기반)
func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	cfg := &database.Config{
		Host:     getEnv("TEST_DB_HOST", "localhost"),
		Port:     getEnv("TEST_DB_PORT", "3306"),
		User:     getEnv("TEST_DB_USER", "root"),
		Password: getEnv("TEST_DB_PASSWORD", ""),
		DBName:   getEnv("TEST_DB_NAME", "table_order_test"),
	}

	db, err := database.Connect(cfg)
	if err != nil {
		t.Skipf("Skipping integration test: DB not available: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		t.Fatalf("AutoMigrate failed: %v", err)
	}

	return db
}

// CleanupTables - 테스트 후 테이블 데이터 정리
func CleanupTables(t *testing.T, db *gorm.DB, tables ...string) {
	t.Helper()
	for _, table := range tables {
		db.Exec(fmt.Sprintf("DELETE FROM %s", table))
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
