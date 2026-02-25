# Contract/Interface Definition - Unit 2: Backend API

## Unit Context
- **Stories**: US-001~010 (전체 Feature Stories)
- **Dependencies**: Unit 1 (Database - model, repository)
- **신규 모델**: Admin, DTO types, SSEEvent, Claims, ErrorResponse

---

## Model Layer (신규/수정)

### model/admin.go
- `Admin` struct: ID(uint), StoreID, Username, PasswordHash, Name, BaseModel

### model/dto.go
- Request/Response DTOs (CustomerLoginRequest, AdminLoginRequest, AuthResponse, MenuCreateRequest, MenuUpdateRequest, MenuOrderRequest, OrderCreateRequest, OrderItemRequest, StatusUpdateRequest, TableSetupRequest, ErrorResponse, ErrorDetail)

### model/claims.go
- `Claims` struct: StoreID, Role, TableID, SessionID, AdminID, jwt.RegisteredClaims

### model/sse_event.go
- `SSEEvent` struct: Type, Data

---

## Service Layer

### service/auth.go
- `AuthService` interface:
  - `CustomerLogin(req CustomerLoginRequest) (*AuthResponse, *TokenPair, error)`
  - `AdminLogin(req AdminLoginRequest) (*AuthResponse, *TokenPair, error)`
  - `RefreshToken(refreshToken string) (string, error)` → new access token
  - `GenerateTokenPair(claims Claims) (*TokenPair, error)`
  - `ValidateToken(token string) (*Claims, error)`

### service/menu.go
- `MenuService` interface:
  - `GetMenusByStore(storeID string) ([]model.Menu, error)`
  - `CreateMenu(storeID string, req MenuCreateRequest) (*model.Menu, error)`
  - `UpdateMenu(id uint, req MenuUpdateRequest) (*model.Menu, error)`
  - `DeleteMenu(id uint) error`
  - `UpdateMenuOrder(items []MenuOrderRequest) error`

### service/order.go
- `OrderService` interface:
  - `CreateOrder(storeID, sessionID string, tableID uint, req OrderCreateRequest) (*model.Order, error)`
  - `GetOrdersBySession(sessionID string) ([]model.Order, error)`
  - `GetOrdersByStore(storeID string) ([]model.Order, error)`
  - `UpdateOrderStatus(id string, req StatusUpdateRequest) (*model.Order, error)`
  - `DeleteOrder(id string) error`

### service/table.go
- `TableService` interface:
  - `SetupTable(storeID string, req TableSetupRequest) (*model.Table, error)`
  - `CompleteTable(tableID uint) error`
  - `GetTableHistory(tableID uint, from, to *time.Time) ([]model.OrderHistory, error)`

### service/sse.go
- `SSEManager` interface:
  - `Subscribe(storeID string) chan SSEEvent`
  - `Unsubscribe(storeID string, ch chan SSEEvent)`
  - `Broadcast(storeID string, event SSEEvent)`

---

## Handler Layer

### handler/auth.go
- `POST /api/customer/login` → CustomerLogin
- `POST /api/admin/login` → AdminLogin
- `POST /api/auth/refresh` → RefreshToken

### handler/menu.go
- `GET /api/customer/menus` → GetMenus
- `POST /api/admin/menus` → CreateMenu
- `PUT /api/admin/menus/:id` → UpdateMenu
- `DELETE /api/admin/menus/:id` → DeleteMenu
- `PATCH /api/admin/menus/order` → UpdateMenuOrder

### handler/order.go
- `POST /api/customer/orders` → CreateOrder
- `GET /api/customer/orders` → GetCustomerOrders
- `GET /api/admin/orders` → GetAdminOrders
- `PATCH /api/admin/orders/:id/status` → UpdateOrderStatus
- `DELETE /api/admin/orders/:id` → DeleteOrder

### handler/table.go
- `POST /api/admin/tables/setup` → SetupTable
- `POST /api/admin/tables/:id/complete` → CompleteTable
- `GET /api/admin/tables/:id/history` → GetTableHistory

### handler/sse.go
- `GET /api/admin/orders/stream` → StreamOrders

---

## Middleware

### middleware/auth.go
- `AuthMiddleware(authService)` → JWT 검증, Claims context 저장
- `RequireRole(role string)` → 역할 기반 접근 제어

### middleware/logger.go
- `LoggerMiddleware()` → slog 구조화 로깅

---

## Repository Layer (신규)

### repository/admin.go
- `AdminRepository` interface:
  - `FindByStoreAndUsername(storeID, username string) (*model.Admin, error)`
  - `Create(admin *model.Admin) error`

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
