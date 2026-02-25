import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface SetupState {
  storeId: string;
  tableNumber: number;
  password: string;
  setSetup: (storeId: string, tableNumber: number, password: string) => void;
  clear: () => void;
}

export const useSetupStore = create<SetupState>()(
  persist(
    (set) => ({
      storeId: '',
      tableNumber: 0,
      password: '',
      setSetup: (storeId, tableNumber, password) => set({ storeId, tableNumber, password }),
      clear: () => set({ storeId: '', tableNumber: 0, password: '' }),
    }),
    { name: 'gowoopi-setup' }
  )
);
