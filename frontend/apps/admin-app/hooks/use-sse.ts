import { useEffect, useRef } from 'react';
import { useQueryClient } from '@tanstack/react-query';
import { orderKeys } from '@gowoopi/shared';
import { createSSEClient } from '@/lib/sse-client';
import { useSSEStore } from '@/stores/sse-store';

const RECONNECT_DELAY = 3000;

export function useSSE(baseUrl: string): void {
  const queryClient = useQueryClient();
  const { setConnected, handleEvent } = useSSEStore();
  const reconnectTimer = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    let client: { close: () => void } | null = null;

    const connect = () => {
      client = createSSEClient(
        `${baseUrl}/api/admin/orders/stream`,
        (event) => {
          handleEvent(event);
          queryClient.invalidateQueries({ queryKey: orderKeys.admin() });
        },
        () => setConnected(true),
        () => {
          setConnected(false);
          reconnectTimer.current = setTimeout(connect, RECONNECT_DELAY);
        }
      );
    };

    connect();

    return () => {
      client?.close();
      if (reconnectTimer.current) clearTimeout(reconnectTimer.current);
    };
  }, [baseUrl, queryClient, setConnected, handleEvent]);
}
