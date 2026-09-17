ALTER TABLE pokemon ADD COLUMN ribbon_pokemon_id INT AUTO_INCREMENT UNIQUE FIRST;

ALTER TABLE pokemon DROP PRIMARY KEY;
ALTER TABLE pokemon ADD PRIMARY KEY (ribbon_pokemon_id);

RENAME TABLE pokemon TO ribbon_pokemon;
RENAME TABLE pokemon_ribbons TO ribbons_earned;
RENAME TABLE pokemon_games TO ribbon_pokemon_games;

ALTER TABLE ribbon_pokemon_games ADD COLUMN ribbon_pokemon_id INT NULL;

UPDATE ribbon_pokemon_games rpg 
JOIN ribbon_pokemon rp ON rpg.pokemon = rp.pokemon
SET rpg.ribbon_pokemon_id = rp.ribbon_pokemon_id;

ALTER TABLE ribbon_pokemon_games DROP PRIMARY KEY;
ALTER TABLE ribbon_pokemon_games ADD PRIMARY KEY(ribbon_pokemon_id, game_key);
ALTER TABLE ribbon_pokemon_games DROP COLUMN pokemon;
ALTER TABLE ribbon_pokemon_games ADD CONSTRAINT fk_ribbon_pokemon_id FOREIGN KEY (ribbon_pokemon_id) REFERENCES ribbon_pokemon(ribbon_pokemon_id);
ALTER TABLE ribbon_pokemon_games ADD CONSTRAINT fk_game_key FOREIGN KEY (game_key) REFERENCES games(game_key);

-- Update ribbons_earned
ALTER TABLE ribbons_earned ADD COLUMN ribbon_pokemon_id INT NULL;
UPDATE ribbons_earned rpg 
JOIN ribbon_pokemon rp ON rpg.pokemon = rp.pokemon
SET rpg.ribbon_pokemon_id = rp.ribbon_pokemon_id;
ALTER TABLE ribbons_earned DROP PRIMARY KEY;
ALTER TABLE ribbons_earned ADD PRIMARY KEY(ribbon_pokemon_id, ribbon_key);
ALTER TABLE ribbons_earned DROP COLUMN pokemon;

ALTER TABLE ribbons_earned ADD CONSTRAINT fk_earned_ribbon_pokemon_id FOREIGN KEY (ribbon_pokemon_id) REFERENCES ribbon_pokemon(ribbon_pokemon_id);
ALTER TABLE ribbons_earned ADD CONSTRAINT fk_earned_ribbon_key FOREIGN KEY (ribbon_key) REFERENCES ribbons(ribbon_key);

-- New events view
CREATE VIEW events_v2 as SELECT * FROM ((SELECT
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
ORDER BY timestamp DESC) as events_v2;

ALTER TABLE ribbon_pokemon
MODIFY COLUMN form_id bigint unsigned NOT NULL;

ALTER TABLE ribbon_pokemon
ADD CONSTRAINT fk_ribbon_pokemon_form
FOREIGN KEY (form_id) REFERENCES pokemon_forms(form_id);

ALTER TABLE pokemon_forms
ADD CONSTRAINT fk_pokemon_form_pokemon
FOREIGN KEY (pokedex_id) REFERENCES pokemon_data(pokedex_id);

ALTER TABLE ribbon_pokemon
ADD CONSTRAINT fk_ribbon_pokemon_pokemon
FOREIGN KEY (pokedex_id) REFERENCES pokemon_data(pokedex_id);

ALTER TABLE game_ribbons
ADD CONSTRAINT fk_game_ribbon_game
FOREIGN KEY (game_key) REFERENCES games(game_key);

ALTER TABLE game_ribbons 
ADD CONSTRAINT fk_game_ribbons_ribbon
FOREIGN KEY (ribbon_key) REFERENCES ribbons(ribbon_key);