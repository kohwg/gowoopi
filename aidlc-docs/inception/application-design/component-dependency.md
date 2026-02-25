# Component Dependencies - 테이블오더 서비스

## 의존성 매트릭스

### 프론트엔드 의존성

| 컴포넌트 | 의존 대상 | 통신 방식 |
|----------|-----------|-----------|
| customer-app | shared | 패키지 import |
| customer-app | Go API Server | HTTP/REST (TanStack Query) |
| admin-app | shared | 패키지 import |
| admin-app | Go API Server | HTTP/REST (TanStack Query) + SSE |
| shared | - | 없음 (공통 라이브러리) |

### 백엔드 의존성

| 컴포넌트 | 의존 대상 | 관계 |
|----------|-----------|------|
| Handler | Service | 호출 (비즈니스 로직 위임) |
| Handler | Middleware | 요청 전처리 |
| Service (Auth) | Repository (Store, Table) | 데이터 조회 |
| Service (Menu) | Repository (Menu) | 데이터 CRUD |
| Service (Order) | Repository (Order) | 데이터 CRUD |
| Service (Order) | SSE Manager | 이벤트 발행 |
| Service (Table) | Repository (Table, Order) | 데이터 CRUD |
| Repository | GORM + MySQL | 데이터베이스 접근 |

---

## 통신 패턴

### 프론트엔드 → 백엔드
- **프로토콜**: HTTP/REST (JSON)
- **인증**: JWT (HTTP-only Cookie)
- **실시간**: SSE (관리자 대시보드)
- **라이브러리**: TanStack Query (캐싱, 재시도, 무효화)

### 백엔드 내부
- **Handler → Service**: 직접 함수 호출 (Go interface)
- **Service → Repository**: 직접 함수 호출 (Go interface)
- **Service → SSE Manager**: 직접 함수 호출

### 백엔드 → 데이터베이스
- **ORM**: GORM
- **연결**: Connection Pool
- **트랜잭션**: Service 레이어에서 관리

---

## 데이터 흐름

### 고객 주문 플로우
```
customer-app → POST /api/customer/orders
  → Auth Middleware (JWT 검증)
  → Order Handler (요청 검증)
  → Order Service (비즈니스 로직)
  → Order Repository (DB 저장, 트랜잭션)
  → SSE Manager (이벤트 브로드캐스트)
  → admin-app (SSE 수신, 대시보드 업데이트)
```

### 인증 플로우
```
customer-app/admin-app → POST /api/*/login
  → Auth Handler (요청 검증)
  → Auth Service (bcrypt 비교, JWT 생성)
  → Store/Table Repository (인증 정보 조회)
  → Set-Cookie (Access Token + Refresh Token)
```

### 토큰 갱신 플로우
```
shared (interceptor) → POST /api/auth/refresh
  → Auth Handler
  → Auth Service (Refresh Token 검증, 새 Access Token 생성)
  → Set-Cookie (New Access Token)
```

---

## 기술 스택 요약

| 레이어 | 기술 | 역할 |
|--------|------|------|
| 프론트엔드 프레임워크 | Next.js | SSR/CSR, 라우팅 |
| UI 라이브러리 | HeroUI v3 + Tailwind CSS | 디자인 시스템 |
| 상태 관리 | Zustand | 클라이언트 상태 (장바구니 등) |
| 서버 상태 | TanStack Query | API 캐싱, 동기화 |
| 언어 | TypeScript | 타입 안전성 |
| 백엔드 프레임워크 | Go + Gin | HTTP 서버 |
| ORM | GORM | 데이터베이스 접근 |
| 데이터베이스 | MySQL | 데이터 저장 |
| 인증 | JWT + bcrypt | 보안 |
| 실시간 | SSE | 주문 모니터링 |
| 다국어 | i18n | 한국어 + 영어 |

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
