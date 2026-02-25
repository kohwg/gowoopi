'use client';

import { Button } from '@heroui/react';
import { useTranslation } from '@gowoopi/shared';

export function LanguageSwitcher() {
  const { locale, setLocale } = useTranslation();

  const toggle = () => setLocale(locale === 'ko' ? 'en' : 'ko');

  return (
    <Button size="sm" variant="light" onPress={toggle}>
      {locale === 'ko' ? 'EN' : '한국어'}
    </Button>
  );
}
