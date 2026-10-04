CREATE VIEW events_v4 as SELECT * FROM ((SELECT
  'RIBBONS' as event_category,
  'RIBBON' as event_type,
  achieved_at as timestamp,
  p.pokemon,
  p.nickname,
  pr.ribbon_key,
  r.name as ribbon_name,
  r.category as ribbon_category,
  r.type as ribbon_type,
  null as pokedex_name
FROM ribbons_earned pr
LEFT JOIN ribbon_pokemon p ON pr.ribbon_pokemon_id = p.ribbon_pokemon_id 
LEFT JOIN ribbons r ON pr.ribbon_key = r.ribbon_key )
UNION
(SELECT
  'RIBBONS' as event_category,
  'RIBBON_CATCH' as event_type,
  caught_at as timestamp,
  pokemon,
  nickname,
  null as ribbon_key,
  null as ribbon_name,
  null as ribbon_category,
  null as ribbon_type,
  null as pokedex_name
FROM ribbon_pokemon
WHERE caught_at is not null)
UNION
(SELECT
  'POKEDEX' as event_category,
  'POKEDEX_CATCH' as event_type,
  pe.caught_at as timestamp,
  pd.species as pokemon,
  null as nickname,
  null as ribbon_key,
  null as ribbon_name,
  null as ribbon_category,
  null as ribbon_type,
  pokedex.name as pokedex_name
FROM pokedex_entries pe
  LEFT JOIN pokemon_forms pf ON pf.form_id = pe.form_id
  LEFT JOIN pokemon_data pd ON pd.pokedex_id = pf.pokedex_id
  LEFT JOIN pokedexes pokedex ON pe.pokedex_id = pokedex.pokedex_id
  WHERE pe.caught_at IS NOT NULL
)
ORDER BY timestamp DESC) as events_v4;