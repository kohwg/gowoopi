import axios, { type AxiosInstance } from 'axios';
import { getAuth } from '../auth/storage';

let apiClient: AxiosInstance | null = null;

export function createApiClient(baseURL: string): AxiosInstance {
  apiClient = axios.create({
    baseURL,
    withCredentials: true,
    headers: { 'Content-Type': 'application/json' },
  });

  // Add auth token to requests
  apiClient.interceptors.request.use((config) => {
    const auth = getAuth();
    if (auth?.accessToken) {
      config.headers.Authorization = `Bearer ${auth.accessToken}`;
    }
    return config;
  });

  apiClient.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config;
      if (error.response?.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true;
        try {
          await apiClient!.post('/api/auth/refresh');
          return apiClient!(originalRequest);
        } catch {
          return Promise.reject(error);
        }
      }
      return Promise.reject(error);
    }
  );

  return apiClient;
}

export function getApiClient(): AxiosInstance {
  if (!apiClient) {
    throw new Error('API client not initialized. Call createApiClient first.');
  }
  return apiClient;
}
