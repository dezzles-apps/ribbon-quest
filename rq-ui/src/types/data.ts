interface PokemonData {
  pokedexNo : number
  species: string
  forms: Form[]
}

interface Form {
  formId: string
  form: string
  imageRef: string
}

interface KeyValuePair {
  key: string
  value: string
}

interface CommonData {
  games: KeyValuePair[]
  characteristics: KeyValuePair[]
  natures: KeyValuePair[]
  methods: KeyValuePair[]
  genders: KeyValuePair[]
}

export type {
  CommonData,
  PokemonData,
  KeyValuePair,
  Form
}