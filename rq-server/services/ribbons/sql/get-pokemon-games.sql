SELECT
  games.game_key as game_key,
  games.name as game_name,
  games.view_order as game_view_order,
  game_ribbons.ribbon_key as ribbon_key
FROM ribbon_pokemon pokemon
  LEFT JOIN ribbon_pokemon_games ON pokemon.ribbon_pokemon_id = ribbon_pokemon_games.ribbon_pokemon_id
  LEFT JOIN games ON ribbon_pokemon_games.game_key = games.game_key
  LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons ON ribbons.ribbon_key = game_ribbons.ribbon_key 
WHERE pokemon.pokemon = ?
ORDER BY game_view_order ASC, ribbons.category ASC, ribbons.view_order ASC
