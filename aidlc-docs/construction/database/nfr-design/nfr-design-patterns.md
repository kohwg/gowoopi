# NFR Design Patterns - Unit 1: Database

## 성능 패턴

### Connection Pool 관리
- MaxOpenConns=25, MaxIdleConns=10, ConnMaxLifetime=5m
- GORM `db.DB()` 통해 `*sql.DB` 레벨에서 설정
- 커넥션 재사용으로 연결 오버헤드 최소화

### 인덱스 전략
- 복합 인덱스: 자주 사용되는 WHERE 조건 조합
- Soft Delete 인덱스: deleted_at 포함하여 필터링 성능 보장
- 커버링 인덱스: 조회 빈도 높은 쿼리 대상

### 쿼리 최적화
- GORM Preload: N+1 문제 방지 (Order → OrderItem 관계)
- 필요한 컬럼만 SELECT (Select 메서드 활용)
- Pagination: LIMIT/OFFSET 기반

---

## 보안 패턴

### 비밀번호 저장
- bcrypt (cost=10): 단방향 해시
- 평문 비밀번호 메모리 즉시 해제

### SQL Injection 방지
- GORM 파라미터 바인딩 (Prepared Statement)
- Raw SQL 사용 금지 (불가피 시 `?` 플레이스홀더)

### DB 접속 정보
- 환경변수: DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
- 코드에 하드코딩 금지

---

## 데이터 무결성 패턴

### Soft Delete
- GORM `gorm.DeletedAt` 타입 사용
- 기본 조회 시 자동 `WHERE deleted_at IS NULL` 적용
- `Unscoped()` 로 삭제된 데이터 포함 조회 가능

### 트랜잭션 관리
- GORM `db.Transaction()` 사용
- 주문 생성, 이용 완료, 주문 삭제: 트랜잭션 필수
- 실패 시 자동 롤백

### 외래키 제약조건
- GORM 태그로 FK 정의
- CASCADE 미사용 (애플리케이션 레벨에서 관리)

---

## 가용성 패턴

### AutoMigrate
- 서버 시작 시 스키마 자동 동기화
- 기존 데이터 보존 (컬럼 추가만, 삭제 안 함)

### 시드 데이터
- 개발 환경: 초기 매장/테이블/메뉴 데이터 자동 생성
- 환경변수 플래그로 시드 실행 제어

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
