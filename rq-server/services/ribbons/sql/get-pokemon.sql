SELECT
  ribbon_pokemon.pokemon AS pokemon,
  ribbon_pokemon.nickname AS nickname,
  games.game_key AS game_key,
  games.name AS game_name,
  games.view_order AS game_view_order,
  ribbons.ribbon_key AS ribbon_key,
  ribbons.name AS ribbon_name,
  IF (ribbons_earned.achieved_at IS NULL, false, true) AS achieved
FROM ribbon_pokemon
  LEFT JOIN ribbon_pokemon_games ON ribbon_pokemon.ribbon_pokemon_id = ribbon_pokemon_games.ribbon_pokemon_id
  LEFT JOIN games ON ribbon_pokemon_games.game_key = games.game_key
  LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
  LEFT JOIN ribbons_earned on ribbons_earned.ribbon_key = ribbons.ribbon_key
WHERE ribbon_pokemon.pokemon = ?
ORDER BY game_view_order ASC
