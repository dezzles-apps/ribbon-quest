interface PokemonData {
  pokedexNo : number
  species: string
  forms: Form[]
}

interface Form {
  form: string
  imageRef: string
}

interface KeyValuePair {
  key: string
  value: string
}

export type {
  PokemonData,
  KeyValuePair,
  Form
}