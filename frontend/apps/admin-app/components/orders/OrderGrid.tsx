'use client';

import { useMemo } from 'react';
import type { Order } from '@gowoopi/shared';
import { TableCard } from './TableCard';
import { useSSEStore } from '@/stores/sse-store';

interface OrderGridProps {
  orders: Order[];
  onOrderClick: (tableNumber: number) => void;
}

export function OrderGrid({ orders, onOrderClick }: OrderGridProps) {
  const { newOrderIds } = useSSEStore();

  const tableGroups = useMemo(() => {
    const groups = new Map<number, { orders: Order[]; total: number }>();
    for (const order of orders) {
      const existing = groups.get(order.tableId) || { orders: [], total: 0 };
      existing.orders.push(order);
      existing.total += order.totalAmount;
      groups.set(order.tableId, existing);
    }
    return Array.from(groups.entries()).sort((a, b) => a[0] - b[0]);
  }, [orders]);

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      {tableGroups.map(([tableId, { orders: tableOrders, total }]) => {
        const hasNew = tableOrders.some((o) => newOrderIds.has(o.id));
        return (
          <TableCard
            key={tableId}
            tableNumber={tableId}
            orders={tableOrders}
            totalAmount={total}
            isNew={hasNew}
            onClick={() => onOrderClick(tableId)}
          />
        );
      })}
    </div>
  );
}
