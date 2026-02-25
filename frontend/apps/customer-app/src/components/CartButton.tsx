'use client';

import { Button, Badge } from '@heroui/react';
import { useCartStore } from '@/stores/cart';

interface CartButtonProps {
  onPress: () => void;
}

export function CartButton({ onPress }: CartButtonProps) {
  const items = useCartStore((s) => s.items);
  const count = items.reduce((sum, i) => sum + i.quantity, 0);

  return (
    <div className="fixed bottom-6 right-6 z-50">
      <Badge content={count} color="danger" isInvisible={count === 0}>
        <Button isIconOnly size="lg" color="primary" className="rounded-full shadow-lg" onPress={onPress}>
          🛒
        </Button>
      </Badge>
    </div>
  );
}
