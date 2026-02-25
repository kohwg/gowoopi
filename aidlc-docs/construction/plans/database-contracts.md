# Contract/Interface Definition - Unit 1: Database

## Unit Context
- **Stories**: 전체 (모든 데이터 관련 스토리의 기반)
- **Dependencies**: 없음 (다른 Unit의 기반)
- **Database Entities**: Store, Category, Table, TableSession, Menu, Order, OrderItem, OrderHistory

---

## Model Layer

### model/base.go
- `BaseModel` struct: CreatedAt, UpdatedAt, DeletedAt 공통 필드

### model/store.go
- `Store` struct: ID(UUID), Name, AdminUsername, AdminPasswordHash, DefaultLanguage, BaseModel

### model/category.go
- `Category` struct: ID(uint), StoreID, Name, SortOrder, BaseModel

### model/table.go
- `Table` struct: ID(uint), StoreID, TableNumber, PasswordHash, IsActive, BaseModel

### model/session.go
- `TableSession` struct: ID(UUID), TableID, StoreID, StartedAt, EndedAt, IsActive, CreatedAt, UpdatedAt

### model/menu.go
- `Menu` struct: ID(uint), StoreID, CategoryID, Name, Description, Price, ImageURL, IsAvailable, SortOrder, BaseModel

### model/order.go
- `Order` struct: ID(UUID), SessionID, StoreID, TableID, Status, TotalAmount, BaseModel
- `OrderItem` struct: ID(uint), OrderID, MenuID, MenuName, Price, Quantity, Subtotal, CreatedAt
- `OrderStatus` type: PENDING, CONFIRMED, PREPARING, COMPLETED

### model/order_history.go
- `OrderHistory` struct: ID(uint), OriginalOrderID, StoreID, TableID, TableNumber, SessionID, Status, TotalAmount, ItemsJSON, OrderedAt, CompletedAt, CreatedAt

---

## Repository Layer

### repository/store.go
- `StoreRepository` interface:
  - `FindByID(id string) (*model.Store, error)`: 매장 조회
  - `FindByIDAndUsername(id, username string) (*model.Store, error)`: 관리자 인증용

### repository/category.go
- `CategoryRepository` interface:
  - `FindByStore(storeID string) ([]model.Category, error)`: 매장별 카테고리 조회
  - `FindByID(id uint) (*model.Category, error)`: 단건 조회
  - `Create(category *model.Category) error`: 생성
  - `Update(category *model.Category) error`: 수정
  - `Delete(id uint) error`: Soft Delete

### repository/table.go
- `TableRepository` interface:
  - `FindByStoreAndNumber(storeID string, number int) (*model.Table, error)`: 테이블 조회
  - `Create(table *model.Table) error`: 생성
  - `Update(table *model.Table) error`: 수정

### repository/session.go
- `SessionRepository` interface:
  - `Create(session *model.TableSession) error`: 세션 생성
  - `FindActiveByTable(tableID uint) (*model.TableSession, error)`: 활성 세션 조회
  - `End(sessionID string) error`: 세션 종료

### repository/menu.go
- `MenuRepository` interface:
  - `FindByStore(storeID string) ([]model.Menu, error)`: 매장별 메뉴 조회
  - `FindByID(id uint) (*model.Menu, error)`: 단건 조회
  - `Create(menu *model.Menu) error`: 생성
  - `Update(menu *model.Menu) error`: 수정
  - `Delete(id uint) error`: Soft Delete
  - `UpdateOrder(items []model.MenuOrderInput) error`: 순서 변경

### repository/order.go
- `OrderRepository` interface:
  - `Create(order *model.Order, items []model.OrderItem) error`: 주문+항목 생성 (트랜잭션)
  - `FindBySession(sessionID string) ([]model.Order, error)`: 세션별 조회
  - `FindByStore(storeID string) ([]model.Order, error)`: 매장별 조회
  - `FindByID(id string) (*model.Order, error)`: 단건 조회
  - `UpdateStatus(id string, status model.OrderStatus) error`: 상태 변경
  - `Delete(id string) error`: Soft Delete
  - `MoveToHistory(sessionID string) error`: 이력 이동 (트랜잭션)
  - `FindHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error)`: 이력 조회

---

## Database Layer

### database/database.go
- `Connect(cfg *Config) (*gorm.DB, error)`: DB 연결 + Connection Pool 설정
- `AutoMigrate(db *gorm.DB) error`: 스키마 자동 마이그레이션

### database/seed.go
- `Seed(db *gorm.DB) error`: 개발용 시드 데이터 생성

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
