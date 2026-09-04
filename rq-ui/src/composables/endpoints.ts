export const Auth = {
  Login: '/api/auth/v1/login',
};

export const Ribbons = {
  GetAllGames: '/api/ribbons/v1/games',
  GetGame: (gameKey: string) => `/api/ribbons/v1/games/${gameKey}`,
  GetAllPokemon: '/api/ribbons/v1/pokemon',
  GetPokemon: (pokemon: string) => `/api/ribbons/v1/pokemon/${pokemon}`,
  CatchPokemon: (pokemon: string) => `/api/ribbons/v1/pokemon/${pokemon}/catch`,
  UpdatePokemon: (pokemon: string) => `/api/ribbons/v1/pokemon/${pokemon}`,
};

export default {
  Auth,
  Ribbons,
}