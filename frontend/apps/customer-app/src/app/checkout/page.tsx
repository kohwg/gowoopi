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
      <div className="min-h-screen flex items-center justify-center p-4 bg-gradient-to-b from-background to-default-50">
        <Card className="shadow-xl">
          <CardBody className="text-center py-12 px-8">
            <p className="text-6xl mb-4">🛒</p>
            <p className="text-lg text-default-500 mb-6">{t('cart.empty')}</p>
            <Button 
              color="primary" 
              size="lg"
              className="font-semibold shadow-md"
              onPress={() => router.push('/')}
            >
              {t('common.back')}
            </Button>
          </CardBody>
        </Card>
      </div>
    );
  }

  return (
    <div className="min-h-screen p-6 bg-gradient-to-b from-background to-default-50">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold mb-6 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          {t('order.title')}
        </h1>
        <Card className="shadow-xl">
          <CardBody className="gap-3 p-6">
            {items.map((item) => (
              <div key={item.menu.id} className="flex justify-between items-center py-3 border-b last:border-0">
                <div>
                  <p className="font-semibold text-lg">{item.menu.name}</p>
                  <p className="text-sm text-default-500">수량: {item.quantity}</p>
                </div>
                <span className="font-bold text-lg">₩{(item.menu.price * item.quantity).toLocaleString()}</span>
              </div>
            ))}
            <div className="flex justify-between items-center mt-4 pt-4 border-t-2 border-primary/20">
              <span className="text-xl font-bold">{t('cart.total')}</span>
              <span className="text-2xl font-bold text-primary">₩{total().toLocaleString()}</span>
            </div>
          </CardBody>
          <CardFooter className="gap-3 p-6 pt-0">
            <Button 
              variant="flat" 
              onPress={() => router.back()} 
              className="flex-1 font-semibold"
              size="lg"
            >
              {t('common.back')}
            </Button>
            <Button 
              color="primary" 
              onPress={handleOrder} 
              isLoading={createOrder.isPending} 
              className="flex-1 font-semibold shadow-md"
              size="lg"
            >
              {t('common.confirm')}
            </Button>
          </CardFooter>
        </Card>
        {createOrder.isError && (
          <p className="text-danger text-center mt-6 text-lg font-medium">{t('order.orderFailed')}</p>
        )}
      </div>
    </div>
  );
}
