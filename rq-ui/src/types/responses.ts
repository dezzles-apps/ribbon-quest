interface Response<T> {
  data: T
  error: string | null
}

interface Page<T> {
  items: T[],
  next: string | null
}

export type {
  Page,
  Response
}