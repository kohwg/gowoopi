# Unit 4: Admin App - TDD Code Generation Plan

## Plan Steps

### Step 0: Project Setup
- [x] Next.js 프로젝트 생성 (`frontend/apps/admin-app/`)
- [x] 의존성 설치 (HeroUI, TanStack Query, Zustand)
- [x] TypeScript, Tailwind 설정
- [x] shared 패키지 연결

### Step 1: Core Infrastructure (RED → GREEN)
- [x] SSE client (`lib/sse-client.ts`)
- [x] SSE store (`stores/sse-store.ts`)
- [x] useSSE hook (`hooks/use-sse.ts`)
- [x] 테스트: SSE 연결/이벤트 처리

### Step 2: Layout & Auth (RED → GREEN)
- [x] Root layout with Providers
- [x] Login page + LoginForm
- [x] Dashboard layout (Sidebar, Header)
- [x] Auth guard (미인증 시 리다이렉트)
- [ ] 테스트: 로그인 flow, 리다이렉트 (E2E 범위)

### Step 3: Dashboard - Order Monitoring (RED → GREEN)
- [x] OrderGrid component
- [x] TableCard component
- [x] OrderDetailModal component
- [x] StatusBadge component
- [x] SSE 연동 (실시간 업데이트)
- [ ] 테스트: 주문 표시, 상태 변경, 삭제 (E2E 범위)

### Step 4: Table Management (RED → GREEN)
- [x] TableList component (Tables page에 통합)
- [x] TableSetupModal component
- [x] TableHistoryModal component
- [x] ConfirmModal component
- [ ] 테스트: 테이블 설정, 이용 완료, 내역 조회 (E2E 범위)

### Step 5: Menu Management (RED → GREEN)
- [x] CategoryTabs component
- [x] MenuList component (Menus page에 통합)
- [x] MenuFormModal component
- [ ] 테스트: 메뉴 CRUD, 순서 변경 (E2E 범위)

### Step 6: i18n Integration
- [x] 한국어/영어 번역 파일 (admin 전용)
- [x] 언어 선택 UI
- [ ] 테스트: 언어 전환 (E2E 범위)

### Step 7: Final Verification
- [x] 전체 lint 통과
- [x] 전체 테스트 통과 (3 unit tests)
- [x] 빌드 성공

---

## Test Strategy

| 영역 | 테스트 유형 | 도구 |
|------|------------|------|
| Components | Unit | Vitest + React Testing Library |
| Hooks | Unit | Vitest + @testing-library/react-hooks |
| Stores | Unit | Vitest |
| Pages | Integration | Vitest + RTL |

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
