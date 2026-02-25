# TDD Code Generation Plan - Shared Package

## Unit Context
- **Workspace Root**: /Users/woo.noh/workspace/personal/aws-workshop/gowoopi
- **Code Location**: `frontend/packages/shared/`
- **Dependencies**: Backend API (Unit 2)

---

## Plan Step 0: 프로젝트 설정
- [x] package.json 생성
- [x] tsconfig.json 생성
- [x] 디렉토리 구조 생성 (api, auth, types, components, i18n, hooks)

## Plan Step 1: Types (TDD)
- [x] types/api.ts - API Request/Response 타입 (Menu, Order, Auth 등)
- [x] types/models.ts - 도메인 모델 타입 (Store, Table, Session 등)
- [x] types/index.ts - 타입 export

## Plan Step 2: API Client (TDD)
- [x] api/client.ts - Axios 기반 HTTP 클라이언트 (토큰 갱신 인터셉터)
- [x] api/auth.ts - 인증 API hooks (useCustomerLogin, useAdminLogin, useRefreshToken)
- [x] api/menu.ts - 메뉴 API hooks (useMenus, useCreateMenu, useUpdateMenu, useDeleteMenu)
- [x] api/order.ts - 주문 API hooks (useCreateOrder, useOrders, useUpdateOrderStatus)
- [x] api/table.ts - 테이블 API hooks (useSetupTable, useCompleteTable, useTableHistory)
- [x] api/index.ts - API export

## Plan Step 3: Auth Utilities (TDD)
- [x] auth/storage.ts - 토큰 저장/조회/삭제 유틸리티
- [x] auth/context.tsx - AuthContext + AuthProvider + useAuth hook
- [x] auth/index.ts - Auth export

## Plan Step 4: i18n (TDD)
- [x] i18n/locales/ko.json - 한국어 번역
- [x] i18n/locales/en.json - 영어 번역
- [x] i18n/config.tsx - i18n 설정 + useTranslation hook
- [x] i18n/index.ts - i18n export

## Plan Step 5: Hooks & Components
- [x] hooks/useSSE.ts - SSE 연결 hook
- [x] hooks/index.ts - Hooks export
- [x] components/index.ts - Components export (placeholder)
- [x] src/index.ts - 메인 export

## Plan Step 6: Integration
- [x] pnpm install 실행
- [x] typecheck 실행
- [ ] lint 실행 (eslint 미설정 - 앱에서 설정)

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
