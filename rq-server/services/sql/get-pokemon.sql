SELECT
  pokemon.pokemon as pokemon,
  pokemon.nickname as nickname,
  pokemon.region as region,
  games.game_key as game_key,
  games.name as game_name,
  games.view_order as game_view_order,
  ribbons.ribbon_key as ribbon_key,
  ribbons.name as ribbon_name,
  pokemon_ribbons.achieved as achieved
FROM pokemon
  LEFT JOIN pokemon_games ON pokemon.pokemon = pokemon_games.pokemon
  LEFT JOIN games ON pokemon_games.game_key = games.game_key
  LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
  LEFT JOIN pokemon_ribbons on pokemon_ribbons.ribbon_key = ribbons.ribbon_key
WHERE pokemon.pokemon = ?
ORDER BY game_view_order ASC