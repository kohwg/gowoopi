'use client';

import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button, Select, SelectItem, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@heroui/react';
import type { Order, OrderStatus } from '@gowoopi/shared';
import { useTranslation } from '@gowoopi/shared';
import { StatusBadge } from './StatusBadge';

interface OrderDetailModalProps {
  order: Order | null;
  isOpen: boolean;
  onClose: () => void;
  onStatusChange: (status: OrderStatus) => void;
  onDelete: () => void;
}

const statuses: OrderStatus[] = ['PENDING', 'CONFIRMED', 'PREPARING', 'COMPLETED'];

export function OrderDetailModal({ order, isOpen, onClose, onStatusChange, onDelete }: OrderDetailModalProps) {
  const { t } = useTranslation();

  if (!order) return null;

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="lg">
      <ModalContent>
        <ModalHeader className="flex justify-between items-center">
          <span>{t('order.detail')}: {order.id.slice(0, 8)}</span>
          <StatusBadge status={order.status} />
        </ModalHeader>
        <ModalBody>
          <Table aria-label="Order items">
            <TableHeader>
              <TableColumn>{t('menu.name')}</TableColumn>
              <TableColumn>{t('order.quantity')}</TableColumn>
              <TableColumn>{t('menu.price')}</TableColumn>
            </TableHeader>
            <TableBody>
              {order.items.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>{item.menuName}</TableCell>
                  <TableCell>{item.quantity}</TableCell>
                  <TableCell>₩{item.subtotal.toLocaleString()}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
          <div className="flex justify-between items-center mt-4">
            <span className="font-bold">{t('order.total')}</span>
            <span className="text-xl font-bold">₩{order.totalAmount.toLocaleString()}</span>
          </div>
          <Select
            label={t('order.status.label')}
            selectedKeys={[order.status]}
            onChange={(e) => onStatusChange(e.target.value as OrderStatus)}
            className="mt-4"
          >
            {statuses.map((s) => (
              <SelectItem key={s}>{t(`order.status.${s.toLowerCase()}`)}</SelectItem>
            ))}
          </Select>
        </ModalBody>
        <ModalFooter>
          <Button color="danger" variant="light" onPress={onDelete}>
            {t('common.delete')}
          </Button>
          <Button onPress={onClose}>{t('common.close')}</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
