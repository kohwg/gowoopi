# TDD Code Generation Plan - Unit 3: Customer App

## Unit Context
- **Workspace Root**: /Users/woo.noh/workspace/personal/aws-workshop/gowoopi
- **Code Location**: `frontend/apps/customer-app/`
- **Dependencies**: Shared Package (@gowoopi/shared)

---

## Plan Step 0: 프로젝트 설정
- [x] Next.js 앱 생성 (create-next-app)
- [x] package.json 수정 (@gowoopi/shared 의존성)
- [x] tailwind.config.ts, postcss.config.js
- [x] HeroUI 설정

## Plan Step 1: 레이아웃 및 Provider
- [x] app/layout.tsx - RootLayout (QueryClient, AuthProvider, I18nProvider)
- [x] app/providers.tsx - Providers 컴포넌트
- [x] components/LanguageSwitcher.tsx - 언어 전환

## Plan Step 2: 인증 (자동 로그인)
- [x] app/setup/page.tsx - 초기 설정 페이지 (매장ID, 테이블번호, 비밀번호)
- [x] stores/setup.ts - Zustand 설정 저장소
- [x] hooks/useAutoLogin.ts - 자동 로그인 hook

## Plan Step 3: 메뉴 조회
- [x] app/page.tsx - 메뉴 목록 (기본 화면)
- [x] components/MenuCard.tsx - 메뉴 카드
- [x] components/CategoryTabs.tsx - 카테고리 탭

## Plan Step 4: 장바구니
- [x] stores/cart.ts - Zustand 장바구니 저장소 (localStorage persist)
- [x] components/CartDrawer.tsx - 장바구니 Drawer
- [x] components/CartItem.tsx - 장바구니 아이템
- [x] components/CartButton.tsx - 장바구니 버튼 (FAB)

## Plan Step 5: 주문
- [x] app/checkout/page.tsx - 주문 확인 페이지
- [x] app/orders/page.tsx - 주문 내역 페이지
- [x] components/OrderCard.tsx - 주문 카드

## Plan Step 6: Integration
- [x] pnpm install
- [x] pnpm build

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
