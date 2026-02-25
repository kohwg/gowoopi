'use client';

import { Chip } from '@heroui/react';
import type { OrderStatus } from '@gowoopi/shared';
import { useTranslation } from '@gowoopi/shared';

const statusConfig: Record<OrderStatus, { color: 'warning' | 'primary' | 'success' | 'secondary'; labelKey: string }> = {
  PENDING: { color: 'warning', labelKey: 'order.status.pending' },
  CONFIRMED: { color: 'secondary', labelKey: 'order.status.confirmed' },
  PREPARING: { color: 'primary', labelKey: 'order.status.preparing' },
  COMPLETED: { color: 'success', labelKey: 'order.status.completed' },
};

interface StatusBadgeProps {
  status: OrderStatus;
}

export function StatusBadge({ status }: StatusBadgeProps) {
  const { t } = useTranslation();
  const config = statusConfig[status];
  return <Chip color={config.color} size="sm">{t(config.labelKey)}</Chip>;
}
