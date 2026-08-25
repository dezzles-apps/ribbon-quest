interface Pokemon {
  pokemon: string
  nickname: string
  region: string
  ribbons: PokemonRibbon[]
  games: PokemonGame[]
}
interface PokemonRibbon {
  ribbonKey: string
  name: string
  achieved: boolean
  achievedAt: string
  category: string
}

interface PokemonGame {
  gameKey: string
  name: string
  viewOrder: number
  ribbons: string[]
}

interface PokemonStats {
  pokemon: string
  nickname: string
  region: string
  current: number
  target: number
}

interface Response<T> {
  data: T
  error: string | null
}

export type {
  Pokemon,
  PokemonRibbon,
  PokemonGame,
  PokemonStats,
  Response
}