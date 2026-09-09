SELECT pokemon, nickname, region, caught_at, nature, characteristic, notes, shiny, achieved, COUNT(*) FROM (
  SELECT DISTINCT
    pokemon.pokemon,
    pokemon.nickname,
    pokemon.region,
    pokemon.caught_at,
    pokemon.nature,
    pokemon.characteristic,
    pokemon.notes,
    pokemon.shiny,
    pokemon.view_order,
    ribbons.ribbon_key as ribbon_key,
    ribbons.name as ribbon_name,
    pokemon_ribbons.achieved as achieved
  FROM pokemon
    LEFT JOIN pokemon_games ON pokemon.pokemon = pokemon_games.pokemon
    LEFT JOIN games ON pokemon_games.game_key = games.game_key
    LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
    LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
    LEFT JOIN pokemon_ribbons on (
      pokemon_ribbons.ribbon_key = ribbons.ribbon_key
      AND pokemon.pokemon = pokemon_ribbons.pokemon
    )
  ORDER BY pokemon.view_order ASC
) as ribbon_types
GROUP BY pokemon, ribbon_types.achieved 