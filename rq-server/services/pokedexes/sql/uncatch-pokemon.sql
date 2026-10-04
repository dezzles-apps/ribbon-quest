UPDATE pokedex_entries pe, pokedexes dexes
SET 
  pe.caught_at = NULL,
  pe.no_event_trigger = false 
WHERE
  dexes.name = ?
  AND dexes.pokedex_id = pe.pokedex_id
AND pe.pokedex_entry_id = ?