UPDATE pokedex_entries pe, pokedexes dexes
SET 
  pe.caught_at = CURRENT_TIMESTAMP,
  pe.no_event_trigger = false 
WHERE
  dexes.name = ?
  AND dexes.pokedex_id = pe.pokedex_id
  AND pe.pokedex_entry_id = ?
  AND pe.caught_at IS NULL