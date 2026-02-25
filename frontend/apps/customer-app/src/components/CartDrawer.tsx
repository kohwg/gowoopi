'use client';

import { useRouter } from 'next/navigation';
import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button } from '@heroui/react';
import { useTranslation } from '@gowoopi/shared';
import { useCartStore } from '@/stores/cart';
import { CartItemRow } from './CartItem';

interface CartDrawerProps {
  isOpen: boolean;
  onClose: () => void;
}

export function CartDrawer({ isOpen, onClose }: CartDrawerProps) {
  const router = useRouter();
  const { t } = useTranslation();
  const { items, total, clear } = useCartStore();

  const handleCheckout = () => {
    onClose();
    router.push('/checkout');
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="2xl" scrollBehavior="inside">
      <ModalContent>
        <ModalHeader className="text-2xl font-bold border-b pb-4">
          {t('cart.title')}
        </ModalHeader>
        <ModalBody className="py-4">
          {items.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-default-400 text-lg mb-2">🛒</p>
              <p className="text-default-500">{t('cart.empty')}</p>
            </div>
          ) : (
            <div className="space-y-2">
              {items.map((item) => <CartItemRow key={item.menu.id} item={item} />)}
            </div>
          )}
        </ModalBody>
        <ModalFooter className="flex-col gap-3 border-t pt-4">
          <div className="w-full flex justify-between items-center text-xl font-bold px-2">
            <span>{t('cart.total')}</span>
            <span className="text-primary">₩{total().toLocaleString()}</span>
          </div>
          <div className="w-full flex gap-3">
            <Button 
              variant="flat" 
              onPress={clear} 
              className="flex-1 font-semibold"
              size="lg"
            >
              {t('cart.clear')}
            </Button>
            <Button 
              color="primary" 
              onPress={handleCheckout} 
              isDisabled={items.length === 0} 
              className="flex-1 font-semibold shadow-md"
              size="lg"
            >
              {t('cart.checkout')}
            </Button>
          </div>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
