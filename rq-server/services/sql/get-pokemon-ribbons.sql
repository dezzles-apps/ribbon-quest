SELECT DISTINCT
  ribbons.ribbon_key as ribbon_key,
  ribbons.name as ribbon_name,
  pokemon_ribbons.achieved as achieved,
  ribbons.category,
  ribbons.view_order
FROM pokemon
  LEFT JOIN pokemon_games ON pokemon.pokemon = pokemon_games.pokemon
  LEFT JOIN games ON pokemon_games.game_key = games.game_key
  LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
  LEFT JOIN pokemon_ribbons on (
    pokemon_ribbons.ribbon_key = ribbons.ribbon_key
    AND pokemon.pokemon = pokemon_ribbons.pokemon
  )
WHERE pokemon.pokemon = ?
ORDER BY ribbons.category ASC, ribbons.view_order ASC
