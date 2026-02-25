import { useMutation } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { AdminLoginRequest, AuthResponse, CustomerLoginRequest } from '../types';

export function useCustomerLogin() {
  return useMutation({
    mutationFn: async (data: CustomerLoginRequest): Promise<AuthResponse> => {
      const res = await getApiClient().post('/api/customer/login', {
        store_id: data.storeId,
        table_number: data.tableNumber,
        password: data.password,
      });
      return res.data;
    },
  });
}

export function useAdminLogin() {
  return useMutation({
    mutationFn: async (data: AdminLoginRequest): Promise<AuthResponse> => {
      const res = await getApiClient().post('/api/admin/login', {
        store_id: data.storeId,
        username: data.username,
        password: data.password,
      });
      return res.data;
    },
  });
}

export function useRefreshToken() {
  return useMutation({
    mutationFn: async (): Promise<void> => {
      await getApiClient().post('/api/auth/refresh');
    },
  });
}
