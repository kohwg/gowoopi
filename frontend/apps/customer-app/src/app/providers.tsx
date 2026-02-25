'use client';

import { HeroUIProvider } from '@heroui/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider, I18nProvider, createApiClient } from '@gowoopi/shared';
import { useState, type ReactNode } from 'react';

export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(() => new QueryClient());

  useState(() => {
    createApiClient(process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080');
  });

  return (
    <QueryClientProvider client={queryClient}>
      <HeroUIProvider>
        <AuthProvider>
          <I18nProvider defaultLocale="ko">{children}</I18nProvider>
        </AuthProvider>
      </HeroUIProvider>
    </QueryClientProvider>
  );
}
