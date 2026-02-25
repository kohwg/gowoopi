'use client';

import { useState, useMemo } from 'react';
import { useRouter } from 'next/navigation';
import { Button, Spinner } from '@heroui/react';
import { useTranslation, useMenus, useAuth, type Category } from '@gowoopi/shared';
import { useAutoLogin } from '@/hooks/useAutoLogin';
import { MenuCard } from '@/components/MenuCard';
import { CategoryTabs } from '@/components/CategoryTabs';
import { CartButton } from '@/components/CartButton';
import { CartDrawer } from '@/components/CartDrawer';
import { LanguageSwitcher } from '@/components/LanguageSwitcher';

export default function HomePage() {
  const router = useRouter();
  const { t } = useTranslation();
  const { auth } = useAuth();
  const { isLoading: authLoading, isAuthenticated } = useAutoLogin();
  const { data: menus, isLoading: menusLoading } = useMenus(auth?.storeId ?? '');

  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [cartOpen, setCartOpen] = useState(false);

  const categories = useMemo(() => {
    if (!menus) return [];
    const catMap = new Map<number, Category>();
    menus.forEach((m) => {
      if (!catMap.has(m.categoryId)) {
        catMap.set(m.categoryId, { id: m.categoryId, storeId: m.storeId, name: m.category?.name || `Category ${m.categoryId}`, sortOrder: m.categoryId });
      }
    });
    return Array.from(catMap.values());
  }, [menus]);

  const filteredMenus = useMemo(() => {
    if (!menus) return [];
    if (!selectedCategory) return menus;
    return menus.filter((m) => m.categoryId === selectedCategory);
  }, [menus, selectedCategory]);

  if (authLoading || !isAuthenticated) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <Spinner size="lg" />
      </div>
    );
  }

  return (
    <div className="min-h-screen pb-20">
      <header className="sticky top-0 bg-background z-40 p-4 border-b flex justify-between items-center">
        <h1 className="text-xl font-bold">{auth?.storeId}</h1>
        <div className="flex gap-2">
          <Button variant="light" size="sm" onPress={() => router.push('/orders')}>
            {t('order.orders')}
          </Button>
          <LanguageSwitcher />
        </div>
      </header>

      <main className="p-4">
        <CategoryTabs categories={categories} selected={selectedCategory} onSelect={setSelectedCategory} />

        {menusLoading ? (
          <div className="flex justify-center py-8">
            <Spinner />
          </div>
        ) : !filteredMenus.length ? (
          <p className="text-center text-default-500 py-8">{t('menu.noMenus')}</p>
        ) : (
          <div className="grid grid-cols-2 gap-4">
            {filteredMenus.map((menu) => (
              <MenuCard key={menu.id} menu={menu} />
            ))}
          </div>
        )}
      </main>

      <CartButton onPress={() => setCartOpen(true)} />
      <CartDrawer isOpen={cartOpen} onClose={() => setCartOpen(false)} />
    </div>
  );
}
