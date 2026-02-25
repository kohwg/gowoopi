'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Card, CardBody, CardHeader, Input, Button } from '@heroui/react';
import { useAdminLogin, useAuth, useTranslation } from '@gowoopi/shared';

export default function LoginPage() {
  const router = useRouter();
  const { login } = useAuth();
  const { t } = useTranslation();
  const loginMutation = useAdminLogin();

  const [storeId, setStoreId] = useState('00000000-0000-0000-0000-000000000001');
  const [username, setUsername] = useState('admin');
  const [password, setPassword] = useState('admin123');
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    try {
      const res = await loginMutation.mutateAsync({ storeId, username, password });
      login(res);
      router.push('/dashboard');
    } catch {
      setError(t('login.error'));
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Card className="w-full max-w-md">
        <CardHeader className="flex justify-center">
          <h1 className="text-2xl font-bold">{t('login.title')}</h1>
        </CardHeader>
        <CardBody>
          <form onSubmit={handleSubmit} className="flex flex-col gap-4">
            <Input
              label={t('login.storeId')}
              value={storeId}
              onChange={(e) => setStoreId(e.target.value)}
              required
            />
            <Input
              label={t('login.username')}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
            <Input
              label={t('login.password')}
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            {error && <p className="text-red-500 text-sm">{error}</p>}
            <Button type="submit" color="primary" isLoading={loginMutation.isPending}>
              {t('login.submit')}
            </Button>
          </form>
        </CardBody>
      </Card>
    </div>
  );
}
