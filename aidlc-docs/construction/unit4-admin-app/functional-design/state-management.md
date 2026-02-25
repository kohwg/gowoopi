# Unit 4: Admin App - State Management

## Zustand Stores

### SSE Store (`stores/sse-store.ts`)

```typescript
interface SSEState {
  isConnected: boolean;
  lastEvent: SSEEvent | null;
  newOrderIds: Set<string>;  // 신규 주문 강조용
  
  // Actions
  setConnected: (connected: boolean) => void;
  addNewOrder: (orderId: string) => void;
  clearNewOrder: (orderId: string) => void;
  handleEvent: (event: SSEEvent) => void;
}
```

**용도**:
- SSE 연결 상태 관리
- 신규 주문 ID 추적 (강조 표시용)
- 이벤트 처리

---

## SSE Event Types

```typescript
type SSEEventType = 'order_created' | 'order_updated' | 'order_deleted';

interface SSEEvent {
  type: SSEEventType;
  data: Order;
}
```

---

## SSE Hook (`hooks/use-sse.ts`)

```typescript
function useSSE(storeId: string): void {
  // 1. EventSource 연결 (/api/admin/orders/stream)
  // 2. 이벤트 수신 시 SSE store 업데이트
  // 3. TanStack Query 캐시 무효화
  // 4. 연결 끊김 시 자동 재연결 (3초 후)
  // 5. 컴포넌트 언마운트 시 연결 해제
}
```

---

## TanStack Query 연동

### Query Invalidation on SSE Events

```typescript
// SSE 이벤트 수신 시
switch (event.type) {
  case 'order_created':
  case 'order_updated':
  case 'order_deleted':
    queryClient.invalidateQueries({ queryKey: orderKeys.admin() });
    break;
}
```

### 사용하는 Shared Hooks

| Hook | 용도 |
|------|------|
| `useAdminLogin` | 로그인 |
| `useAdminOrders` | 주문 목록 조회 |
| `useUpdateOrderStatus` | 주문 상태 변경 |
| `useDeleteOrder` | 주문 삭제 |
| `useSetupTable` | 테이블 설정 |
| `useCompleteTable` | 테이블 이용 완료 |
| `useTableHistory` | 과거 내역 조회 |
| `useMenus` | 메뉴 목록 조회 |
| `useCreateMenu` | 메뉴 등록 |
| `useUpdateMenu` | 메뉴 수정 |
| `useDeleteMenu` | 메뉴 삭제 |
| `useUpdateMenuOrder` | 메뉴 순서 변경 |

---

## Auth Context 사용

```typescript
// 로그인 성공 시
const { login } = useAuth();
login(authResponse);

// 인증 상태 확인
const { auth, isAuthenticated } = useAuth();
if (!isAuthenticated) redirect('/login');

// 로그아웃
const { logout } = useAuth();
logout();
```

---

## i18n 사용

```typescript
const { t, locale, setLocale } = useTranslation();

// 텍스트 번역
<Button>{t('common.login')}</Button>

// 언어 변경
<Dropdown onSelectionChange={(key) => setLocale(key as Locale)}>
```

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
