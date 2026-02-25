import { create } from 'zustand';
import type { SSEEvent } from '@/lib/sse-client';

interface SSEState {
  isConnected: boolean;
  newOrderIds: Set<string>;
  setConnected: (connected: boolean) => void;
  addNewOrder: (orderId: string) => void;
  clearNewOrder: (orderId: string) => void;
  handleEvent: (event: SSEEvent) => void;
}

export const useSSEStore = create<SSEState>((set) => ({
  isConnected: false,
  newOrderIds: new Set(),

  setConnected: (connected) => set({ isConnected: connected }),

  addNewOrder: (orderId) =>
    set((state) => ({
      newOrderIds: new Set(state.newOrderIds).add(orderId),
    })),

  clearNewOrder: (orderId) =>
    set((state) => {
      const next = new Set(state.newOrderIds);
      next.delete(orderId);
      return { newOrderIds: next };
    }),

  handleEvent: (event) => {
    if (event.type === 'order_created') {
      set((state) => ({
        newOrderIds: new Set(state.newOrderIds).add(event.data.id),
      }));
    }
  },
}));
