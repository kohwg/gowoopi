'use client';

import { Tabs, Tab } from '@heroui/react';
import type { Category } from '@gowoopi/shared';

interface CategoryTabsProps {
  categories: Category[];
  selectedId: number | null;
  onSelect: (id: number) => void;
}

export function CategoryTabs({ categories, selectedId, onSelect }: CategoryTabsProps) {
  if (!categories.length || selectedId === null) return null;
  
  return (
    <Tabs
      selectedKey={String(selectedId)}
      onSelectionChange={(key) => onSelect(Number(key))}
      className="mb-4"
    >
      {categories.map((cat) => (
        <Tab key={String(cat.id)} title={cat.name} />
      ))}
    </Tabs>
  );
}
