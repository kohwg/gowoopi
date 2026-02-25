# Unit 4: Admin App - Component Design

## 컴포넌트 계층

```
App
├── Providers (QueryClient, Auth, I18n, HeroUI)
├── LoginPage
│   └── LoginForm
└── DashboardLayout
    ├── Sidebar
    ├── Header (언어 선택, 로그아웃)
    ├── DashboardPage
    │   ├── OrderGrid
    │   │   └── TableCard[]
    │   └── OrderDetailModal
    ├── TablesPage
    │   ├── TableList
    │   ├── TableSetupModal
    │   └── TableHistoryModal
    └── MenusPage
        ├── CategoryTabs
        ├── MenuList
        └── MenuFormModal
```

---

## 컴포넌트 상세

### Layout Components

#### Sidebar
```typescript
interface SidebarProps {
  currentPath: string;
}
// 네비게이션: 대시보드, 테이블 관리, 메뉴 관리
// HeroUI: Listbox
```

#### Header
```typescript
interface HeaderProps {
  storeName: string;
}
// 매장명, 언어 선택 (Dropdown), 로그아웃 버튼
// HeroUI: Navbar, Dropdown, Button
```

---

### Order Components

#### OrderGrid
```typescript
interface OrderGridProps {
  orders: Order[];
  onOrderClick: (order: Order) => void;
}
// 테이블별로 그룹화하여 TableCard 렌더링
// SSE로 실시간 업데이트
```

#### TableCard
```typescript
interface TableCardProps {
  tableNumber: number;
  orders: Order[];
  totalAmount: number;
  isNew?: boolean;
  onClick: () => void;
}
// HeroUI: Card
// 신규 주문 시 isNew=true → 강조 애니메이션
```

#### OrderDetailModal
```typescript
interface OrderDetailModalProps {
  order: Order | null;
  isOpen: boolean;
  onClose: () => void;
  onStatusChange: (status: OrderStatus) => void;
  onDelete: () => void;
}
// HeroUI: Modal, Table, Select, Button
// 메뉴 목록, 상태 변경, 삭제 기능
```

#### StatusBadge
```typescript
interface StatusBadgeProps {
  status: OrderStatus;
}
// HeroUI: Chip
// pending: warning, preparing: primary, completed: success
```

---

### Table Components

#### TableList
```typescript
interface TableListProps {
  tables: TableWithOrders[];
  onSetup: () => void;
  onComplete: (tableId: number) => void;
  onHistory: (tableId: number) => void;
}
// HeroUI: Table, Button
```

#### TableSetupModal
```typescript
interface TableSetupModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: TableSetupRequest) => void;
}
// HeroUI: Modal, Input, Button
// 테이블 번호, 비밀번호 입력
```

#### TableHistoryModal
```typescript
interface TableHistoryModalProps {
  tableId: number;
  isOpen: boolean;
  onClose: () => void;
}
// HeroUI: Modal, Table, DatePicker
// 날짜 필터, 과거 주문 목록
```

---

### Menu Components

#### CategoryTabs
```typescript
interface CategoryTabsProps {
  categories: Category[];
  selectedId: number;
  onSelect: (id: number) => void;
}
// HeroUI: Tabs
```

#### MenuList
```typescript
interface MenuListProps {
  menus: Menu[];
  onEdit: (menu: Menu) => void;
  onDelete: (id: number) => void;
  onReorder: (items: MenuOrderItem[]) => void;
}
// HeroUI: Table, Button
// 드래그 앤 드롭 또는 위/아래 버튼
```

#### MenuFormModal
```typescript
interface MenuFormModalProps {
  menu?: Menu;  // undefined = 등록, defined = 수정
  categories: Category[];
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: MenuCreateRequest | MenuUpdateRequest) => void;
}
// HeroUI: Modal, Input, Textarea, Select, Button
```

---

### Shared Components

#### ConfirmModal
```typescript
interface ConfirmModalProps {
  title: string;
  message: string;
  isOpen: boolean;
  onConfirm: () => void;
  onCancel: () => void;
}
// HeroUI: Modal, Button
// 삭제, 이용 완료 등 확인용
```

---

**문서 버전**: 1.0
**작성일**: 2026-02-25
