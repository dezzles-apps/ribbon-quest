interface Response<T> {
  data: T
  error: string | null
  errors: Error[] | null
}

interface Page<T> {
  items: T[],
  next: string | null
}

interface Error {
  field: string,
  message: string
}

export type {
  Error,
  Page,
  Response
}