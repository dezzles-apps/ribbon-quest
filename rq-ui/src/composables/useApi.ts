
export function useApi() {

  async function apiFetch(input: string, init: RequestInit = {}): Promise<Response> {
    const headers = new Headers(init.headers)
    return fetch(input, {
      ...init,
      headers,
    })
  }

  return { apiFetch }
}
