import type { KeyValuePair } from './data'

interface PokemonDetails {
  pokemon: string
  species: PokemonSpecies
  details: CatchDetails
}

interface PokemonSpecies {
  pokedexId: number
  name: string
  form: string
  imageRef: string
}

interface ShinyPokemon {
  shinyId: number
  species: PokemonSpecies
  details: CatchDetails
}

interface CatchDetails {
  nickname: string
  caughtAt: string
  caught: boolean
  nature: KeyValuePair
  notes: string
  characteristic: KeyValuePair
  shiny: boolean
  game: KeyValuePair
  level: number
  method: KeyValuePair
  gender: KeyValuePair
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
  ribbons: PokemonRibbon[]
}

interface GameWithStats {
  gameKey: string
  name: string
  achieved: number
  total: number
}

interface RibbonStats {
  current: number
  total: number
  pokemon: PokemonDetails[]
}

interface NewRibbonPokemonRequest {
  pokemon: string
  pokedexId: number,
  form: string,
  games: string[]
}

export type {
  Game,
  GamePokemon,
  GameWithStats,
  NewRibbonPokemonRequest,
  Pokemon,
  PokemonDetails,
  PokemonRibbon,
  PokemonGame,
  PokemonStats,
  RibbonStats,
  ShinyPokemon
}