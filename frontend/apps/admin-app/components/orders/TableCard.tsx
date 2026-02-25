'use client';

import { Card, CardBody, CardHeader } from '@heroui/react';
import type { Order } from '@gowoopi/shared';
import { useTranslation } from '@gowoopi/shared';
import { StatusBadge } from './StatusBadge';

interface TableCardProps {
  tableNumber: number;
  orders: Order[];
  totalAmount: number;
  isNew?: boolean;
  onClick: () => void;
}

export function TableCard({ tableNumber, orders, totalAmount, isNew, onClick }: TableCardProps) {
  const { t } = useTranslation();
  const latestOrder = orders[0];

  return (
    <Card
      isPressable
      onPress={onClick}
      className={`transition-all ${isNew ? 'ring-2 ring-primary animate-pulse' : ''}`}
    >
      <CardHeader className="flex justify-between">
        <span className="font-bold">{t('table.number.label')} {tableNumber}</span>
        <span className="text-lg font-semibold">₩{totalAmount.toLocaleString()}</span>
      </CardHeader>
      <CardBody>
        {orders.length === 0 ? (
          <p className="text-gray-400">{t('order.empty')}</p>
        ) : (
          <div className="space-y-2">
            <div className="flex justify-between items-center">
              <span className="text-sm text-gray-500">
                {t('order.latest')}: {latestOrder.items.length} {t('order.items')}
              </span>
              <StatusBadge status={latestOrder.status} />
            </div>
            {orders.length > 1 && (
              <p className="text-xs text-gray-400">+{orders.length - 1} {t('order.more')}</p>
            )}
          </div>
        )}
      </CardBody>
    </Card>
  );
}
