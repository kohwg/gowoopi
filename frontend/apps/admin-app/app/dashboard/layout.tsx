'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@gowoopi/shared';
import { Sidebar } from '@/components/layout/Sidebar';
import { Header } from '@/components/layout/Header';
import { useSSE } from '@/hooks/use-sse';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const { auth, isAuthenticated } = useAuth();

  useSSE(API_URL);

  useEffect(() => {
    if (!isAuthenticated) router.replace('/login');
  }, [isAuthenticated, router]);

  if (!isAuthenticated) return null;

  return (
    <div className="flex min-h-screen bg-gray-50">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Header storeName={auth?.storeName || ''} />
        <main className="flex-1 p-6">{children}</main>
      </div>
    </div>
  );
}
