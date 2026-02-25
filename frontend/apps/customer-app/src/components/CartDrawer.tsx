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
    <Modal isOpen={isOpen} onClose={onClose} size="lg">
      <ModalContent>
        <ModalHeader>{t('cart.title')}</ModalHeader>
        <ModalBody>
          {items.length === 0 ? (
            <p className="text-center text-default-500 py-8">{t('cart.empty')}</p>
          ) : (
            items.map((item) => <CartItemRow key={item.menu.id} item={item} />)
          )}
        </ModalBody>
        <ModalFooter className="flex-col gap-2">
          <div className="w-full flex justify-between text-lg font-bold">
            <span>{t('cart.total')}</span>
            <span>₩{total().toLocaleString()}</span>
          </div>
          <div className="w-full flex gap-2">
            <Button variant="flat" onPress={clear} className="flex-1">
              {t('cart.clear')}
            </Button>
            <Button color="primary" onPress={handleCheckout} isDisabled={items.length === 0} className="flex-1">
              {t('cart.checkout')}
            </Button>
          </div>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
