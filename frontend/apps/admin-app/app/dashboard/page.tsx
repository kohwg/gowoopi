'use client';

import { useState } from 'react';
import { Spinner } from '@heroui/react';
import { useAdminOrders, useUpdateOrderStatus, useDeleteOrder, useTranslation, type Order, type OrderStatus } from '@gowoopi/shared';
import { OrderGrid, OrderDetailModal } from '@/components/orders';
import { ConfirmModal } from '@/components/shared/ConfirmModal';
import { useSSEStore } from '@/stores/sse-store';

export default function DashboardPage() {
  const { t } = useTranslation();
  const { data: orders, isLoading } = useAdminOrders();
  const updateStatus = useUpdateOrderStatus();
  const deleteOrder = useDeleteOrder();
  const { clearNewOrder } = useSSEStore();

  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Order | null>(null);

  const handleTableClick = (tableId: number) => {
    const tableOrderList = orders?.filter((o) => o.tableId === tableId) || [];
    if (tableOrderList.length > 0) {
      setSelectedOrder(tableOrderList[0]);
      tableOrderList.forEach((o) => clearNewOrder(o.id));
    }
  };

  const handleStatusChange = async (status: OrderStatus) => {
    if (!selectedOrder) return;
    await updateStatus.mutateAsync({ id: selectedOrder.id, data: { status } });
    setSelectedOrder((prev) => (prev ? { ...prev, status } : null));
  };

  const handleDelete = async () => {
    if (!deleteTarget) return;
    await deleteOrder.mutateAsync(deleteTarget.id);
    setDeleteTarget(null);
    setSelectedOrder(null);
  };

  if (isLoading) {
    return <div className="flex justify-center p-8"><Spinner size="lg" /></div>;
  }

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6">{t('dashboard.title')}</h1>
      <OrderGrid orders={orders || []} onOrderClick={handleTableClick} />
      <OrderDetailModal
        order={selectedOrder}
        isOpen={!!selectedOrder}
        onClose={() => setSelectedOrder(null)}
        onStatusChange={handleStatusChange}
        onDelete={() => {
          setDeleteTarget(selectedOrder);
        }}
      />
      <ConfirmModal
        title={t('order.deleteConfirm.title')}
        message={t('order.deleteConfirm.message')}
        isOpen={!!deleteTarget}
        onConfirm={handleDelete}
        onCancel={() => setDeleteTarget(null)}
      />
    </div>
  );
}
