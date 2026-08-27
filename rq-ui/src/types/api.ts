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
  caughtAt: string
  nature: string
  characteristic: string
  shiny: boolean
  current: number
  target: number
}

interface Game {
  gameKey: string
  name: string
  pokemon: GamePokemon[]
}

interface GamePokemon {
  pokemon: string
  nickname: string
  region: string
  caughtAt: string
  nature: string
  characteristic: string
  shiny: boolean
  ribbons: PokemonRibbon[]
}

interface GameWithStats {
  gameKey: string
  name: string
  achieved: number
  total: number
}

interface Response<T> {
  data: T
  error: string | null
}

export type {
  Game,
  GamePokemon,
  GameWithStats,
  Pokemon,
  PokemonRibbon,
  PokemonGame,
  PokemonStats,
  Response
}