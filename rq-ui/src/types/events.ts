interface Event {
  eventType: string
  eventTime: string
  metadata: EventMetadata
}

interface EventMetadata {
  nickname: string
  pokemon: string
  ribbonKey: string
  ribbonName: string
  ribbonCategory: string
  ribbonType: string
}

export type {
  Event,
  EventMetadata
}