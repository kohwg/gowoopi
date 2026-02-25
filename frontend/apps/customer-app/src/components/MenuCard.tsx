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
    <Card className="w-full">
      {menu.imageUrl && (
        <Image
          alt={menu.name}
          src={menu.imageUrl}
          className="w-full h-40 object-cover"
        />
      )}
      <CardBody className="p-3">
        <h3 className="font-semibold">{menu.name}</h3>
        {menu.description && (
          <p className="text-sm text-default-500 line-clamp-2">{menu.description}</p>
        )}
        <p className="text-lg font-bold mt-2">₩{menu.price.toLocaleString()}</p>
      </CardBody>
      <CardFooter className="pt-0">
        <Button
          color="primary"
          className="w-full"
          onPress={() => addItem(menu)}
        >
          {t('menu.addToCart')}
        </Button>
      </CardFooter>
    </Card>
  );
}
