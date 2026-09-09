SELECT
  pokemon.pokemon,
  ribbons.ribbon_key,
  ribbons.name,
  if (pr.achieved_at IS NULL, false, true) as achieved,
  pr.achieved_at,
  ribbons.category
FROM
  pokemon
LEFT JOIN pokemon_games ON pokemon.pokemon = pokemon_games.pokemon 
LEFT JOIN game_ribbons ON pokemon_games.game_key = game_ribbons.game_key
LEFT JOIN ribbons ON game_ribbons.ribbon_key = ribbons.ribbon_key
LEFT JOIN pokemon_ribbons pr on (
  pokemon.pokemon = pr.pokemon
  AND ribbons.ribbon_key = pr.ribbon_key
)
WHERE
  game_ribbons.game_key = ?
ORDER BY pokemon.view_order, ribbons.category ASC, ribbons.view_order ASC
