SELECT pokemon, pokedex_id, species, form_name, sprite_ref, nickname, caught_at, nature, characteristic, notes, shiny, achieved, COUNT(*) FROM (
  SELECT DISTINCT
    pokemon.pokemon,
    species.pokedex_id,
    species.species,
    form.form_name,
    form.sprite_ref,
    pokemon.nickname,
    pokemon.caught_at,
    pokemon.nature,
    pokemon.characteristic,
    pokemon.notes,
    pokemon.shiny,
    pokemon.view_order,
    ribbons.ribbon_key as ribbon_key,
    ribbons.name as ribbon_name,
    ribbons_earned.achieved as achieved
  FROM ribbon_pokemon pokemon
  	LEFT JOIN pokemon_data species ON pokemon.pokedex_id = species.pokedex_id
  	LEFT JOIN pokemon_forms form ON pokemon.form_id = form.form_id
    LEFT JOIN ribbon_pokemon_games pokemon_games ON pokemon.ribbon_pokemon_id = pokemon_games.ribbon_pokemon_id
    LEFT JOIN games ON pokemon_games.game_key = games.game_key
    LEFT JOIN game_ribbons ON games.game_key = game_ribbons.game_key
    LEFT JOIN ribbons on game_ribbons.ribbon_key = ribbons.ribbon_key
    LEFT JOIN ribbons_earned on (
      ribbons_earned.ribbon_key = ribbons.ribbon_key
      AND pokemon.ribbon_pokemon_id = ribbons_earned.ribbon_pokemon_id
    )
  ORDER BY pokemon.view_order ASC
) as ribbon_types
GROUP BY pokemon, pokedex_id, ribbon_types.achieved 
