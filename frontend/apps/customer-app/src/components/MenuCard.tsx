'use client';

import { Card, CardBody, CardFooter, Image, Button } from '@heroui/react';
import { useTranslation, type Menu } from '@gowoopi/shared';
import { useCartStore } from '@/stores/cart';

interface MenuCardProps {
  menu: Menu;
}

export function MenuCard({ menu }: MenuCardProps) {
  const { t } = useTranslation();
  const addItem = useCartStore((s) => s.addItem);

  return (
    <Card className="w-full hover:scale-105 transition-transform duration-200 shadow-md hover:shadow-xl">
      {menu.imageUrl && (
        <Image
          alt={menu.name}
          src={menu.imageUrl}
          className="w-full h-48 object-cover"
        />
      )}
      <CardBody className="p-4 gap-2">
        <h3 className="font-bold text-lg">{menu.name}</h3>
        {menu.description && (
          <p className="text-sm text-default-500 line-clamp-2 min-h-[2.5rem]">{menu.description}</p>
        )}
        <p className="text-xl font-bold text-primary mt-1">₩{menu.price.toLocaleString()}</p>
      </CardBody>
      <CardFooter className="pt-0 px-4 pb-4">
        <Button
          color="primary"
          className="w-full font-semibold shadow-md"
          size="lg"
          onPress={() => addItem(menu)}
        >
          {t('menu.addToCart')}
        </Button>
      </CardFooter>
    </Card>
  );
}
