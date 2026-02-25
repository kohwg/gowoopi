# Functional Design Plan - Unit 1: Database

## 범위
7개 엔티티(Store, Table, TableSession, Menu, Order, OrderItem, OrderHistory)의 상세 스키마 설계

---

## 질문

### Question 1: 주문 상태 값
주문 상태(Order Status)에 어떤 값들을 사용하시겠습니까?

A) 3단계: PENDING → PREPARING → COMPLETED
B) 4단계: PENDING → CONFIRMED → PREPARING → COMPLETED
C) 5단계: PENDING → CONFIRMED → PREPARING → READY → COMPLETED
D) Other

[Answer]: B

### Question 2: 메뉴 카테고리 관리 방식
메뉴 카테고리를 어떻게 관리하시겠습니까?

A) Menu 테이블의 category 컬럼 (문자열, 자유 입력)
B) 별도 Category 테이블 분리 (정규화)
C) Menu 테이블의 category 컬럼 (ENUM, 고정 목록)

[Answer]: B

### Question 3: Soft Delete vs Hard Delete
메뉴/주문 삭제 시 어떤 방식을 사용하시겠습니까?

A) Soft Delete (deleted_at 컬럼, 데이터 보존)
B) Hard Delete (실제 삭제)
C) 메뉴는 Soft Delete, 주문은 Hard Delete

[Answer]: A

### Question 4: 마이그레이션 도구
데이터베이스 마이그레이션에 어떤 도구를 사용하시겠습니까?

A) GORM AutoMigrate (Go 코드 기반, 개발 편의)
B) golang-migrate (SQL 파일 기반, 버전 관리)
C) 순수 SQL 파일 (수동 실행)

[Answer]: A

---

**모든 질문에 답변 후 "완료"라고 말씀해주세요.**

---

## 생성 실행 계획
- [x] domain-entities.md - 엔티티 상세 정의 (컬럼, 타입, 제약조건)
- [x] business-rules.md - 비즈니스 규칙 및 제약조건
- [x] business-logic-model.md - 데이터 흐름 및 트랜잭션 로직

---

**Plan Status**: 완료
