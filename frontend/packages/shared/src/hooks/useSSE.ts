import { useEffect, useRef, useCallback, useState } from 'react';
import type { Order } from '../types';

interface SSEEvent {
  type: 'new_order' | 'order_updated' | 'order_deleted';
  data: Order;
}

interface UseSSEOptions {
  onNewOrder?: (order: Order) => void;
  onOrderUpdated?: (order: Order) => void;
  onOrderDeleted?: (order: Order) => void;
}

export function useSSE(baseUrl: string, options: UseSSEOptions = {}) {
  const [isConnected, setIsConnected] = useState(false);
  const eventSourceRef = useRef<EventSource | null>(null);

  const connect = useCallback(() => {
    if (eventSourceRef.current) return;

    const es = new EventSource(`${baseUrl}/api/admin/orders/stream`, { withCredentials: true });
    eventSourceRef.current = es;

    es.onopen = () => setIsConnected(true);
    es.onerror = () => {
      setIsConnected(false);
      es.close();
      eventSourceRef.current = null;
      setTimeout(connect, 3000);
    };

    es.onmessage = (event) => {
      try {
        const data: SSEEvent = JSON.parse(event.data);
        if (data.type === 'new_order') options.onNewOrder?.(data.data);
        else if (data.type === 'order_updated') options.onOrderUpdated?.(data.data);
        else if (data.type === 'order_deleted') options.onOrderDeleted?.(data.data);
      } catch { /* ignore parse errors */ }
    };
  }, [baseUrl, options]);

  const disconnect = useCallback(() => {
    eventSourceRef.current?.close();
    eventSourceRef.current = null;
    setIsConnected(false);
  }, []);

  useEffect(() => {
    return () => disconnect();
  }, [disconnect]);

  return { isConnected, connect, disconnect };
}
