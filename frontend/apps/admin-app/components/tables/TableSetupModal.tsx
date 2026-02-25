'use client';

import { useState } from 'react';
import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button, Input } from '@heroui/react';
import { useTranslation, type TableSetupRequest } from '@gowoopi/shared';

interface TableSetupModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: TableSetupRequest) => void;
}

export function TableSetupModal({ isOpen, onClose, onSubmit }: TableSetupModalProps) {
  const { t } = useTranslation();
  const [tableNumber, setTableNumber] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = () => {
    onSubmit({ tableNumber: parseInt(tableNumber, 10), password });
    setTableNumber('');
    setPassword('');
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>{t('table.setup.title')}</ModalHeader>
        <ModalBody className="flex flex-col gap-4">
          <Input
            label={t('table.setup.number')}
            type="number"
            value={tableNumber}
            onChange={(e) => setTableNumber(e.target.value)}
            required
          />
          <Input
            label={t('table.setup.password')}
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </ModalBody>
        <ModalFooter>
          <Button variant="light" onPress={onClose}>{t('common.cancel')}</Button>
          <Button color="primary" onPress={handleSubmit}>{t('common.save')}</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
