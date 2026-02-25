// Domain Models
export interface Store {
  id: string;
  name: string;
}

export interface Table {
  id: number;
  storeId: string;
  tableNumber: number;
  currentSessionId?: string;
}

export interface Session {
  id: string;
  tableId: number;
  startTime: string;
  endTime?: string;
  isActive: boolean;
}

export interface Category {
  id: number;
  storeId: string;
  name: string;
  displayOrder: number;
}

export interface Menu {
  id: number;
  storeId: string;
  categoryId: number;
  name: string;
  price: number;
  description?: string;
  imageUrl?: string;
  displayOrder: number;
}

export interface OrderItem {
  id: number;
  orderId: string;
  menuId: number;
  menuName: string;
  quantity: number;
  unitPrice: number;
}

export type OrderStatus = 'pending' | 'preparing' | 'completed';

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
