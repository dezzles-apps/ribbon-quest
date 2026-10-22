interface PokedexEntry {
  pokedexNo: PokedexNo
  pokemon: string
  spriteRef: string
  caught: boolean
}

interface PokedexNo {
  number: number
  text: string
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
  PokedexNo,
  PokedexStats
}