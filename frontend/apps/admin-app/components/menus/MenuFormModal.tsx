'use client';

import { useState, useEffect } from 'react';
import { Modal, ModalContent, ModalHeader, ModalBody, ModalFooter, Button, Input, Textarea, Select, SelectItem } from '@heroui/react';
import { useTranslation, type Menu, type Category, type MenuCreateRequest, type MenuUpdateRequest } from '@gowoopi/shared';

interface MenuFormModalProps {
  menu?: Menu;
  categories: Category[];
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: MenuCreateRequest | MenuUpdateRequest) => void;
}

export function MenuFormModal({ menu, categories, isOpen, onClose, onSubmit }: MenuFormModalProps) {
  const { t } = useTranslation();
  const [categoryId, setCategoryId] = useState('');
  const [name, setName] = useState('');
  const [price, setPrice] = useState('');
  const [description, setDescription] = useState('');
  const [imageUrl, setImageUrl] = useState('');

  useEffect(() => {
    if (menu) {
      setCategoryId(String(menu.categoryId));
      setName(menu.name);
      setPrice(String(menu.price));
      setDescription(menu.description || '');
      setImageUrl(menu.imageUrl || '');
    } else {
      setCategoryId(categories[0]?.id ? String(categories[0].id) : '');
      setName('');
      setPrice('');
      setDescription('');
      setImageUrl('');
    }
  }, [menu, categories, isOpen]);

  const handleSubmit = () => {
    const catId = parseInt(categoryId, 10);
    const priceVal = parseInt(price, 10);
    
    if (isNaN(catId) || isNaN(priceVal)) {
      console.error('Invalid categoryId or price');
      return;
    }
    
    const data = {
      categoryId: catId,
      name,
      price: priceVal,
      description: description || undefined,
      imageUrl: imageUrl || undefined,
    };
    onSubmit(data);
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalContent>
        <ModalHeader>{menu ? t('menu.edit') : t('menu.add')}</ModalHeader>
        <ModalBody className="flex flex-col gap-4">
          <Select
            label={t('menu.category')}
            selectedKeys={categoryId ? [categoryId] : []}
            onChange={(e) => setCategoryId(e.target.value)}
          >
            {categories.map((cat) => (
              <SelectItem key={String(cat.id)}>{cat.name}</SelectItem>
            ))}
          </Select>
          <Input label={t('menu.name')} value={name} onChange={(e) => setName(e.target.value)} required />
          <Input label={t('menu.price')} type="number" value={price} onChange={(e) => setPrice(e.target.value)} required />
          <Textarea label={t('menu.description')} value={description} onChange={(e) => setDescription(e.target.value)} />
          <Input label={t('menu.imageUrl')} value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} />
        </ModalBody>
        <ModalFooter>
          <Button variant="light" onPress={onClose}>{t('common.cancel')}</Button>
          <Button color="primary" onPress={handleSubmit}>{t('common.save')}</Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
}
