import type { Menu, Order, OrderHistory, OrderStatus, Table } from './models';

// Auth
export interface CustomerLoginRequest {
  storeId: string;
  tableNumber: number;
  password: string;
}

export interface AdminLoginRequest {
  storeId: string;
  username: string;
  password: string;
}

export interface AuthResponse {
  role: 'customer' | 'admin';
  storeId: string;
  accessToken: string;
  refreshToken: string;
}

// Menu
export interface MenuCreateRequest {
  categoryId: number;
  name: string;
  price: number;
  description?: string;
  imageUrl?: string;
  isAvailable?: boolean;
}

export interface MenuUpdateRequest {
  categoryId?: number;
  name?: string;
  price?: number;
  description?: string;
  imageUrl?: string;
  isAvailable?: boolean;
}

export interface MenuOrderItem {
  id: number;
  sortOrder: number;
}

// Order
export interface OrderItemRequest {
  menuId: number;
  quantity: number;
}

export interface OrderCreateRequest {
  items: OrderItemRequest[];
}

export interface StatusUpdateRequest {
  status: OrderStatus;
}

// Table
export interface TableSetupRequest {
  tableNumber: number;
  password: string;
}

export interface TableHistoryQuery {
  from?: string;
  to?: string;
}

// Error
export interface ErrorDetail {
  code: string;
  message: string;
}

export interface ErrorResponse {
  error: ErrorDetail;
}

// API Responses
export interface MenusResponse {
  menus: Menu[];
}

export interface OrdersResponse {
  orders: Order[];
}

export interface TableHistoryResponse {
  history: OrderHistory[];
}

export interface TableResponse {
  table: Table;
  sessionId: string;
}
