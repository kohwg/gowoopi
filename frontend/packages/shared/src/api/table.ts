import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { getApiClient } from './client';
import type { OrderHistory, Table, TableHistoryQuery, TableSetupRequest } from '../types';

export const tableKeys = {
  all: ['tables'] as const,
  history: (tableId: number) => [...tableKeys.all, tableId, 'history'] as const,
};

export function useSetupTable() {
  return useMutation({
    mutationFn: async (data: TableSetupRequest): Promise<{ table: Table; sessionId: string }> => {
      const res = await getApiClient().post('/api/admin/tables/setup', {
        table_number: data.tableNumber,
        password: data.password,
      });
      return res.data;
    },
  });
}

export function useCompleteTable() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (tableId: number): Promise<void> => {
      await getApiClient().post(`/api/admin/tables/${tableId}/complete`);
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: tableKeys.all }),
  });
}

export function useTableHistory(tableId: number, query?: TableHistoryQuery) {
  return useQuery({
    queryKey: tableKeys.history(tableId),
    queryFn: async (): Promise<OrderHistory[]> => {
      const params = new URLSearchParams();
      if (query?.from) params.set('from', query.from);
      if (query?.to) params.set('to', query.to);
      const res = await getApiClient().get(`/api/admin/tables/${tableId}/history?${params}`);
      return res.data.history;
    },
    enabled: tableId > 0,
  });
}
