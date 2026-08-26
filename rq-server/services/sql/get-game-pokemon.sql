SELECT
  pokemon.pokemon,
  pokemon.nickname,
  pokemon.region,
  pokemon.caught_at,
  pokemon.nature,
  pokemon.characteristic,
  pokemon.shiny
FROM pokemon, pokemon_games
WHERE
  pokemon.pokemon = pokemon_games.pokemon
  AND pokemon_games.game_key = ?
ORDER BY pokemon.view_order 