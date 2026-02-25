'use client';

import { Button } from '@heroui/react';
import { useTranslation } from '@gowoopi/shared';
import type { CartItem } from '@/stores/cart';
import { useCartStore } from '@/stores/cart';

interface CartItemRowProps {
  item: CartItem;
}

export function CartItemRow({ item }: CartItemRowProps) {
  const { updateQuantity, removeItem } = useCartStore();

  return (
    <div className="flex items-center justify-between py-2 border-b">
      <div className="flex-1">
        <p className="font-medium">{item.menu.name}</p>
        <p className="text-sm text-default-500">₩{item.menu.price.toLocaleString()}</p>
      </div>
      <div className="flex items-center gap-2">
        <Button size="sm" variant="flat" onPress={() => updateQuantity(item.menu.id, item.quantity - 1)}>
          -
        </Button>
        <span className="w-8 text-center">{item.quantity}</span>
        <Button size="sm" variant="flat" onPress={() => updateQuantity(item.menu.id, item.quantity + 1)}>
          +
        </Button>
        <Button size="sm" color="danger" variant="light" onPress={() => removeItem(item.menu.id)}>
          ✕
        </Button>
      </div>
    </div>
  );
}
