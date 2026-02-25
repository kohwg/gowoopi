# Unit of Work - 테이블오더 서비스

## 분해 전략
- **분해 방식**: 4 Units (customer-app, admin-app, 백엔드 API, 데이터베이스)
- **개발 순서**: 병렬 개발 (API 계약 기반)

---

## Unit 1: Database (데이터베이스)
**경로**: `database/`  
**개발 순서**: 1순위 (다른 Unit의 기반)

**책임**:
- MySQL 스키마 정의 (7개 테이블)
- 마이그레이션 스크립트
- 초기 시드 데이터

**산출물**:
- `database/migrations/` - DDL 스크립트
- `database/seeds/` - 초기 데이터

---

## Unit 2: Backend API (백엔드)
**경로**: `backend/`  
**개발 순서**: 1순위 (DB와 병렬, 프론트엔드에 API 제공)

**책임**:
- Go + Gin HTTP 서버
- Handler / Service / Repository 레이어
- JWT 인증 (Access + Refresh Token)
- SSE 실시간 이벤트
- 구조화된 로깅

**산출물**:
- `backend/cmd/server/main.go` - 엔트리포인트
- `backend/internal/` - handler, service, repository, model, middleware, sse
- `backend/go.mod`, `backend/go.sum`

---

## Unit 3: Customer App (고객용 프론트엔드)
**경로**: `frontend/apps/customer-app/`  
**개발 순서**: 2순위 (API 계약 기반 병렬 개발 가능)

**책임**:
- 자동 로그인 및 세션 관리
- 메뉴 조회 (카테고리별)
- 장바구니 관리 (Zustand)
- 주문 생성 및 확인
- 주문 내역 조회
- 다국어 (한국어 + 영어)

**산출물**:
- Next.js 앱 (TypeScript)
- HeroUI v3 + Tailwind CSS
- TanStack Query hooks
- Zustand stores

---

## Unit 4: Admin App (관리자용 프론트엔드)
**경로**: `frontend/apps/admin-app/`  
**개발 순서**: 2순위 (API 계약 기반 병렬 개발 가능)

**책임**:
- 매장 인증 (로그인)
- 실시간 주문 모니터링 (SSE)
- 주문 상태 관리
- 테이블 관리 (설정, 이용 완료, 주문 삭제, 과거 내역)
- 메뉴 관리 (CRUD, 순서)
- 다국어 (한국어 + 영어)

**산출물**:
- Next.js 앱 (TypeScript)
- HeroUI v3 + Tailwind CSS
- TanStack Query hooks + SSE 연동
- Zustand stores

---

## Shared Package (공통)
**경로**: `frontend/packages/shared/`  
**참고**: 독립 Unit이 아닌 Unit 3, 4에서 공유하는 패키지

**책임**:
- API 클라이언트 (TanStack Query hooks)
- 인증 유틸리티 (토큰 갱신)
- 공통 TypeScript 타입
- 공통 UI 컴포넌트 (HeroUI 래퍼)
- i18n 설정

---

## 코드 구성 전략 (Greenfield)

```
gowoopi/
├── frontend/
│   ├── apps/
│   │   ├── customer-app/        # Unit 3
│   │   └── admin-app/           # Unit 4
│   ├── packages/
│   │   └── shared/              # 공통 패키지
│   ├── package.json
│   └── pnpm-workspace.yaml
├── backend/                     # Unit 2
│   ├── cmd/server/
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   ├── middleware/
│   │   └── sse/
│   ├── go.mod
│   └── go.sum
├── database/                    # Unit 1
│   ├── migrations/
│   └── seeds/
└── aidlc-docs/
```

---

**문서 버전**: 1.0  
**작성일**: 2026-02-25
