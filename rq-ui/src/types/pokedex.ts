interface PokedexEntry {
  pokedexNo: string
  pokemon: string
  spriteRef: string
  caught: boolean
}

interface PokedexStats {
  caught: number,
  total: number
}

interface Pokedex {
  name: string,
  entries: PokedexEntry[]
  stats: PokedexStats
  image: string
  notes: string
}

export type {
  PokedexEntry,
  Pokedex,
  PokedexStats
}