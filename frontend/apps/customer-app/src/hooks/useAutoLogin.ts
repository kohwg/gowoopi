'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth, useCustomerLogin } from '@gowoopi/shared';
import { useSetupStore } from '@/stores/setup';

export function useAutoLogin() {
  const router = useRouter();
  const { isAuthenticated, login } = useAuth();
  const { storeId, tableNumber, password } = useSetupStore();
  const customerLogin = useCustomerLogin();

  useEffect(() => {
    if (isAuthenticated) return;
    if (!storeId || !tableNumber || !password) {
      router.replace('/setup');
      return;
    }

    customerLogin.mutate(
      { storeId, tableNumber, password },
      {
        onSuccess: (data) => login(data),
        onError: () => router.replace('/setup'),
      }
    );
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated, storeId, tableNumber, password]);

  return { isLoading: customerLogin.isPending, isAuthenticated };
}
