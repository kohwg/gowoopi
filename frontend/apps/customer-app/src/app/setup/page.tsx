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
    <div className="min-h-screen flex items-center justify-center p-4 bg-gradient-to-br from-primary-50 via-background to-secondary-50">
      <Card className="w-full max-w-md shadow-2xl">
        <CardBody className="gap-6 p-8">
          <div className="flex justify-between items-center">
            <h1 className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
              {t('table.setup')}
            </h1>
            <LanguageSwitcher />
          </div>
          <Input
            label={t('auth.storeId')}
            value={storeId}
            onValueChange={setStoreId}
            autoComplete="off"
            variant="bordered"
            size="lg"
          />
          <Input
            label={t('auth.tableNumber')}
            type="number"
            value={tableNumber}
            onValueChange={setTableNumber}
            autoComplete="off"
            variant="bordered"
            size="lg"
          />
          <Input
            label={t('auth.password')}
            type="password"
            value={password}
            onValueChange={setPassword}
            autoComplete="off"
            variant="bordered"
            size="lg"
          />
          <Button
            color="primary"
            onPress={handleSubmit}
            isLoading={customerLogin.isPending}
            size="lg"
            className="font-semibold shadow-lg"
          >
            {t('auth.login')}
          </Button>
          {customerLogin.isError && (
            <p className="text-danger text-sm text-center">{t('auth.loginFailed')}</p>
          )}
        </CardBody>
      </Card>
    </div>
  );
}
