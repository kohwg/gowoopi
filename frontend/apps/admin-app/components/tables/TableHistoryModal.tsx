'use client';

import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button, Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, Input, Spinner } from '@heroui/react';
import { useState } from 'react';
import { useTableHistory, useTranslation } from '@gowoopi/shared';

interface TableHistoryModalProps {
  tableId: number;
  isOpen: boolean;
  onClose: () => void;
}

export function TableHistoryModal({ tableId, isOpen, onClose }: TableHistoryModalProps) {
  const { t } = useTranslation();
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const { data: history, isLoading } = useTableHistory(tableId, { from: from || undefined, to: to || undefined });

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="2xl">
      <ModalContent>
        <ModalHeader>{t('table.history.title')}</ModalHeader>
        <ModalBody>
          <div className="flex gap-4 mb-4">
            <Input type="date" label={t('table.history.from')} value={from} onChange={(e) => setFrom(e.target.value)} />
            <Input type="date" label={t('table.history.to')} value={to} onChange={(e) => setTo(e.target.value)} />
          </div>
          {isLoading ? (
            <div className="flex justify-center p-4"><Spinner /></div>
          ) : (
            <Table aria-label="Order history">
              <TableHeader>
                <TableColumn>{t('order.id')}</TableColumn>
                <TableColumn>{t('order.time')}</TableColumn>
                <TableColumn>{t('order.total')}</TableColumn>
                <TableColumn>{t('order.completedAt')}</TableColumn>
              </TableHeader>
              <TableBody emptyContent={t('table.history.empty')}>
                {(history || []).map((h) => (
                  <TableRow key={h.id}>
                    <TableCell>{h.id.slice(0, 8)}</TableCell>
                    <TableCell>{new Date(h.createdAt).toLocaleString()}</TableCell>
                    <TableCell>₩{h.totalAmount.toLocaleString()}</TableCell>
                    <TableCell>{new Date(h.completedAt).toLocaleString()}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          )}
        </ModalBody>
        <ModalFooter>
          <Button onPress={onClose}>{t('common.close')}</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
