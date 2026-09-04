ALTER TABLE games ADD COLUMN ribbon_quest BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE games SET ribbon_quest = TRUE WHERE game_key IN ('Y', 'OMEGA_RUBY', 'ULTRA_SUN', 'SWORD', 'SCARLET', 'BRILLIANT_DIAMOND', 'LEGENDS_ARCEUS');

INSERT INTO games (game_key, name, view_order) VALUES
  ('X', 'X', 3),
  ('ALPHA_SAPPHIRE', 'Alpha Sapphire', 2),
  ('SUN', 'Sun', 5),
  ('MOON', 'Moon', 6),
  ('ULTRA_MOON', 'Ultra Moon', 8),
  ('LGE', 'Lets Go Eevee', 9),
  ('LGP', 'Lets Go Pikachu', 10),
  ('SWORD', 'Sword', 11),
  ('SHINING_PEARL', 'Shining Pearl', 15),
  ('VIOLET', 'Violet', 17),
  ('LZA', 'Legends: ZA', 18),
  ('GO', 'Pokemon GO', 19);

UPDATE games SET view_order = 1 WHERE game_key = 'OMEGA_RUBY';
UPDATE games SET view_order = 4 WHERE game_key = 'Y';
UPDATE games SET view_order = 7 WHERE game_key = 'ULTRA_SUN';
UPDATE games SET view_order = 12 WHERE game_key = 'SHIELD';
UPDATE games SET view_order = 13 WHERE game_key = 'LEGENDS_ARCEUS';
UPDATE games SET view_order = 14 WHERE game_key = 'BRILLIANT_DIAMOND';
UPDATE games SET view_order = 16 WHERE game_key = 'SCARLET';

ALTER TABLE pokemon ADD COLUMN notes TEXT NULL;