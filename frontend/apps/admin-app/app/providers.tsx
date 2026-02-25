'use client';

import { HeroUIProvider } from '@heroui/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider, I18nProvider, createApiClient } from '@gowoopi/shared';
import { useState, type ReactNode } from 'react';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(() => new QueryClient());

  useState(() => createApiClient(API_URL));

  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <I18nProvider defaultLocale="ko">
          <HeroUIProvider>{children}</HeroUIProvider>
        </I18nProvider>
      </AuthProvider>
    </QueryClientProvider>
  );
}
