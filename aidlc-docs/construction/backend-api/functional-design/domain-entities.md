# Domain Entities - Unit 2: Backend API

## 신규 모델: Admin (매장당 다수 관리자)

### model/admin.go
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 관리자 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| username | VARCHAR(50) | NOT NULL | 관리자 아이디 |
| password_hash | VARCHAR(255) | NOT NULL | bcrypt 해시 |
| name | VARCHAR(50) | NOT NULL | 관리자 이름 |
| BaseModel | | | Soft Delete 포함 |

**UNIQUE**: (store_id, username, deleted_at)

**참고**: Store 테이블의 admin_username/admin_password_hash 컬럼은 제거

---

## Request/Response DTOs

### Auth
```go
// CustomerLoginRequest
type CustomerLoginRequest struct {
    StoreID     string `json:"store_id" binding:"required"`
    TableNumber int    `json:"table_number" binding:"required"`
    Password    string `json:"password" binding:"required"`
}

// AdminLoginRequest
type AdminLoginRequest struct {
    StoreID  string `json:"store_id" binding:"required"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// TokenResponse (Cookie로 전달, body는 사용자 정보만)
type AuthResponse struct {
    Role    string `json:"role"`    // "customer" | "admin"
    StoreID string `json:"store_id"`
}
```

### Menu
```go
type MenuCreateRequest struct {
    CategoryID  uint   `json:"category_id" binding:"required"`
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
    Price       uint   `json:"price" binding:"required"`
    ImageURL    string `json:"image_url"`
    IsAvailable *bool  `json:"is_available"`
}

type MenuUpdateRequest struct {
    CategoryID  *uint   `json:"category_id"`
    Name        *string `json:"name"`
    Description *string `json:"description"`
    Price       *uint   `json:"price"`
    ImageURL    *string `json:"image_url"`
    IsAvailable *bool   `json:"is_available"`
}

type MenuOrderRequest struct {
    ID        uint `json:"id" binding:"required"`
    SortOrder int  `json:"sort_order" binding:"required"`
}
```

### Order
```go
type OrderCreateRequest struct {
    Items []OrderItemRequest `json:"items" binding:"required,min=1"`
}

type OrderItemRequest struct {
    MenuID   uint `json:"menu_id" binding:"required"`
    Quantity uint `json:"quantity" binding:"required,min=1"`
}

type StatusUpdateRequest struct {
    Status string `json:"status" binding:"required"`
}
```

### Table
```go
type TableSetupRequest struct {
    TableNumber int    `json:"table_number" binding:"required"`
    Password    string `json:"password" binding:"required"`
}
```

### Error Response
```go
type ErrorResponse struct {
    Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
```

### JWT Claims
```go
type Claims struct {
    StoreID   string `json:"store_id"`
    Role      string `json:"role"`       // "customer" | "admin"
    TableID   uint   `json:"table_id"`   // customer only
    SessionID string `json:"session_id"` // customer only
    AdminID   uint   `json:"admin_id"`   // admin only
    jwt.RegisteredClaims
}
```

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
