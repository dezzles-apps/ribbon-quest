SELECT ribbons.game_key, ribbons.game_name, ribbons.achieved, count(*) FROM (
  SELECT
    games.game_key,
    games.name as game_name,
    games.view_order,
    pokemon.pokemon,
    ribbons.ribbon_key,
    ribbons.name,
    if (pr.achieved_at IS NULL, false, true) as achieved,
    pr.achieved_at,
    ribbons.category
  FROM
    ribbon_pokemon pokemon
  LEFT JOIN ribbon_pokemon_games pokemon_games ON pokemon.ribbon_pokemon_id = pokemon_games.ribbon_pokemon_id 
  LEFT JOIN game_ribbons ON pokemon_games.game_key = game_ribbons.game_key
  LEFT JOIN games ON games.game_key = game_ribbons.game_key
  LEFT JOIN ribbons ON game_ribbons.ribbon_key = ribbons.ribbon_key
  LEFT JOIN ribbons_earned pr on (
    pokemon.ribbon_pokemon_id = pr.ribbon_pokemon_id
    AND ribbons.ribbon_key = pr.ribbon_key
  )
  ORDER BY games.view_order, pokemon.view_order, ribbons.category ASC, ribbons.view_order ASC
) as ribbons
GROUP BY ribbons.game_key, ribbons.game_name, achieved
ORDER BY ribbons.view_order ASC