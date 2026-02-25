import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { Menu, MenuCreateRequest, MenuOrderItem, MenuUpdateRequest } from '../types';

export const menuKeys = {
  all: ['menus'] as const,
  byStore: (storeId: string) => [...menuKeys.all, storeId] as const,
};

export function useMenus(storeId: string) {
  return useQuery({
    queryKey: menuKeys.byStore(storeId),
    queryFn: async (): Promise<Menu[]> => {
      const res = await getApiClient().get('/api/customer/menus');
      return res.data.menus;
    },
    enabled: !!storeId,
  });
}

export function useCreateMenu() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: MenuCreateRequest): Promise<Menu> => {
      const res = await getApiClient().post('/api/admin/menus', data);
      return res.data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}

export function useUpdateMenu() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: number; data: MenuUpdateRequest }): Promise<Menu> => {
      const res = await getApiClient().put(`/api/admin/menus/${id}`, data);
      return res.data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}

export function useDeleteMenu() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (id: number): Promise<void> => {
      await getApiClient().delete(`/api/admin/menus/${id}`);
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}

export function useUpdateMenuOrder() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (items: MenuOrderItem[]): Promise<void> => {
      await getApiClient().patch('/api/admin/menus/order', { items });
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}
