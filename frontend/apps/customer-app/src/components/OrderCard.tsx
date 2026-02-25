'use client';

import { Card, CardBody, Chip } from '@heroui/react';
import { useTranslation, type Order, type OrderStatus } from '@gowoopi/shared';

const statusColorMap: Record<OrderStatus, 'warning' | 'primary' | 'success' | 'secondary'> = {
  PENDING: 'warning',
  CONFIRMED: 'secondary',
  PREPARING: 'primary',
  COMPLETED: 'success',
};

interface OrderCardProps {
  order: Order;
}

export function OrderCard({ order }: OrderCardProps) {
  const { t } = useTranslation();

  return (
    <Card className="w-full">
      <CardBody>
        <div className="flex justify-between items-start mb-2">
          <div>
            <p className="text-sm text-default-500">#{order.id.slice(0, 8)}</p>
            <p className="text-xs text-default-400">
              {new Date(order.createdAt).toLocaleString()}
            </p>
          </div>
          <Chip color={statusColorMap[order.status]} size="sm">
            {t(`order.${order.status}`)}
          </Chip>
        </div>
        <div className="space-y-1">
          {order.items.map((item) => (
            <div key={item.id} className="flex justify-between text-sm">
              <span>{item.menuName} x {item.quantity}</span>
              <span>₩{item.subtotal.toLocaleString()}</span>
            </div>
          ))}
        </div>
        <div className="flex justify-between mt-3 pt-2 border-t font-bold">
          <span>{t('order.totalAmount')}</span>
          <span>₩{order.totalAmount.toLocaleString()}</span>
        </div>
      </CardBody>
    </Card>
  );
}
