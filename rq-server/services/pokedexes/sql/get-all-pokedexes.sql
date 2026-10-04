
SELECT totals.name, catches.caught, totals.total, totals.image, totals.notes FROM (
  SELECT p.pokedex_id, p.name, p.image, p.notes, COUNT(*) as total FROM pokedex_entries pe, pokedexes p
  WHERE p.pokedex_id = pe.pokedex_id
  GROUP BY p.pokedex_id
) as totals LEFT JOIN (
  SELECT p.pokedex_id, p.name, COUNT(*) as caught FROM pokedex_entries pe, pokedexes p
  WHERE p.pokedex_id = pe.pokedex_id
  AND pe.caught_at IS NOT NULL
  GROUP BY p.pokedex_id
) as catches
ON totals.pokedex_id = catches.pokedex_id