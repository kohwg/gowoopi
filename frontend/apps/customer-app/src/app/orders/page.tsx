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
    <div className="min-h-screen p-6 bg-gradient-to-b from-background to-default-50">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            {t('order.orders')}
          </h1>
          <Button 
            variant="flat" 
            onPress={() => router.push('/')}
            className="font-semibold"
            size="lg"
          >
            {t('menu.title')}
          </Button>
        </div>
        {isLoading ? (
          <div className="flex justify-center py-12">
            <Spinner size="lg" />
          </div>
        ) : !orders?.length ? (
          <Card className="shadow-xl">
            <CardBody className="text-center py-12">
              <p className="text-6xl mb-4">📋</p>
              <p className="text-default-400 text-lg">{t('order.noOrders')}</p>
            </CardBody>
          </Card>
        ) : (
          <div className="space-y-4">
            {orders.map((order) => (
              <OrderCard key={order.id} order={order} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
