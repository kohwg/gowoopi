import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { Menu } from '@gowoopi/shared';

export interface CartItem {
  menu: Menu;
  quantity: number;
}

interface CartState {
  items: CartItem[];
  addItem: (menu: Menu) => void;
  removeItem: (menuId: number) => void;
  updateQuantity: (menuId: number, quantity: number) => void;
  clear: () => void;
  total: () => number;
}

export const useCartStore = create<CartState>()(
  persist(
    (set, get) => ({
      items: [],
      addItem: (menu) =>
        set((state) => {
          const existing = state.items.find((i) => i.menu.id === menu.id);
          if (existing) {
            return {
              items: state.items.map((i) =>
                i.menu.id === menu.id ? { ...i, quantity: i.quantity + 1 } : i
              ),
            };
          }
          return { items: [...state.items, { menu, quantity: 1 }] };
        }),
      removeItem: (menuId) =>
        set((state) => ({ items: state.items.filter((i) => i.menu.id !== menuId) })),
      updateQuantity: (menuId, quantity) =>
        set((state) => ({
          items: quantity <= 0
            ? state.items.filter((i) => i.menu.id !== menuId)
            : state.items.map((i) => (i.menu.id === menuId ? { ...i, quantity } : i)),
        })),
      clear: () => set({ items: [] }),
      total: () => get().items.reduce((sum, i) => sum + i.menu.price * i.quantity, 0),
    }),
    { name: 'gowoopi-cart' }
  )
);
