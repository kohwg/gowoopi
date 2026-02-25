# Domain Entities - Unit 1: Database

## 엔티티 목록
Store, Category, Table, TableSession, Menu, Order, OrderItem, OrderHistory

---

## Store (매장)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | UUID (CHAR(36)) | PK | 매장 고유 식별자 |
| name | VARCHAR(100) | NOT NULL | 매장명 |
| admin_username | VARCHAR(50) | NOT NULL | 관리자 아이디 |
| admin_password_hash | VARCHAR(255) | NOT NULL | bcrypt 해시 |
| default_language | VARCHAR(5) | NOT NULL, DEFAULT 'ko' | 기본 언어 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |
| deleted_at | DATETIME | NULL, INDEX | Soft Delete |

---

## Category (메뉴 카테고리)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 카테고리 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| name | VARCHAR(50) | NOT NULL | 카테고리명 |
| sort_order | INT | NOT NULL, DEFAULT 0 | 정렬 순서 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |
| deleted_at | DATETIME | NULL, INDEX | Soft Delete |

**UNIQUE**: (store_id, name, deleted_at)

---

## Table (테이블)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 테이블 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| table_number | INT | NOT NULL | 테이블 번호 |
| password_hash | VARCHAR(255) | NOT NULL | 테이블 비밀번호 해시 |
| is_active | TINYINT(1) | NOT NULL, DEFAULT 1 | 활성 상태 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |
| deleted_at | DATETIME | NULL, INDEX | Soft Delete |

**UNIQUE**: (store_id, table_number, deleted_at)

---

## TableSession (테이블 세션)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | UUID (CHAR(36)) | PK | 세션 고유 식별자 |
| table_id | BIGINT UNSIGNED | FK(Table.id), NOT NULL, INDEX | 테이블 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| started_at | DATETIME | NOT NULL | 세션 시작 |
| ended_at | DATETIME | NULL | 세션 종료 |
| is_active | TINYINT(1) | NOT NULL, DEFAULT 1 | 활성 여부 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |

---

## Menu (메뉴)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 메뉴 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| category_id | BIGINT UNSIGNED | FK(Category.id), NOT NULL, INDEX | 카테고리 ID |
| name | VARCHAR(100) | NOT NULL | 메뉴명 |
| description | TEXT | NULL | 메뉴 설명 |
| price | INT UNSIGNED | NOT NULL | 가격 (원) |
| image_url | VARCHAR(500) | NULL | 이미지 URL |
| is_available | TINYINT(1) | NOT NULL, DEFAULT 1 | 판매 가능 여부 |
| sort_order | INT | NOT NULL, DEFAULT 0 | 정렬 순서 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |
| deleted_at | DATETIME | NULL, INDEX | Soft Delete |

---

## Order (주문)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | UUID (CHAR(36)) | PK | 주문 고유 식별자 |
| session_id | CHAR(36) | FK(TableSession.id), NOT NULL, INDEX | 세션 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| table_id | BIGINT UNSIGNED | FK(Table.id), NOT NULL | 테이블 ID |
| status | ENUM('PENDING','CONFIRMED','PREPARING','COMPLETED') | NOT NULL, DEFAULT 'PENDING' | 주문 상태 |
| total_amount | INT UNSIGNED | NOT NULL, DEFAULT 0 | 총 금액 |
| created_at | DATETIME | NOT NULL | 생성일시 |
| updated_at | DATETIME | NOT NULL | 수정일시 |
| deleted_at | DATETIME | NULL, INDEX | Soft Delete |

---

## OrderItem (주문 항목)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 항목 ID |
| order_id | CHAR(36) | FK(Order.id), NOT NULL, INDEX | 주문 ID |
| menu_id | BIGINT UNSIGNED | FK(Menu.id), NOT NULL | 메뉴 ID |
| menu_name | VARCHAR(100) | NOT NULL | 주문 시점 메뉴명 (스냅샷) |
| price | INT UNSIGNED | NOT NULL | 주문 시점 가격 (스냅샷) |
| quantity | INT UNSIGNED | NOT NULL | 수량 |
| subtotal | INT UNSIGNED | NOT NULL | 소계 (price × quantity) |
| created_at | DATETIME | NOT NULL | 생성일시 |

---

## OrderHistory (주문 이력)
| 컬럼 | 타입 | 제약조건 | 설명 |
|------|------|----------|------|
| id | BIGINT UNSIGNED | PK, AUTO_INCREMENT | 이력 ID |
| original_order_id | CHAR(36) | NOT NULL, INDEX | 원본 주문 ID |
| store_id | CHAR(36) | FK(Store.id), NOT NULL, INDEX | 매장 ID |
| table_id | BIGINT UNSIGNED | FK(Table.id), NOT NULL | 테이블 ID |
| table_number | INT | NOT NULL | 테이블 번호 (스냅샷) |
| session_id | CHAR(36) | NOT NULL | 원본 세션 ID |
| status | VARCHAR(20) | NOT NULL | 최종 상태 |
| total_amount | INT UNSIGNED | NOT NULL | 총 금액 |
| items_json | JSON | NOT NULL | 주문 항목 JSON 스냅샷 |
| ordered_at | DATETIME | NOT NULL | 원본 주문 일시 |
| completed_at | DATETIME | NOT NULL | 이력 이동 일시 |
| created_at | DATETIME | NOT NULL | 생성일시 |

---

## 관계도

```
Store (1) ──── (N) Category
Store (1) ──── (N) Table
Store (1) ──── (N) Order
Table (1) ──── (N) TableSession
TableSession (1) ──── (N) Order
Category (1) ──── (N) Menu
Order (1) ──── (N) OrderItem
OrderItem (N) ──── (1) Menu
Store (1) ──── (N) OrderHistory
Table (1) ──── (N) OrderHistory
```

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
