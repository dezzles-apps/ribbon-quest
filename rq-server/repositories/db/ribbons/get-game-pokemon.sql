SELECT
  pokemon.pokemon,
  pokemon.pokedex_id,
  species.species,
  pf.form_name,
  pf.sprite_ref,
  pokemon.nickname,
  pokemon.caught_at,
  pokemon.nature,
  pokemon.characteristic,
  pokemon.notes,
  pokemon.shiny
FROM ribbon_pokemon_games pg
LEFT JOIN ribbon_pokemon pokemon ON pokemon.ribbon_pokemon_id = pg.ribbon_pokemon_id
LEFT JOIN pokemon_data species ON species.pokedex_id = pokemon.pokedex_id
LEFT JOIN pokemon_forms pf ON pokemon.form_id = pf.form_id 
WHERE pg.game_key = ?
ORDER BY pokemon.view_order 