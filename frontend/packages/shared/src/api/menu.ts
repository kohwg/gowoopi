import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { Menu, MenuCreateRequest, MenuOrderItem, MenuUpdateRequest } from '../types';

export const menuKeys = {
  all: ['menus'] as const,
  byStore: (storeId: string) => [...menuKeys.all, storeId] as const,
};

export function useMenus(storeId: string, role: 'customer' | 'admin' = 'customer') {
  return useQuery({
    queryKey: menuKeys.byStore(storeId),
    queryFn: async (): Promise<Menu[]> => {
      const endpoint = role === 'admin' ? '/api/admin/menus' : '/api/customer/menus';
      const res = await getApiClient().get(endpoint);
      const data = res.data ?? [];
      return data.map((m: Record<string, unknown>) => ({
        id: m.ID,
        storeId: m.StoreID,
        categoryId: m.CategoryID,
        name: m.Name,
        price: m.Price,
        description: m.Description,
        imageUrl: m.ImageURL,
        displayOrder: m.SortOrder,
        category: m.Category ? {
          id: (m.Category as Record<string, unknown>).ID,
          name: (m.Category as Record<string, unknown>).Name,
        } : undefined,
      }));
    },
    enabled: !!storeId,
  });
}

export function useCreateMenu() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (data: MenuCreateRequest): Promise<Menu> => {
      const res = await getApiClient().post('/api/admin/menus', {
        category_id: data.categoryId,
        name: data.name,
        price: data.price,
        description: data.description,
        image_url: data.imageUrl,
      });
      return res.data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}

export function useUpdateMenu() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ id, data }: { id: number; data: MenuUpdateRequest }): Promise<Menu> => {
      const res = await getApiClient().put(`/api/admin/menus/${id}`, {
        category_id: data.categoryId,
        name: data.name,
        price: data.price,
        description: data.description,
        image_url: data.imageUrl,
      });
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
      await getApiClient().patch('/api/admin/menus/order', items.map((i) => ({ id: i.id, sort_order: i.displayOrder })));
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: menuKeys.all }),
  });
}
