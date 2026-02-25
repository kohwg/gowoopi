package database

import (
	"fmt"
	"time"

	"github.com/kohwg/gowoopi/backend/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config - DB 연결 설정
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Connect - DB 연결 + Connection Pool 설정
func Connect(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// AutoMigrate - 스키마 자동 마이그레이션
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Store{},
		&model.Category{},
		&model.Table{},
		&model.TableSession{},
		&model.Menu{},
		&model.Order{},
		&model.OrderItem{},
		&model.OrderHistory{},
	)
}
