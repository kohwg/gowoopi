'use client';

import { Navbar, NavbarContent, NavbarItem, Button, Dropdown, DropdownTrigger, DropdownMenu, DropdownItem } from '@heroui/react';
import { useRouter } from 'next/navigation';
import { useAuth, useTranslation, type Locale } from '@gowoopi/shared';

interface HeaderProps {
  storeName: string;
}

export function Header({ storeName }: HeaderProps) {
  const router = useRouter();
  const { logout } = useAuth();
  const { t, locale, setLocale } = useTranslation();

  const handleLogout = () => {
    logout();
    router.push('/login');
  };

  return (
    <Navbar maxWidth="full" className="border-b">
      <NavbarContent justify="start">
        <NavbarItem className="font-semibold">{storeName}</NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem>
          <Dropdown>
            <DropdownTrigger>
              <Button variant="light" size="sm">
                {locale === 'ko' ? '한국어' : 'English'}
              </Button>
            </DropdownTrigger>
            <DropdownMenu onAction={(key) => setLocale(key as Locale)}>
              <DropdownItem key="ko">한국어</DropdownItem>
              <DropdownItem key="en">English</DropdownItem>
            </DropdownMenu>
          </Dropdown>
        </NavbarItem>
        <NavbarItem>
          <Button color="danger" variant="light" size="sm" onPress={handleLogout}>
            {t('common.logout')}
          </Button>
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
