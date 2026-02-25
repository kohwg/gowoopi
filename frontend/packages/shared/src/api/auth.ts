import { useMutation } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { AdminLoginRequest, AuthResponse, CustomerLoginRequest } from '../types';

export function useCustomerLogin() {
  return useMutation({
    mutationFn: async (data: CustomerLoginRequest): Promise<AuthResponse> => {
      const res = await getApiClient().post('/api/customer/login', data);
      return res.data;
    },
  });
}

export function useAdminLogin() {
  return useMutation({
    mutationFn: async (data: AdminLoginRequest): Promise<AuthResponse> => {
      const res = await getApiClient().post('/api/admin/login', data);
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
