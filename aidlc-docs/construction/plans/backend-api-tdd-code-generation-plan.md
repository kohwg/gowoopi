# TDD Code Generation Plan - Unit 2: Backend API

## Unit Context
- **Workspace Root**: /Users/woo.noh/workspace/personal/aws-workshop/gowoopi
- **Code Location**: `backend/`
- **Dependencies**: Unit 1 (model, repository)

---

## Plan Step 0: 신규 모델 + Repository + 스켈레톤
- [ ] model/admin.go - Admin 모델
- [ ] model/dto.go - Request/Response DTOs
- [ ] model/claims.go - JWT Claims
- [ ] model/sse_event.go - SSE 이벤트
- [ ] model/errors.go - 커스텀 에러 타입
- [ ] repository/admin.go 인터페이스 추가 + impl/admin.go 구현
- [ ] service/ 인터페이스 스켈레톤 (auth, menu, order, table, sse)
- [ ] handler/ 스켈레톤 (auth, menu, order, table, sse)
- [ ] middleware/ 스켈레톤 (auth, logger)
- [ ] cmd/server/main.go - 엔트리포인트 스켈레톤
- [ ] 컴파일 확인

## Plan Step 1: Service Layer (TDD)
- [ ] service/auth.go - CustomerLogin, AdminLogin, RefreshToken, GenerateTokenPair, ValidateToken
- [ ] service/menu.go - GetMenusByStore, CreateMenu, UpdateMenu, DeleteMenu, UpdateMenuOrder
- [ ] service/order.go - CreateOrder, GetOrdersBySession, GetOrdersByStore, UpdateOrderStatus, DeleteOrder
- [ ] service/table.go - SetupTable, CompleteTable, GetTableHistory
- [ ] service/sse.go - Subscribe, Unsubscribe, Broadcast

## Plan Step 2: Middleware (TDD)
- [ ] middleware/auth.go - AuthMiddleware, RequireRole
- [ ] middleware/logger.go - LoggerMiddleware

## Plan Step 3: Handler Layer (TDD)
- [ ] handler/auth.go - CustomerLogin, AdminLogin, RefreshToken
- [ ] handler/menu.go - GetMenus, CreateMenu, UpdateMenu, DeleteMenu, UpdateMenuOrder
- [ ] handler/order.go - CreateOrder, GetCustomerOrders, GetAdminOrders, UpdateOrderStatus, DeleteOrder
- [ ] handler/table.go - SetupTable, CompleteTable, GetTableHistory
- [ ] handler/sse.go - StreamOrders

## Plan Step 4: Integration
- [ ] cmd/server/main.go - 라우터 설정, DI, 서버 시작
- [ ] Store 모델에서 admin 컬럼 제거 (Admin 테이블로 이전)
- [ ] AutoMigrate에 Admin 추가
- [ ] Seed에 Admin 데이터 추가
- [ ] 컴파일 + 전체 테스트 실행

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
