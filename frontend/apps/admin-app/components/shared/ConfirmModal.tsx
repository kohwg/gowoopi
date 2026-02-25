'use client';

import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button } from '@heroui/react';
import { useTranslation } from '@gowoopi/shared';

interface ConfirmModalProps {
  title: string;
  message: string;
  isOpen: boolean;
  onConfirm: () => void;
  onCancel: () => void;
}

export function ConfirmModal({ title, message, isOpen, onConfirm, onCancel }: ConfirmModalProps) {
  const { t } = useTranslation();

  return (
    <Modal isOpen={isOpen} onClose={onCancel}>
      <ModalContent>
        <ModalHeader>{title}</ModalHeader>
        <ModalBody><p>{message}</p></ModalBody>
        <ModalFooter>
          <Button variant="light" onPress={onCancel}>{t('common.cancel')}</Button>
          <Button color="danger" onPress={onConfirm}>{t('common.confirm')}</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
