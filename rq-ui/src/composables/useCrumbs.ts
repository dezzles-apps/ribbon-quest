import { useRoute, useRouter } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router';

interface CrumbData {
  title: string | null | ((params: Record<string, any>) => string)
  href: string | null | ((params: Record<string, any>) => string)
  parent: string | null
}


function process(
  value : string | null | ((params: Record<string, any>) => string),
  params: Record<string, any>
) : string {
  if (typeof value === 'string') {
    return value as string;
  } else if (typeof value === 'function') {
    return value(params)
  }
  return 'Default'
}

export function useCrumbs() {
  const router = useRouter();
  const route = useRoute();
  function getCrumbs() : any[] {
    const routes = router.options.routes.reduce((prev, n) => {
      if (n.name) {
        prev.set(n.name as string, n);
      }
      return prev;
    }, new Map<string | null, RouteRecordRaw>())

    let result = [] as any[]
    let next = route.matched[0] as RouteRecordRaw | undefined
    while (next && next.meta && next.meta.crumbs) {
      let crumbData = next.meta.crumbs as CrumbData
      
      result.unshift({
        title: process(crumbData.title, route.params),
        disabled: false,
        href: process(crumbData.href, route.params),
      })

      next = routes.get(crumbData.parent)
    }
    return result
  }

  return { getCrumbs }
}
