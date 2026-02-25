# Unit 4: Admin App - Page Structure

## 페이지 구조

```
frontend/apps/admin-app/
├── app/
│   ├── layout.tsx              # Root layout (Providers)
│   ├── page.tsx                # Redirect to /login or /dashboard
│   ├── login/
│   │   └── page.tsx            # 로그인 페이지 (US-006)
│   ├── dashboard/
│   │   ├── layout.tsx          # Dashboard layout (Sidebar, Header)
│   │   ├── page.tsx            # 주문 모니터링 (US-007, US-008)
│   │   ├── tables/
│   │   │   └── page.tsx        # 테이블 관리 (US-009)
│   │   └── menus/
│   │       └── page.tsx        # 메뉴 관리 (US-010)
│   └── globals.css
├── components/
│   ├── layout/
│   │   ├── Sidebar.tsx
│   │   └── Header.tsx
│   ├── orders/
│   │   ├── OrderGrid.tsx       # 테이블별 주문 그리드
│   │   ├── TableCard.tsx       # 테이블 카드 (주문 요약)
│   │   ├── OrderDetailModal.tsx
│   │   └── StatusBadge.tsx
│   ├── tables/
│   │   ├── TableList.tsx
│   │   ├── TableSetupModal.tsx
│   │   ├── TableHistoryModal.tsx
│   │   └── ConfirmModal.tsx
│   └── menus/
│       ├── MenuList.tsx
│       ├── MenuFormModal.tsx
│       └── CategoryTabs.tsx
├── stores/
│   └── sse-store.ts            # SSE 연결 상태 관리
├── hooks/
│   └── use-sse.ts              # SSE 연결 hook
└── lib/
    └── sse-client.ts           # SSE 클라이언트
```

---

## 페이지별 상세

### 1. 로그인 페이지 (`/login`)
**Story**: US-006

| 요소 | 설명 |
|------|------|
| 입력 필드 | Store ID (UUID), Username, Password |
| 버튼 | 로그인 |
| 에러 표시 | 인증 실패 시 에러 메시지 |
| 리다이렉트 | 성공 시 `/dashboard` |

### 2. 대시보드 - 주문 모니터링 (`/dashboard`)
**Story**: US-007, US-008

| 요소 | 설명 |
|------|------|
| OrderGrid | 테이블별 카드 그리드 |
| TableCard | 테이블 번호, 총액, 최근 주문 미리보기 |
| OrderDetailModal | 주문 상세 (메뉴 목록, 상태 변경, 삭제) |
| SSE 연동 | 실시간 주문 업데이트 |
| 신규 주문 강조 | 애니메이션/색상 강조 |

### 3. 테이블 관리 (`/dashboard/tables`)
**Story**: US-009

| 요소 | 설명 |
|------|------|
| TableList | 테이블 목록 (번호, 세션 상태, 총액) |
| TableSetupModal | 테이블 초기 설정 (번호, 비밀번호) |
| 이용 완료 버튼 | 확인 후 세션 종료 |
| TableHistoryModal | 과거 주문 내역 (날짜 필터) |

### 4. 메뉴 관리 (`/dashboard/menus`)
**Story**: US-010

| 요소 | 설명 |
|------|------|
| CategoryTabs | 카테고리별 탭 |
| MenuList | 메뉴 목록 (이름, 가격, 이미지) |
| MenuFormModal | 메뉴 등록/수정 폼 |
| 삭제 버튼 | 확인 후 삭제 |
| 순서 변경 | 드래그 앤 드롭 또는 위/아래 버튼 |

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
