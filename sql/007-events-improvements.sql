-- New events view
CREATE VIEW events_v3 as SELECT * FROM ((SELECT
  'RIBBONS' as event_category,
  'RIBBON' as event_type,
  achieved_at as timestamp,
  p.pokemon,
  p.nickname,
  pr.ribbon_key,
  r.name as ribbon_name,
  r.category as ribbon_category,
  r.type as ribbon_type
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
  null as ribbon_type
FROM ribbon_pokemon
WHERE caught_at is not null)
ORDER BY timestamp DESC) as events_v3;