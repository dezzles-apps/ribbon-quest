(SELECT
  'RIBBON' as event_type,
  achieved_at as timestamp,
  pr.pokemon,
  p.nickname,
  pr.ribbon_key,
  r.name as ribbon_name,
  r.category as ribbon_category,
  r.type as ribbon_type
FROM pokemon_ribbons pr
LEFT JOIN pokemon p ON pr.pokemon = p.pokemon 
LEFT JOIN ribbons r ON pr.ribbon_key = r.ribbon_key )
UNION
(SELECT
  'RIBBON_CATCH' as event_type,
  caught_at as timestamp,
  pokemon,
  nickname,
  null as ribbon_key,
  null as ribbon_name,
  null as ribbon_category,
  null as ribbon_type
FROM pokemon
WHERE caught_at is not null)
ORDER BY timestamp DESC