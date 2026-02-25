# Logical Components - Unit 1: Database

## 컴포넌트 구성

```
backend/internal/
├── model/          # GORM 모델 정의
│   ├── store.go
│   ├── category.go
│   ├── table.go
│   ├── session.go
│   ├── menu.go
│   ├── order.go
│   └── base.go     # 공통 필드 (BaseModel)
├── repository/     # 데이터 접근 레이어
│   ├── store.go
│   ├── category.go
│   ├── table.go
│   ├── menu.go
│   └── order.go
└── database/       # DB 초기화
    ├── database.go  # 연결 + AutoMigrate
    └── seed.go      # 시드 데이터
```

## 공통 베이스 모델

```go
// BaseModel - GORM 공통 필드
type BaseModel struct {
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

## 환경변수

| 변수 | 기본값 | 설명 |
|------|--------|------|
| DB_HOST | localhost | MySQL 호스트 |
| DB_PORT | 3306 | MySQL 포트 |
| DB_USER | root | DB 사용자 |
| DB_PASSWORD | - | DB 비밀번호 |
| DB_NAME | table_order | 데이터베이스명 |
| DB_SEED | false | 시드 데이터 실행 여부 |

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
