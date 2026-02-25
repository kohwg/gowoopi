# Business Logic Model - Unit 1: Database

## 트랜잭션 경계

### TX-1: 주문 생성
```
BEGIN TRANSACTION
  1. INSERT Order (UUID, session_id, store_id, table_id, status=PENDING)
  2. INSERT OrderItem[] (menu_name/price 스냅샷)
  3. UPDATE Order.total_amount = SUM(OrderItem.subtotal)
COMMIT
```

### TX-2: 테이블 이용 완료
```
BEGIN TRANSACTION
  1. SELECT 활성 세션의 모든 Order + OrderItem
  2. INSERT OrderHistory[] (items_json으로 항목 직렬화)
  3. Soft DELETE Order (deleted_at 설정)
  4. UPDATE TableSession (is_active=false, ended_at=NOW)
COMMIT
```

### TX-3: 주문 삭제
```
BEGIN TRANSACTION
  1. Soft DELETE Order (deleted_at 설정)
  2. (OrderItem은 Order 조회 시 자동 제외)
COMMIT
```

---

## 인덱스 전략

| 테이블 | 인덱스 | 컬럼 | 용도 |
|--------|--------|------|------|
| Store | idx_store_deleted | deleted_at | Soft Delete 필터 |
| Category | idx_category_store | store_id, deleted_at | 매장별 카테고리 조회 |
| Table | uk_table_store_num | store_id, table_number, deleted_at | 유니크 + 조회 |
| TableSession | idx_session_table_active | table_id, is_active | 활성 세션 조회 |
| Menu | idx_menu_store | store_id, deleted_at | 매장별 메뉴 조회 |
| Menu | idx_menu_category | category_id | 카테고리별 메뉴 |
| Order | idx_order_session | session_id, deleted_at | 세션별 주문 조회 |
| Order | idx_order_store | store_id, deleted_at | 매장별 주문 조회 |
| OrderItem | idx_orderitem_order | order_id | 주문별 항목 조회 |
| OrderHistory | idx_history_store | store_id | 매장별 이력 |
| OrderHistory | idx_history_table | table_id | 테이블별 이력 |

---

## 마이그레이션 전략
- **도구**: GORM AutoMigrate
- **실행**: 서버 시작 시 자동 마이그레이션
- **시드 데이터**: 개발용 초기 데이터 (Go 코드)

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
