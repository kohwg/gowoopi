'use client';

import { useRouter } from 'next/navigation';
import { Card, CardBody, CardFooter, Button } from '@heroui/react';
import { useTranslation, useCreateOrder } from '@gowoopi/shared';
import { useCartStore } from '@/stores/cart';

export default function CheckoutPage() {
  const router = useRouter();
  const { t } = useTranslation();
  const { items, total, clear } = useCartStore();
  const createOrder = useCreateOrder();

  const handleOrder = () => {
    createOrder.mutate(
      { items: items.map((i) => ({ menuId: i.menu.id, quantity: i.quantity })) },
      {
        onSuccess: () => {
          clear();
          router.replace('/orders');
        },
      }
    );
  };

  if (items.length === 0) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <Card>
          <CardBody className="text-center py-8">
            <p>{t('cart.empty')}</p>
            <Button className="mt-4" onPress={() => router.push('/')}>
              {t('common.back')}
            </Button>
          </CardBody>
        </Card>
      </div>
    );
  }

  return (
    <div className="min-h-screen p-4 max-w-lg mx-auto">
      <h1 className="text-xl font-bold mb-4">{t('order.title')}</h1>
      <Card>
        <CardBody className="gap-2">
          {items.map((item) => (
            <div key={item.menu.id} className="flex justify-between">
              <span>{item.menu.name} x {item.quantity}</span>
              <span>₩{(item.menu.price * item.quantity).toLocaleString()}</span>
            </div>
          ))}
          <div className="flex justify-between mt-4 pt-2 border-t font-bold text-lg">
            <span>{t('cart.total')}</span>
            <span>₩{total().toLocaleString()}</span>
          </div>
        </CardBody>
        <CardFooter className="gap-2">
          <Button variant="flat" onPress={() => router.back()} className="flex-1">
            {t('common.back')}
          </Button>
          <Button color="primary" onPress={handleOrder} isLoading={createOrder.isPending} className="flex-1">
            {t('common.confirm')}
          </Button>
        </CardFooter>
      </Card>
      {createOrder.isError && (
        <p className="text-danger text-center mt-4">{t('order.orderFailed')}</p>
      )}
    </div>
  );
}
