'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Listbox, ListboxItem } from '@heroui/react';
import { useTranslation } from '@gowoopi/shared';

const navItems = [
  { key: '/dashboard', labelKey: 'nav.dashboard' },
  { key: '/dashboard/tables', labelKey: 'nav.tables' },
  { key: '/dashboard/menus', labelKey: 'nav.menus' },
];

export function Sidebar() {
  const pathname = usePathname();
  const { t } = useTranslation();

  return (
    <aside className="w-64 bg-white border-r min-h-screen p-4">
      <h2 className="text-xl font-bold mb-6">{t('app.title')}</h2>
      <Listbox aria-label="Navigation" selectionMode="single" selectedKeys={[pathname]}>
        {navItems.map((item) => (
          <ListboxItem key={item.key} as={Link} href={item.key}>
            {t(item.labelKey)}
          </ListboxItem>
        ))}
      </Listbox>
    </aside>
  );
}
