'use client';

import { Button, Badge } from '@heroui/react';
import { useCartStore } from '@/stores/cart';

interface CartButtonProps {
  onPress: () => void;
}

export function CartButton({ onPress }: CartButtonProps) {
  const items = useCartStore((s) => s.items);
  const count = items.reduce((sum, i) => sum + i.quantity, 0);

  if (count === 0) return null;

  return (
    <div className="fixed bottom-8 right-8 z-50 animate-in fade-in slide-in-from-bottom-4">
      <Badge content={count} color="danger" size="lg" className="text-base font-bold">
        <Button 
          isIconOnly
          size="lg" 
          className="rounded-full bg-white shadow-[0_10px_40px_rgba(0,0,0,0.5),0_0_20px_rgba(0,0,0,0.3)] h-16 w-16 text-3xl hover:scale-110 transition-transform"
          onPress={onPress}
        >
          🛒
        </Button>
      </Badge>
    </div>
  );
}
