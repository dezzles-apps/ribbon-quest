interface PokemonData {
  pokedexNo : number
  species: string
  forms: Form[]
}

interface Form {
  form: string
  imageRef: string
}

export type {
  PokemonData,
  Form
}