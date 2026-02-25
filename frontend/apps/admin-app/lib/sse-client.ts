import type { Order } from '@gowoopi/shared';

export type SSEEventType = 'order_created' | 'order_updated' | 'order_deleted';

export interface SSEEvent {
  type: SSEEventType;
  data: Order;
}

export function createSSEClient(
  url: string,
  onEvent: (event: SSEEvent) => void,
  onConnect: () => void,
  onDisconnect: () => void
): { close: () => void } {
  const eventSource = new EventSource(url, { withCredentials: true });

  eventSource.onopen = () => onConnect();
  eventSource.onerror = () => onDisconnect();

  eventSource.onmessage = (e) => {
    try {
      const event: SSEEvent = JSON.parse(e.data);
      onEvent(event);
    } catch {
      // ignore parse errors
    }
  };

  return {
    close: () => eventSource.close(),
  };
}
