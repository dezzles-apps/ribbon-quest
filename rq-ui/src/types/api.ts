interface PokemonDetails {
  pokemon: string
  nickname: string
  region: string
  caughtAt: string
  caught: boolean
  nature: string
  notes: string
  characteristic: string
  shiny: boolean
}

interface Pokemon extends PokemonDetails {
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

interface PokemonStats extends PokemonDetails {
  current: number
  total: number
}

interface Game {
  gameKey: string
  name: string
  pokemon: GamePokemon[]
}

interface GamePokemon extends PokemonDetails {
  pokemon: string
  nickname: string
  region: string
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
  PokemonDetails,
  PokemonRibbon,
  PokemonGame,
  PokemonStats,
  Response
}