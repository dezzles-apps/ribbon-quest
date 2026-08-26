import { useAuthStore } from '@/stores/authStore';

export function useApi() {
  const authStore = useAuthStore();
  async function apiFetch(input: string, init: RequestInit = {}): Promise<Response> {
    const headers = new Headers(init.headers)
    return fetch(input, {
      ...init,
      headers,
    });
  }

  async function apiFetchWithAuth(input: string, init: RequestInit = {}): Promise<Response> {
    const headers = new Headers(init.headers)
    if (authStore.isAuthenticated) {
      headers.set('Authorization', `Bearer ${authStore.token}`);
    }
    return fetch(input, {
      ...init,
      headers,
    });
  }

  return { apiFetch, apiFetchWithAuth }
}
