# TDD Code Generation Plan - Unit 1: Database

## Unit Context
- **Workspace Root**: /Users/woo.noh/workspace/personal/aws-workshop/gowoopi
- **Project Type**: Greenfield (multi-unit monolith)
- **Code Location**: `backend/`

---

## Plan Step 0: Project Setup + Contract Skeletons
- [x] Go module 초기화 (`backend/go.mod`)
- [x] 디렉토리 구조 생성
- [x] 모든 모델 struct 스켈레톤 생성 (NotImplemented 아닌 빈 struct)
- [x] 모든 Repository interface 스켈레톤 생성
- [x] Database 패키지 스켈레톤 생성
- [ ] 컴파일 확인 (Go 미설치 - 설치 후 확인)

## Plan Step 1: Model Layer (TDD)
- [x] model/order.go - OrderStatus 타입 + IsValid() - RED-GREEN-REFACTOR (TC-DB-002)
- [x] model/order.go - CanTransitionTo() - RED-GREEN-REFACTOR (TC-DB-003)
- [x] model/store.go - BeforeCreate UUID 생성 훅 - RED-GREEN-REFACTOR (TC-DB-001)

## Plan Step 2: Database Layer (TDD)
- [x] database/database.go - Connect() - RED-GREEN-REFACTOR (TC-DB-070)
- [x] database/database.go - AutoMigrate() - RED-GREEN-REFACTOR (TC-DB-071)
- [x] database/seed.go - Seed()

## Plan Step 3: Repository Layer (TDD)
- [x] repository/store.go - FindByID, FindByIDAndUsername - RED-GREEN-REFACTOR (TC-DB-010~012)
- [x] repository/category.go - Create, FindByStore, Delete - RED-GREEN-REFACTOR (TC-DB-020~022)
- [x] repository/table.go - Create, FindByStoreAndNumber - RED-GREEN-REFACTOR (TC-DB-030~031)
- [x] repository/session.go - Create, FindActiveByTable, End - RED-GREEN-REFACTOR (TC-DB-040~042)
- [x] repository/menu.go - Create, FindByStore, Delete, UpdateOrder - RED-GREEN-REFACTOR (TC-DB-050~053)
- [x] repository/order.go - Create (트랜잭션) - RED-GREEN-REFACTOR (TC-DB-060)
- [x] repository/order.go - FindBySession, FindByStore, UpdateStatus, Delete - RED-GREEN-REFACTOR (TC-DB-061~064)
- [x] repository/order.go - MoveToHistory (트랜잭션), FindHistory - RED-GREEN-REFACTOR (TC-DB-065~066)

## Plan Step 4: Additional Artifacts
- [x] 테스트 헬퍼 (DB 테스트용 setup/teardown)
- [ ] 코드 문서 요약 (`aidlc-docs/construction/database/code/code-summary.md`)

---

**Total Steps**: 4 Plan Steps, 17 sub-items  
**Total Test Cases**: 25  
**문서 버전**: 1.0  
**작성일**: 2026-02-25
