'use client';

import { Button } from '@heroui/react';
import type { CartItem } from '@/stores/cart';
import { useCartStore } from '@/stores/cart';

interface CartItemRowProps {
  item: CartItem;
}

export function CartItemRow({ item }: CartItemRowProps) {
  const { updateQuantity, removeItem } = useCartStore();

  return (
    <div className="flex items-center justify-between py-3 px-2 border-b hover:bg-default-50 rounded-lg transition-colors">
      <div className="flex-1">
        <p className="font-semibold text-base">{item.menu.name}</p>
        <p className="text-sm text-default-500 mt-1">₩{item.menu.price.toLocaleString()}</p>
      </div>
      <div className="flex items-center gap-3">
        <div className="flex items-center gap-2 bg-default-100 rounded-full px-1">
          <Button 
            size="sm" 
            variant="light" 
            isIconOnly
            className="min-w-8 h-8 rounded-full"
            onPress={() => updateQuantity(item.menu.id, item.quantity - 1)}
          >
            −
          </Button>
          <span className="w-8 text-center font-semibold">{item.quantity}</span>
          <Button 
            size="sm" 
            variant="light" 
            isIconOnly
            className="min-w-8 h-8 rounded-full"
            onPress={() => updateQuantity(item.menu.id, item.quantity + 1)}
          >
            +
          </Button>
        </div>
        <Button 
          size="sm" 
          color="danger" 
          variant="light" 
          isIconOnly
          className="min-w-8 h-8"
          onPress={() => removeItem(item.menu.id)}
        >
          ✕
        </Button>
      </div>
    </div>
  );
}
