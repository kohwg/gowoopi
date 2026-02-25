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
    <div className="fixed bottom-8 right-8 z-50">
      <Badge 
        content={count} 
        color="danger" 
        isInvisible={count === 0}
        size="lg"
        className="font-bold"
      >
        <Button 
          isIconOnly 
          size="lg" 
          color="primary" 
          className="rounded-full shadow-2xl w-16 h-16 text-2xl hover:scale-110 transition-transform" 
          onPress={onPress}
        >
          🛒
        </Button>
      </Badge>
    </div>
  );
}
