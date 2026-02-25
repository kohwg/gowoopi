# Business Rules - Unit 1: Database

## 주문 상태 전이 규칙
| 현재 상태 | 허용 전이 | 설명 |
|-----------|-----------|------|
| PENDING | CONFIRMED | 관리자 주문 확인 |
| PENDING | (삭제) | 관리자 주문 삭제 (Soft Delete) |
| CONFIRMED | PREPARING | 조리 시작 |
| PREPARING | COMPLETED | 조리 완료 |
| COMPLETED | - | 최종 상태 (변경 불가) |

**규칙**: 역방향 전이 불가. 상태 건너뛰기 불가.

---

## 데이터 무결성 규칙

### Store
- admin_username: 매장 내 유일
- admin_password_hash: bcrypt 해시 필수

### Category
- (store_id, name): 매장 내 카테고리명 유일 (Soft Delete 고려)
- sort_order: 매장 내 정렬 순서

### Table
- (store_id, table_number): 매장 내 테이블 번호 유일 (Soft Delete 고려)
- 활성 세션은 테이블당 최대 1개

### Menu
- price: 0 이상 정수 (원 단위)
- category_id: 유효한 카테고리 참조 필수
- Soft Delete 시 기존 주문의 menu_name/price 스냅샷은 영향 없음

### Order
- total_amount: OrderItem.subtotal 합계와 일치
- 주문 삭제 시 Soft Delete (deleted_at 설정)
- 삭제된 주문은 조회에서 제외

### OrderItem
- quantity: 1 이상
- subtotal = price × quantity (자동 계산)
- menu_name, price: 주문 시점 스냅샷 (메뉴 변경 영향 없음)

### OrderHistory
- items_json: 주문 항목 전체를 JSON으로 보존
- 이력 데이터는 수정/삭제 불가 (append-only)

---

## Soft Delete 규칙
- 대상: Store, Category, Table, Menu, Order
- 비대상: TableSession, OrderItem, OrderHistory
- GORM의 `gorm.DeletedAt` 타입 사용
- 조회 시 기본적으로 deleted_at IS NULL 조건 적용
- UNIQUE 제약조건에 deleted_at 포함하여 삭제 후 동일 이름 재생성 허용

---

## 세션 규칙
- 테이블당 활성 세션 최대 1개
- 세션 시작: is_active = true, started_at 설정
- 세션 종료: is_active = false, ended_at 설정
- 이용 완료 시: 활성 주문 → OrderHistory 이동 후 세션 종료

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
