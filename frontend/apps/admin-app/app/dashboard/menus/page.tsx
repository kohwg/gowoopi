'use client';

import { useState, useMemo, useEffect } from 'react';
import { Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Spinner, Image } from '@heroui/react';
import { useMenus, useCategories, useCreateMenu, useUpdateMenu, useDeleteMenu, useUpdateMenuOrder, useAuth, useTranslation, type Menu, type MenuCreateRequest, type MenuUpdateRequest } from '@gowoopi/shared';
import { CategoryTabs, MenuFormModal } from '@/components/menus';
import { ConfirmModal } from '@/components/shared/ConfirmModal';

export default function MenusPage() {
  const { t } = useTranslation();
  const { auth } = useAuth();
  const { data: menus, isLoading } = useMenus(auth?.storeId || '', 'admin');
  const { data: categories } = useCategories(auth?.storeId || '');
  const createMenu = useCreateMenu();
  const updateMenu = useUpdateMenu();
  const deleteMenu = useDeleteMenu();
  const updateOrder = useUpdateMenuOrder();

  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [editMenu, setEditMenu] = useState<Menu | undefined>();
  const [showForm, setShowForm] = useState(false);
  const [deleteTarget, setDeleteTarget] = useState<Menu | null>(null);

  useEffect(() => {
    if (categories?.length && !selectedCategory) {
      setSelectedCategory(categories[0].id);
    }
  }, [categories, selectedCategory]);

  const filteredMenus = useMemo(() => {
    if (!selectedCategory) return [];
    return (menus || [])
      .filter((m) => m.categoryId === selectedCategory)
      .sort((a, b) => a.displayOrder - b.displayOrder);
  }, [menus, selectedCategory]);

  const handleSubmit = async (data: MenuCreateRequest | MenuUpdateRequest) => {
    if (editMenu) {
      await updateMenu.mutateAsync({ id: editMenu.id, data });
    } else {
      await createMenu.mutateAsync(data as MenuCreateRequest);
    }
    setEditMenu(undefined);
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    await deleteMenu.mutateAsync(deleteTarget.id);
    setDeleteTarget(null);
  };

  const handleMoveUp = async (menu: Menu, index: number) => {
    if (index === 0) return;
    const items = filteredMenus.map((m, i) => ({ id: m.id, displayOrder: i }));
    [items[index - 1], items[index]] = [items[index], items[index - 1]];
    await updateOrder.mutateAsync(items.map((item, i) => ({ id: item.id, displayOrder: i })));
  };

  const handleMoveDown = async (menu: Menu, index: number) => {
    if (index === filteredMenus.length - 1) return;
    const items = filteredMenus.map((m, i) => ({ id: m.id, displayOrder: i }));
    [items[index], items[index + 1]] = [items[index + 1], items[index]];
    await updateOrder.mutateAsync(items.map((item, i) => ({ id: item.id, displayOrder: i })));
  };

  if (isLoading) {
    return <div className="flex justify-center p-8"><Spinner size="lg" /></div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">{t('menu.management')}</h1>
        <Button color="primary" onPress={() => { setEditMenu(undefined); setShowForm(true); }}>
          {t('menu.add')}
        </Button>
      </div>

      <CategoryTabs categories={categories || []} selectedId={selectedCategory} onSelect={setSelectedCategory} />

      <Table aria-label="Menus">
        <TableHeader>
          <TableColumn>{t('menu.image')}</TableColumn>
          <TableColumn>{t('menu.name')}</TableColumn>
          <TableColumn>{t('menu.price')}</TableColumn>
          <TableColumn>{t('menu.order')}</TableColumn>
          <TableColumn>{t('common.actions')}</TableColumn>
        </TableHeader>
        <TableBody emptyContent={t('menu.empty')}>
          {filteredMenus.map((menu, index) => (
            <TableRow key={menu.id}>
              <TableCell>
                {menu.imageUrl && <Image src={menu.imageUrl} alt={menu.name} width={50} height={50} className="rounded" />}
              </TableCell>
              <TableCell>{menu.name}</TableCell>
              <TableCell>₩{menu.price.toLocaleString()}</TableCell>
              <TableCell>
                <div className="flex gap-1">
                  <Button size="sm" variant="light" isIconOnly onPress={() => handleMoveUp(menu, index)} isDisabled={index === 0}>↑</Button>
                  <Button size="sm" variant="light" isIconOnly onPress={() => handleMoveDown(menu, index)} isDisabled={index === filteredMenus.length - 1}>↓</Button>
                </div>
              </TableCell>
              <TableCell>
                <div className="flex gap-2">
                  <Button size="sm" variant="light" onPress={() => { setEditMenu(menu); setShowForm(true); }}>
                    {t('common.edit')}
                  </Button>
                  <Button size="sm" color="danger" variant="light" onPress={() => setDeleteTarget(menu)}>
                    {t('common.delete')}
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <MenuFormModal
        menu={editMenu}
        categories={categories || []}
        isOpen={showForm}
        onClose={() => { setShowForm(false); setEditMenu(undefined); }}
        onSubmit={handleSubmit}
      />
      <ConfirmModal
        title={t('menu.deleteConfirm.title')}
        message={t('menu.deleteConfirm.message')}
        isOpen={!!deleteTarget}
        onConfirm={handleDelete}
        onCancel={() => setDeleteTarget(null)}
      />
    </div>
  );
}
