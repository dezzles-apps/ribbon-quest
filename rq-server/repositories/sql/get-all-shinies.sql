SELECT
  shiny_id,
  pd.pokedex_id,
  pd.species,
  pf.form_name,
  pf.sprite_ref,
  s.catch_time,
  s.nature,
  s.characteristic,
  s.level,
  s.notes,
  s.game
FROM shinies s
LEFT JOIN pokemon_forms pf ON pf.form_id = s.form_id
LEFT JOIN pokemon_data pd ON pd.pokedex_id = pf.pokedex_id