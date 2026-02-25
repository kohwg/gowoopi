'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Card, CardBody, Input, Button } from '@heroui/react';
import { useTranslation, useCustomerLogin, useAuth } from '@gowoopi/shared';
import { useSetupStore } from '@/stores/setup';
import { LanguageSwitcher } from '@/components/LanguageSwitcher';

export default function SetupPage() {
  const router = useRouter();
  const { t } = useTranslation();
  const { login } = useAuth();
  const { setSetup } = useSetupStore();
  const customerLogin = useCustomerLogin();

  const [storeId, setStoreId] = useState('00000000-0000-0000-0000-000000000001');
  const [tableNumber, setTableNumber] = useState('1');
  const [password, setPassword] = useState('admin123');
  const [mounted, setMounted] = useState(false);

  useState(() => { setMounted(true); });

  if (!mounted) return null;

  const handleSubmit = () => {
    const tableNum = parseInt(tableNumber, 10);
    customerLogin.mutate(
      { storeId, tableNumber: tableNum, password },
      {
        onSuccess: (data) => {
          setSetup(storeId, tableNum, password);
          login(data);
          router.replace('/');
        },
      }
    );
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <Card className="w-full max-w-md">
        <CardBody className="gap-4">
          <div className="flex justify-between items-center">
            <h1 className="text-xl font-bold">{t('table.setup')}</h1>
            <LanguageSwitcher />
          </div>
          <Input
            label={t('auth.storeId')}
            value={storeId}
            onValueChange={setStoreId}
            autoComplete="off"
          />
          <Input
            label={t('auth.tableNumber')}
            type="number"
            value={tableNumber}
            onValueChange={setTableNumber}
            autoComplete="off"
          />
          <Input
            label={t('auth.password')}
            type="password"
            value={password}
            onValueChange={setPassword}
            autoComplete="off"
          />
          <Button
            color="primary"
            onPress={handleSubmit}
            isLoading={customerLogin.isPending}
          >
            {t('auth.login')}
          </Button>
          {customerLogin.isError && (
            <p className="text-danger text-sm">{t('auth.loginFailed')}</p>
          )}
        </CardBody>
      </Card>
    </div>
  );
}
