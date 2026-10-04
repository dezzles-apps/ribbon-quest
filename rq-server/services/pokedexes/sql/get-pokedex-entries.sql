SELECT
  pe.pokedex_entry_id,
  pd.species,
  pf.sprite_ref,
  IF (pe.caught_at IS NOT NULL, TRUE, FALSE) AS caught
FROM pokedex_entries pe
  LEFT JOIN pokemon_forms pf ON pe.form_id = pf.form_id
  LEFT JOIN pokedexes p ON p.pokedex_id = pe.pokedex_id
  LEFT JOIN pokemon_data pd ON pd.pokedex_id  = pf.pokedex_id
WHERE p.name = ?
ORDER BY pe.pokedex_entry_id ASC