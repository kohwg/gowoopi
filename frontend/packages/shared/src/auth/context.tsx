import { createContext, useContext, useState, useCallback, useEffect, type ReactNode } from 'react';
import type { AuthResponse } from '../types';
import { saveAuth, getAuth, clearAuth } from './storage';

interface AuthContextValue {
  auth: AuthResponse | null;
  isAuthenticated: boolean;
  login: (auth: AuthResponse) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextValue | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [auth, setAuth] = useState<AuthResponse | null>(null);

  useEffect(() => {
    setAuth(getAuth());
  }, []);

  const login = useCallback((data: AuthResponse) => {
    saveAuth(data);
    setAuth(data);
  }, []);

  const logout = useCallback(() => {
    clearAuth();
    setAuth(null);
  }, []);

  return (
    <AuthContext.Provider value={{ auth, isAuthenticated: !!auth, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth(): AuthContextValue {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within AuthProvider');
  }
  return context;
}
