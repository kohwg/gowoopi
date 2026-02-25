// Domain Models
export interface Store {
  id: string;
  name: string;
}

export interface Table {
  id: number;
  storeId: string;
  tableNumber: number;
  isActive: boolean;
}

export interface Session {
  id: string;
  tableId: number;
  storeId: string;
  startedAt: string;
  endedAt?: string;
  isActive: boolean;
}

export interface Category {
  id: number;
  storeId: string;
  name: string;
  sortOrder: number;
}

export interface Menu {
  id: number;
  storeId: string;
  categoryId: number;
  name: string;
  price: number;
  description?: string;
  imageUrl?: string;
  isAvailable: boolean;
  sortOrder: number;
  category?: { id: number; name: string };
}

export interface OrderItem {
  id: number;
  orderId: string;
  menuId: number;
  menuName: string;
  quantity: number;
  price: number;
  subtotal: number;
}

export type OrderStatus = 'PENDING' | 'CONFIRMED' | 'PREPARING' | 'COMPLETED';

export interface Order {
  id: string;
  sessionId: string;
  tableId: number;
  storeId: string;
  totalAmount: number;
  status: OrderStatus;
  createdAt: string;
  items: OrderItem[];
}

export interface OrderHistory extends Order {
  completedAt: string;
}

export type UserRole = 'customer' | 'admin';
