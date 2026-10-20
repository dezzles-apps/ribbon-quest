SELECT
  shiny_id,
  pd.pokedex_id,
  pd.species,
  pf.form_name,
  pf.sprite_ref,
  s.caught_at,
  s.nature,
  s.characteristic,
  s.level,
  COALESCE(s.notes, '') AS notes,
  s.game,
  s.method,
  s.gender
FROM shinies s
LEFT JOIN pokemon_forms pf ON pf.form_id = s.form_id
LEFT JOIN pokemon_data pd ON pd.pokedex_id = pf.pokedex_id
ORDER BY caught_at DESC