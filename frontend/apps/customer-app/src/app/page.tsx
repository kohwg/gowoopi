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
        catMap.set(m.categoryId, { id: m.categoryId, storeId: m.storeId, name: `Category ${m.categoryId}`, displayOrder: m.categoryId });
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
    <div className="min-h-screen pb-20 bg-gradient-to-b from-background to-default-50">
      <header className="sticky top-0 bg-background/80 backdrop-blur-lg z-40 px-6 py-4 border-b border-divider shadow-sm">
        <div className="max-w-7xl mx-auto flex justify-between items-center">
          <h1 className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            {auth?.storeName}
          </h1>
          <div className="flex gap-3">
            <Button 
              variant="flat" 
              size="sm" 
              onPress={() => router.push('/orders')}
              className="font-medium"
            >
              {t('order.orders')}
            </Button>
            <LanguageSwitcher />
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-6 py-6">
        <CategoryTabs categories={categories} selected={selectedCategory} onSelect={setSelectedCategory} />

        {menusLoading ? (
          <div className="flex justify-center py-12">
            <Spinner size="lg" />
          </div>
        ) : !filteredMenus.length ? (
          <p className="text-center text-default-400 py-12 text-lg">{t('menu.noMenus')}</p>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
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
