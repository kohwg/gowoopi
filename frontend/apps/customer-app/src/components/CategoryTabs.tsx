'use client';

import { Tabs, Tab } from '@heroui/react';
import type { Category } from '@gowoopi/shared';

interface CategoryTabsProps {
  categories: Category[];
  selected: number | null;
  onSelect: (id: number | null) => void;
}

export function CategoryTabs({ categories, selected, onSelect }: CategoryTabsProps) {
  return (
    <Tabs
      selectedKey={selected?.toString() ?? 'all'}
      onSelectionChange={(key) => onSelect(key === 'all' ? null : Number(key))}
      className="mb-4"
    >
      <Tab key="all" title="전체" />
      {categories.map((cat) => (
        <Tab key={cat.id.toString()} title={cat.name} />
      ))}
    </Tabs>
  );
}
