# Components - 테이블오더 서비스

## 프로젝트 구조 개요

```
gowoopi/
├── frontend/                    # 모노레포 (pnpm workspace)
│   ├── apps/
│   │   ├── customer-app/        # 고객용 Next.js 앱
│   │   └── admin-app/           # 관리자용 Next.js 앱
│   ├── packages/
│   │   └── shared/              # 공통 컴포넌트/유틸리티
│   ├── package.json
│   └── pnpm-workspace.yaml
├── backend/                     # Go 백엔드
│   ├── cmd/server/              # 엔트리포인트
│   ├── internal/
│   │   ├── handler/             # HTTP 핸들러
│   │   ├── service/             # 비즈니스 로직
│   │   ├── repository/          # 데이터 접근
│   │   ├── model/               # 도메인 모델
│   │   ├── middleware/          # 미들웨어
│   │   └── sse/                 # SSE 관리
│   ├── go.mod
│   └── go.sum
└── database/                    # DB 마이그레이션
    └── migrations/
```

---

## 프론트엔드 컴포넌트

### FE-001: customer-app (고객용 Next.js 앱)
**책임**: 고객이 테이블에서 메뉴를 탐색하고 주문하는 인터페이스
- 자동 로그인 및 세션 관리
- 메뉴 조회 및 카테고리 탐색
- 장바구니 관리 (Zustand 로컬 상태)
- 주문 생성 및 확인
- 주문 내역 조회

### FE-002: admin-app (관리자용 Next.js 앱)
**책임**: 매장 관리자가 주문, 테이블, 메뉴를 관리하는 인터페이스
- 매장 인증 (로그인)
- 실시간 주문 모니터링 대시보드 (SSE)
- 주문 상태 관리
- 테이블 관리 (설정, 이용 완료, 주문 삭제)
- 메뉴 관리 (CRUD, 순서 조정)
- 과거 주문 내역 조회

### FE-003: shared (공통 패키지)
**책임**: 두 앱에서 공유하는 코드
- API 클라이언트 (TanStack Query hooks)
- 인증 유틸리티 (토큰 갱신 로직)
- 공통 타입 정의 (TypeScript interfaces)
- 공통 UI 컴포넌트 (HeroUI 래퍼)
- 다국어 설정 (i18n)

---

## 백엔드 컴포넌트

### BE-001: Handler Layer (HTTP 핸들러)
**책임**: HTTP 요청 수신, 검증, 응답 반환
- `handler/auth.go` - 인증 관련 핸들러 (로그인, 토큰 갱신)
- `handler/menu.go` - 메뉴 관련 핸들러 (CRUD)
- `handler/order.go` - 주문 관련 핸들러 (생성, 조회, 상태 변경, 삭제)
- `handler/table.go` - 테이블 관련 핸들러 (설정, 이용 완료, 내역)
- `handler/sse.go` - SSE 스트림 핸들러

### BE-002: Service Layer (비즈니스 로직)
**책임**: 비즈니스 규칙 실행, 트랜잭션 관리
- `service/auth.go` - 인증 서비스 (JWT 발급/검증, bcrypt)
- `service/menu.go` - 메뉴 서비스 (CRUD, 순서 관리)
- `service/order.go` - 주문 서비스 (생성, 상태 변경, 삭제, 이력 이동)
- `service/table.go` - 테이블 서비스 (세션 관리, 이용 완료)
- `service/sse.go` - SSE 이벤트 브로드캐스트 서비스

### BE-003: Repository Layer (데이터 접근)
**책임**: 데이터베이스 CRUD 작업, GORM 사용
- `repository/store.go` - Store 엔티티 접근
- `repository/table.go` - Table, TableSession 엔티티 접근
- `repository/menu.go` - Menu 엔티티 접근
- `repository/order.go` - Order, OrderItem, OrderHistory 엔티티 접근

### BE-004: Model Layer (도메인 모델)
**책임**: 데이터 구조 정의
- `model/store.go` - Store 모델
- `model/table.go` - Table, TableSession 모델
- `model/menu.go` - Menu 모델
- `model/order.go` - Order, OrderItem, OrderHistory 모델
- `model/auth.go` - JWT Claims, 로그인 요청/응답 모델

### BE-005: Middleware Layer
**책임**: 요청 전처리 (인증, 로깅, CORS)
- `middleware/auth.go` - JWT 인증 미들웨어
- `middleware/cors.go` - CORS 설정
- `middleware/logger.go` - 구조화된 로깅

### BE-006: SSE Manager
**책임**: Server-Sent Events 연결 관리 및 이벤트 브로드캐스트
- `sse/manager.go` - SSE 연결 관리, 이벤트 발행

---

## 데이터베이스 컴포넌트

### DB-001: MySQL Database
**책임**: 데이터 영구 저장
- Store, Table, TableSession, Menu, Order, OrderItem, OrderHistory 테이블
- 마이그레이션 스크립트

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
