import type { AuthResponse } from '../types';

const AUTH_KEY = 'gowoopi_auth';

export function saveAuth(auth: AuthResponse): void {
  localStorage.setItem(AUTH_KEY, JSON.stringify(auth));
}

export function getAuth(): AuthResponse | null {
  const data = localStorage.getItem(AUTH_KEY);
  return data ? JSON.parse(data) : null;
}

export function clearAuth(): void {
  localStorage.removeItem(AUTH_KEY);
}
