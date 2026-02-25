import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { Order, OrderCreateRequest, StatusUpdateRequest } from '../types';

export const orderKeys = {
  all: ['orders'] as const,
  customer: () => [...orderKeys.all, 'customer'] as const,
  admin: () => [...orderKeys.all, 'admin'] as const,
};

export function useCustomerOrders() {
  return useQuery({
    queryKey: orderKeys.customer(),
    queryFn: async (): Promise<Order[]> => {
      const res = await getApiClient().get('/api/customer/orders');
      return res.data ?? [];
    },
  });
}

export function useAdminOrders() {
  return useQuery({
    queryKey: orderKeys.admin(),
    queryFn: async (): Promise<Order[]> => {
      const res = await getApiClient().get('/api/admin/orders');
      return res.data ?? [];
    },
  });
}

export function useCreateOrder() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: OrderCreateRequest): Promise<Order> => {
      const res = await getApiClient().post('/api/customer/orders', {
        items: data.items.map((i) => ({ menu_id: i.menuId, quantity: i.quantity })),
      });
      return res.data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: orderKeys.all }),
  });
}

export function useUpdateOrderStatus() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: StatusUpdateRequest }): Promise<Order> => {
      const res = await getApiClient().patch(`/api/admin/orders/${id}/status`, data);
      return res.data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: orderKeys.all }),
  });
}

export function useDeleteOrder() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (id: string): Promise<void> => {
      await getApiClient().delete(`/api/admin/orders/${id}`);
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: orderKeys.all }),
  });
}
