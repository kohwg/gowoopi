'use client';

import { Card, CardBody, Chip } from '@heroui/react';
import { useTranslation, type Order, type OrderStatus } from '@gowoopi/shared';

const statusColorMap: Record<OrderStatus, 'warning' | 'primary' | 'success'> = {
  pending: 'warning',
  preparing: 'primary',
  completed: 'success',
};

interface OrderCardProps {
  order: Order;
}

export function OrderCard({ order }: OrderCardProps) {
  const { t } = useTranslation();

  return (
    <Card className="w-full shadow-lg hover:shadow-xl transition-shadow">
      <CardBody className="p-6">
        <div className="flex justify-between items-start mb-4">
          <div>
            <p className="text-sm text-default-400 font-mono">#{order.id.slice(0, 8)}</p>
            <p className="text-xs text-default-400 mt-1">
              {new Date(order.createdAt).toLocaleString()}
            </p>
          </div>
          <Chip color={statusColorMap[order.status]} size="lg" variant="flat" className="font-semibold">
            {t(`order.status.${order.status}`)}
          </Chip>
        </div>
        <div className="space-y-2 mb-4">
          {order.items.map((item) => (
            <div key={item.id} className="flex justify-between text-sm py-2 border-b last:border-0">
              <span className="font-medium">{item.menuName} × {item.quantity}</span>
              <span className="font-semibold">₩{(item.unitPrice * item.quantity).toLocaleString()}</span>
            </div>
          ))}
        </div>
        <div className="flex justify-between items-center pt-3 border-t-2 border-primary/20">
          <span className="text-lg font-bold">{t('order.totalAmount')}</span>
          <span className="text-xl font-bold text-primary">₩{order.totalAmount.toLocaleString()}</span>
        </div>
      </CardBody>
    </Card>
  );
}
