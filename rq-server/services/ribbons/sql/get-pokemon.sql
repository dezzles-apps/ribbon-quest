SELECT
  pokemon.pokemon AS pokemon,
  pokemon.nickname AS nickname,
  pokemon.region AS region,
  games.game_key AS game_key,
  games.name AS game_name,
  games.view_order AS game_view_order,
  ribbons.ribbon_key AS ribbon_key,
  ribbons.name AS ribbon_name,
  IF (pokemon_ribbons.achieved_at IS NULL, false, true) AS achieved,
FROM pokemon
  LEFT JOIN pokemon_games ON pokemon.pokemon = pokemon_games.pokemon
  LEFT JOIN games ON pokemon_games.game_key = games.game_key
  LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
  LEFT JOIN pokemon_ribbons on pokemon_ribbons.ribbon_key = ribbons.ribbon_key
WHERE pokemon.pokemon = ?
ORDER BY game_view_order ASC