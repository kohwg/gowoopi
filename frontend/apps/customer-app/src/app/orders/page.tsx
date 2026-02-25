'use client';

import { useRouter } from 'next/navigation';
import { Button, Spinner } from '@heroui/react';
import { useTranslation, useCustomerOrders } from '@gowoopi/shared';
import { OrderCard } from '@/components/OrderCard';

export default function OrdersPage() {
  const router = useRouter();
  const { t } = useTranslation();
  const { data: orders, isLoading } = useCustomerOrders();

  return (
    <div className="min-h-screen p-4 max-w-lg mx-auto">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-xl font-bold">{t('order.orders')}</h1>
        <Button variant="light" onPress={() => router.push('/')}>
          {t('menu.title')}
        </Button>
      </div>
      {isLoading ? (
        <div className="flex justify-center py-8">
          <Spinner />
        </div>
      ) : !orders?.length ? (
        <p className="text-center text-default-500 py-8">{t('order.noOrders')}</p>
      ) : (
        <div className="space-y-4">
          {orders.map((order) => (
            <OrderCard key={order.id} order={order} />
          ))}
        </div>
      )}
    </div>
  );
}
