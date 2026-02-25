'use client';

import { useState, useMemo } from 'react';
import { Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Spinner, Image } from '@heroui/react';
import { useMenus, useCreateMenu, useUpdateMenu, useDeleteMenu, useUpdateMenuOrder, useAuth, useTranslation, type Menu, type Category, type MenuCreateRequest, type MenuUpdateRequest } from '@gowoopi/shared';
import { CategoryTabs, MenuFormModal } from '@/components/menus';
import { ConfirmModal } from '@/components/shared/ConfirmModal';

// TODO: API에서 카테고리 조회 추가 필요. 임시로 하드코딩
const mockCategories: Category[] = [
  { id: 1, storeId: '', name: '메인', displayOrder: 1 },
  { id: 2, storeId: '', name: '사이드', displayOrder: 2 },
  { id: 3, storeId: '', name: '음료', displayOrder: 3 },
];

export default function MenusPage() {
  const { t } = useTranslation();
  const { auth } = useAuth();
  const { data: menus, isLoading } = useMenus(auth?.storeId || '', 'admin');
  const createMenu = useCreateMenu();
  const updateMenu = useUpdateMenu();
  const deleteMenu = useDeleteMenu();
  const updateOrder = useUpdateMenuOrder();

  const [selectedCategory, setSelectedCategory] = useState(mockCategories[0]?.id || 1);
  const [editMenu, setEditMenu] = useState<Menu | undefined>();
  const [showForm, setShowForm] = useState(false);
  const [deleteTarget, setDeleteTarget] = useState<Menu | null>(null);

  const filteredMenus = useMemo(() => {
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

      <CategoryTabs categories={mockCategories} selectedId={selectedCategory} onSelect={setSelectedCategory} />

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
        categories={mockCategories}
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
