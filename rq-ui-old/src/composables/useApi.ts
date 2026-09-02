import { useAuthStore } from '@/stores/authStore';
import { useRouter } from 'vue-router'
export function useApi() {
  const authStore = useAuthStore();
  const router = useRouter();
  async function apiFetch(input: string, init: RequestInit = {}): Promise<Response> {
    const headers = new Headers(init.headers)
    if (authStore.isAuthenticated) {
      headers.set('Authorization', `Bearer ${authStore.token}`);
    }
    return fetch(input, {
      ...init,
      headers,
    }).then(resp => {
      if (resp.status == 401) {
        authStore.clearToken()
        router.push('/login')

      }
      return resp
    });
  }

  return { apiFetch }
}
