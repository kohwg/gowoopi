'use client';

import { useState, useMemo } from 'react';
import { Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Chip, Spinner } from '@heroui/react';
import { useAdminOrders, useSetupTable, useCompleteTable, useTranslation } from '@gowoopi/shared';
import { TableSetupModal, TableHistoryModal } from '@/components/tables';
import { ConfirmModal } from '@/components/shared/ConfirmModal';

export default function TablesPage() {
  const { t } = useTranslation();
  const { data: orders, isLoading } = useAdminOrders();
  const setupTable = useSetupTable();
  const completeTable = useCompleteTable();

  const [showSetup, setShowSetup] = useState(false);
  const [completeTarget, setCompleteTarget] = useState<number | null>(null);
  const [historyTarget, setHistoryTarget] = useState<number | null>(null);

  const tableData = useMemo(() => {
    if (!orders) return [];
    const map = new Map<number, { total: number; orderCount: number }>();
    for (const o of orders) {
      const existing = map.get(o.tableId) || { total: 0, orderCount: 0 };
      existing.total += o.totalAmount;
      existing.orderCount += 1;
      map.set(o.tableId, existing);
    }
    return Array.from(map.entries())
      .map(([id, data]) => ({ id, ...data }))
      .sort((a, b) => a.id - b.id);
  }, [orders]);

  const handleSetup = async (data: { tableNumber: number; password: string }) => {
    await setupTable.mutateAsync(data);
  };

  const handleComplete = async () => {
    if (completeTarget === null) return;
    await completeTable.mutateAsync(completeTarget);
    setCompleteTarget(null);
  };

  if (isLoading) {
    return <div className="flex justify-center p-8"><Spinner size="lg" /></div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">{t('table.management')}</h1>
        <Button color="primary" onPress={() => setShowSetup(true)}>{t('table.setup.button')}</Button>
      </div>

      <Table aria-label="Tables">
        <TableHeader>
          <TableColumn>{t('table.number.label')}</TableColumn>
          <TableColumn>{t('table.status')}</TableColumn>
          <TableColumn>{t('table.totalAmount')}</TableColumn>
          <TableColumn>{t('common.actions')}</TableColumn>
        </TableHeader>
        <TableBody emptyContent={t('table.empty')}>
          {tableData.map((table) => (
            <TableRow key={table.id}>
              <TableCell>{table.id}</TableCell>
              <TableCell>
                <Chip color={table.orderCount > 0 ? 'success' : 'default'} size="sm">
                  {table.orderCount > 0 ? t('table.active') : t('table.inactive')}
                </Chip>
              </TableCell>
              <TableCell>₩{table.total.toLocaleString()}</TableCell>
              <TableCell>
                <div className="flex gap-2">
                  <Button size="sm" variant="light" onPress={() => setHistoryTarget(table.id)}>
                    {t('table.history.button')}
                  </Button>
                  <Button
                    size="sm"
                    color="warning"
                    variant="flat"
                    onPress={() => setCompleteTarget(table.id)}
                    isDisabled={table.orderCount === 0}
                  >
                    {t('table.complete')}
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <TableSetupModal isOpen={showSetup} onClose={() => setShowSetup(false)} onSubmit={handleSetup} />
      <TableHistoryModal tableId={historyTarget || 0} isOpen={historyTarget !== null} onClose={() => setHistoryTarget(null)} />
      <ConfirmModal
        title={t('table.completeConfirm.title')}
        message={t('table.completeConfirm.message')}
        isOpen={completeTarget !== null}
        onConfirm={handleComplete}
        onCancel={() => setCompleteTarget(null)}
      />
    </div>
  );
}
