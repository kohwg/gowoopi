# Test Plan - Unit 1: Database

## Unit Overview
- **Unit**: Database
- **Stories**: 전체 (데이터 기반)
- **Tech**: Go, GORM, MySQL 8.0, testify

---

## Model Layer Tests

### model validation
- **TC-DB-001**: Store 모델 UUID PK 생성 확인
  - Given: Store 생성 시
  - When: ID가 비어있으면
  - Then: UUID가 자동 생성되어야 한다
  - Status: ⬜ Not Started

- **TC-DB-002**: OrderStatus 유효성 검증
  - Given: OrderStatus 값이 주어졌을 때
  - When: IsValid() 호출 시
  - Then: PENDING/CONFIRMED/PREPARING/COMPLETED만 true
  - Status: ⬜ Not Started

- **TC-DB-003**: OrderStatus 전이 검증
  - Given: 현재 상태가 주어졌을 때
  - When: CanTransitionTo() 호출 시
  - Then: 허용된 전이만 true (PENDING→CONFIRMED, CONFIRMED→PREPARING, PREPARING→COMPLETED)
  - Status: ⬜ Not Started

---

## Repository Layer Tests (Integration - MySQL)

### StoreRepository
- **TC-DB-010**: FindByID - 존재하는 매장 조회
  - Given: 매장이 DB에 존재할 때
  - When: FindByID 호출
  - Then: 매장 정보 반환
  - Status: ⬜ Not Started

- **TC-DB-011**: FindByID - 존재하지 않는 매장
  - Given: 매장이 DB에 없을 때
  - When: FindByID 호출
  - Then: nil, error 반환
  - Status: ⬜ Not Started

- **TC-DB-012**: FindByIDAndUsername - 관리자 인증
  - Given: 매장+관리자 정보가 DB에 존재할 때
  - When: FindByIDAndUsername 호출
  - Then: 매장 정보 반환
  - Status: ⬜ Not Started

### CategoryRepository
- **TC-DB-020**: Create - 카테고리 생성
  - Given: 유효한 카테고리 데이터
  - When: Create 호출
  - Then: 카테고리 생성, ID 할당
  - Status: ⬜ Not Started

- **TC-DB-021**: FindByStore - 매장별 카테고리 조회
  - Given: 매장에 카테고리가 존재할 때
  - When: FindByStore 호출
  - Then: sort_order 순으로 카테고리 목록 반환
  - Status: ⬜ Not Started

- **TC-DB-022**: Delete - Soft Delete 확인
  - Given: 카테고리가 존재할 때
  - When: Delete 호출
  - Then: deleted_at 설정, 일반 조회에서 제외
  - Status: ⬜ Not Started

### TableRepository
- **TC-DB-030**: Create - 테이블 생성
  - Given: 유효한 테이블 데이터
  - When: Create 호출
  - Then: 테이블 생성
  - Status: ⬜ Not Started

- **TC-DB-031**: FindByStoreAndNumber - 테이블 조회
  - Given: 매장+테이블번호가 존재할 때
  - When: FindByStoreAndNumber 호출
  - Then: 테이블 정보 반환
  - Status: ⬜ Not Started

### SessionRepository
- **TC-DB-040**: Create - 세션 생성
  - Given: 유효한 세션 데이터
  - When: Create 호출
  - Then: 세션 생성, UUID ID 할당
  - Status: ⬜ Not Started

- **TC-DB-041**: FindActiveByTable - 활성 세션 조회
  - Given: 테이블에 활성 세션이 존재할 때
  - When: FindActiveByTable 호출
  - Then: 활성 세션 반환
  - Status: ⬜ Not Started

- **TC-DB-042**: End - 세션 종료
  - Given: 활성 세션이 존재할 때
  - When: End 호출
  - Then: is_active=false, ended_at 설정
  - Status: ⬜ Not Started

### MenuRepository
- **TC-DB-050**: Create - 메뉴 생성
  - Given: 유효한 메뉴 데이터
  - When: Create 호출
  - Then: 메뉴 생성
  - Status: ⬜ Not Started

- **TC-DB-051**: FindByStore - 매장별 메뉴 조회
  - Given: 매장에 메뉴가 존재할 때
  - When: FindByStore 호출
  - Then: 메뉴 목록 반환 (Category Preload)
  - Status: ⬜ Not Started

- **TC-DB-052**: Delete - Soft Delete 확인
  - Given: 메뉴가 존재할 때
  - When: Delete 호출
  - Then: deleted_at 설정
  - Status: ⬜ Not Started

- **TC-DB-053**: UpdateOrder - 순서 변경
  - Given: 메뉴 목록이 존재할 때
  - When: UpdateOrder 호출
  - Then: sort_order 업데이트
  - Status: ⬜ Not Started

### OrderRepository
- **TC-DB-060**: Create - 주문+항목 트랜잭션 생성
  - Given: 유효한 주문+항목 데이터
  - When: Create 호출
  - Then: Order + OrderItem 모두 생성, total_amount 계산
  - Status: ⬜ Not Started

- **TC-DB-061**: FindBySession - 세션별 주문 조회
  - Given: 세션에 주문이 존재할 때
  - When: FindBySession 호출
  - Then: 주문 목록 반환 (OrderItem Preload)
  - Status: ⬜ Not Started

- **TC-DB-062**: FindByStore - 매장별 활성 주문 조회
  - Given: 매장에 활성 주문이 존재할 때
  - When: FindByStore 호출
  - Then: deleted_at IS NULL인 주문만 반환
  - Status: ⬜ Not Started

- **TC-DB-063**: UpdateStatus - 상태 변경
  - Given: 주문이 존재할 때
  - When: UpdateStatus 호출
  - Then: 상태 변경
  - Status: ⬜ Not Started

- **TC-DB-064**: Delete - Soft Delete
  - Given: 주문이 존재할 때
  - When: Delete 호출
  - Then: deleted_at 설정
  - Status: ⬜ Not Started

- **TC-DB-065**: MoveToHistory - 이력 이동 트랜잭션
  - Given: 세션에 주문이 존재할 때
  - When: MoveToHistory 호출
  - Then: OrderHistory 생성, 원본 Order Soft Delete
  - Status: ⬜ Not Started

- **TC-DB-066**: FindHistory - 이력 조회
  - Given: 테이블에 이력이 존재할 때
  - When: FindHistory 호출
  - Then: 날짜 필터 적용된 이력 반환
  - Status: ⬜ Not Started

---

## Database Layer Tests

### database.go
- **TC-DB-070**: Connect - DB 연결 성공
  - Given: 유효한 DB 설정
  - When: Connect 호출
  - Then: *gorm.DB 반환, Connection Pool 설정 확인
  - Status: ⬜ Not Started

- **TC-DB-071**: AutoMigrate - 스키마 마이그레이션
  - Given: DB 연결 상태
  - When: AutoMigrate 호출
  - Then: 모든 테이블 생성 확인
  - Status: ⬜ Not Started

---

## Requirements Coverage
| 요구사항 | 테스트 케이스 | Status |
|----------|-------------|--------|
| 데이터 무결성 | TC-DB-001~003, TC-DB-060, TC-DB-065 | ⬜ Pending |
| Soft Delete | TC-DB-022, TC-DB-052, TC-DB-064 | ⬜ Pending |
| 트랜잭션 | TC-DB-060, TC-DB-065 | ⬜ Pending |
| 세션 관리 | TC-DB-040~042 | ⬜ Pending |
| DB 연결 | TC-DB-070~071 | ⬜ Pending |

---

**Total Test Cases**: 25  
**문서 버전**: 1.0  
**작성일**: 2026-02-25
