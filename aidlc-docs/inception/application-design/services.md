# Services - 테이블오더 서비스

## 서비스 아키텍처 개요

```
+------------------+     +------------------+
| customer-app     |     | admin-app        |
| (Next.js)        |     | (Next.js)        |
+--------+---------+     +--------+---------+
         |                        |
         | HTTP/REST              | HTTP/REST + SSE
         |                        |
+--------v------------------------v---------+
|              Go API Server (Gin)           |
|                                            |
|  +----------+  +----------+  +----------+ |
|  | Auth MW   |  | CORS MW  |  | Logger   | |
|  +----------+  +----------+  +----------+ |
|                                            |
|  +----------+  +----------+  +----------+ |
|  | Auth     |  | Menu     |  | Order    | |
|  | Handler  |  | Handler  |  | Handler  | |
|  +----+-----+  +----+-----+  +----+-----+ |
|       |              |              |       |
|  +----v-----+  +----v-----+  +----v-----+ |
|  | Auth     |  | Menu     |  | Order    | |
|  | Service  |  | Service  |  | Service  | |
|  +----+-----+  +----+-----+  +----+-----+ |
|       |              |              |       |
|  +----v-----+  +----v-----+  +----v-----+ |
|  | Store    |  | Menu     |  | Order    | |
|  | Repo     |  | Repo     |  | Repo     | |
|  +----+-----+  +----+-----+  +----+-----+ |
|       |              |              |       |
+-------+--------------+--------------+------+
        |              |              |
+-------v--------------v--------------v------+
|              MySQL Database                 |
+--------------------------------------------+
```

---

## 서비스 정의

### SVC-001: Auth Service
**책임**: 인증 및 권한 관리
- JWT Access/Refresh Token 발급 및 검증
- bcrypt 비밀번호 해싱 및 비교
- 고객(테이블) 로그인 처리
- 관리자 로그인 처리
- 토큰 자동 갱신

**의존성**: Store Repository, Table Repository

### SVC-002: Menu Service
**책임**: 메뉴 비즈니스 로직
- 매장별 메뉴 조회 (카테고리별 정렬)
- 메뉴 CRUD (등록, 수정, 삭제)
- 메뉴 노출 순서 관리
- 필수 필드 검증

**의존성**: Menu Repository

### SVC-003: Order Service
**책임**: 주문 비즈니스 로직
- 주문 생성 (항목 포함, 트랜잭션)
- 세션별/매장별 주문 조회
- 주문 상태 변경 (대기중→준비중→완료)
- 주문 삭제 및 총액 재계산
- SSE 이벤트 발행 (주문 생성/상태 변경/삭제 시)

**의존성**: Order Repository, SSE Manager

### SVC-004: Table Service
**책임**: 테이블 및 세션 관리
- 테이블 초기 설정 (인증 정보 저장)
- 테이블 세션 시작/종료
- 이용 완료 처리 (주문→이력 이동, 리셋)
- 과거 주문 내역 조회

**의존성**: Table Repository, Order Repository

### SVC-005: SSE Manager
**책임**: 실시간 이벤트 관리
- 매장별 SSE 클라이언트 연결 관리
- 이벤트 브로드캐스트 (주문 생성, 상태 변경, 삭제)
- 연결 해제 처리
- 하트비트 전송

**의존성**: 없음 (다른 서비스에서 호출됨)

---

## 서비스 오케스트레이션 패턴

### 주문 생성 플로우
```
Customer → Handler → OrderService.CreateOrder()
  → OrderRepo.Create() (트랜잭션)
  → SSEManager.Broadcast(NewOrderEvent)
  → Response to Customer
```

### 주문 상태 변경 플로우
```
Admin → Handler → OrderService.UpdateOrderStatus()
  → OrderRepo.UpdateStatus()
  → SSEManager.Broadcast(StatusChangeEvent)
  → Response to Admin
```

### 테이블 이용 완료 플로우
```
Admin → Handler → TableService.CompleteTable()
  → OrderRepo.MoveToHistory() (트랜잭션)
  → TableRepo.EndSession()
  → SSEManager.Broadcast(TableResetEvent)
  → Response to Admin
```

### 실시간 모니터링 플로우
```
Admin → Handler → SSEManager.Subscribe(StoreID)
  → SSE Stream 유지
  → 이벤트 수신 시 자동 전송
  → 연결 해제 시 Unsubscribe
```

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
