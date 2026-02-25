import { describe, it, expect, beforeEach } from 'vitest';
import { useSSEStore } from '@/stores/sse-store';
import type { SSEEvent } from '@/lib/sse-client';

describe('SSE Store', () => {
  beforeEach(() => {
    useSSEStore.setState({ isConnected: false, newOrderIds: new Set() });
  });

  it('SSE-01: should set connected state', () => {
    const { setConnected } = useSSEStore.getState();
    setConnected(true);
    expect(useSSEStore.getState().isConnected).toBe(true);
  });

  it('SSE-02: should add new order on order_created event', () => {
    const { handleEvent } = useSSEStore.getState();
    const event: SSEEvent = {
      type: 'order_created',
      data: { id: 'order-1' } as SSEEvent['data'],
    };
    handleEvent(event);
    expect(useSSEStore.getState().newOrderIds.has('order-1')).toBe(true);
  });

  it('SSE-03: should clear new order', () => {
    const { addNewOrder, clearNewOrder } = useSSEStore.getState();
    addNewOrder('order-1');
    expect(useSSEStore.getState().newOrderIds.has('order-1')).toBe(true);
    clearNewOrder('order-1');
    expect(useSSEStore.getState().newOrderIds.has('order-1')).toBe(false);
  });
});
