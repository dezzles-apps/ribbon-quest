CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS config (
  config_key VARCHAR(255) PRIMARY KEY,
  value TEXT NOT NULL
);

INSERT INTO config (config_key, value) VALUES
  ('register_enabled', 'true');

CREATE TABLE IF NOT EXISTS pokemon (
    pokemon VARCHAR(255) PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    view_order int NOT NULL
);

CREATE TABLE IF NOT EXISTS games (
    game_key VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    view_order int NOT NULL
);

INSERT INTO pokemon (pokemon, nickname, region, view_order) VALUES
  ('Vulpix', 'Kiri', 'Kanto', 1),
  ('Bellossom', 'Bellossom', 'Johto', 2),
  ('Swablu', 'Swablu', 'Hoenn', 3),
  ('Leafeon', 'Leafeon', 'Sinnoh', 4),
  ('Lilligant', 'Lilligant', 'Unova', 5),
  ('Amaura', 'Amaura', 'Kalos', 6),
  ('Poipole', 'Poipole', 'Alola', 7),
  ('Ponyta', 'Ponyta', 'Galar', 8),
  ('Sprigatito', 'Mint', 'Paldea', 9),
  ('Overqwil', 'Carmilla', 'Hisui', 10),
  ('Meltan', 'Meltan', 'Unknown', 11);

INSERT INTO games (game_key, name, view_order) VALUES
  ('OMEGA_RUBY', 'Omega Ruby', 1),
  ('Y', 'Y', 2),
  ('ULTRA_SUN', 'Ultra Sun', 3),
  ('SHIELD', 'Shield', 4),
  ('LEGENDS_ARCEUS', 'Legends: Arceus', 5),
  ('BRILLIANT_DIAMOND', 'Brilliant Diamond', 6),
  ('SCARLET', 'Scarlet', 7);

CREATE TABLE IF NOT EXISTS pokemon_games (
    pokemon VARCHAR(255) REFERENCES pokemon(pokemon),
    game_key VARCHAR(255) REFERENCES games(game_key),
    PRIMARY KEY (pokemon, game_key)
);

INSERT INTO pokemon_games (pokemon, game_key) VALUES
  ('Vulpix', 'Y'),
  ('Vulpix', 'OMEGA_RUBY'),
  ('Vulpix', 'ULTRA_SUN'),
  ('Vulpix', 'SHIELD'),
  ('Vulpix', 'LEGENDS_ARCEUS'),
  ('Vulpix', 'BRILLIANT_DIAMOND'),
  ('Vulpix', 'SCARLET'),
  ('Bellossom', 'Y'),
  ('Bellossom', 'OMEGA_RUBY'),
  ('Bellossom', 'ULTRA_SUN'),
  ('Bellossom', 'SHIELD'),
  ('Bellossom', 'BRILLIANT_DIAMOND'),
  ('Bellossom', 'SCARLET'),
  ('Swablu', 'Y'),
  ('Swablu', 'OMEGA_RUBY'),
  ('Swablu', 'ULTRA_SUN'),
  ('Swablu', 'SHIELD'),
  ('Swablu', 'BRILLIANT_DIAMOND'),
  ('Swablu', 'SCARLET'),
  ('Leafeon', 'Y'),
  ('Leafeon', 'OMEGA_RUBY'),
  ('Leafeon', 'ULTRA_SUN'),
  ('Leafeon', 'SHIELD'),
  ('Leafeon', 'LEGENDS_ARCEUS'),
  ('Leafeon', 'BRILLIANT_DIAMOND'),
  ('Leafeon', 'SCARLET'),
  ('Lilligant', 'Y'),
  ('Lilligant', 'OMEGA_RUBY'),
  ('Lilligant', 'ULTRA_SUN'),
  ('Lilligant', 'SHIELD'),
  ('Lilligant', 'LEGENDS_ARCEUS'),
  ('Lilligant', 'SCARLET'),
  ('Amaura', 'Y'),
  ('Amaura', 'OMEGA_RUBY'),
  ('Amaura', 'ULTRA_SUN'),
  ('Amaura', 'SHIELD'),
  ('Poipole', 'ULTRA_SUN'),
  ('Poipole', 'SHIELD'),
  ('Ponyta', 'SHIELD'),
  ('Sprigatito', 'SCARLET'),
  ('Overqwil', 'LEGENDS_ARCEUS'),
  ('Overqwil', 'SCARLET'),
  ('Meltan', 'SHIELD');

CREATE TABLE IF NOT EXISTS ribbons (
  ribbon_key VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  enabled BOOLEAN NOT NULL DEFAULT TRUE,
  type VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  view_order int NOT NULL
);

INSERT INTO ribbons (ribbon_key, name, enabled, type, category, view_order) VALUES
('SINNOH_CHAMPION', 'Sinnoh Champion', 1,'RIBBON','CHAMPION',5),
('KALOS_CHAMPION', 'Kalos Champion', 1,'RIBBON','CHAMPION',2),
('HOENN_CHAMPION', 'Hoenn Champion', 1,'RIBBON','CHAMPION',1),
('ALOLA_CHAMPION', 'Alola Champion', 1,'RIBBON','CHAMPION',3),
('TOWER_MASTER', 'Tower Master', 0,'RIBBON','BATTLE',1),
('MASTER_RANK', 'Master Rank', 0,'RIBBON','BATTLE',2),
('GALAR_CHAMPION', 'Galar Champion', 1,'RIBBON','CHAMPION',4),
('PALDEA_CHAMPION', 'Paldea Champion', 1,'RIBBON','CHAMPION',6),
('COOLNESS_MASTER', 'Coolness Master', 1,'RIBBON','CONTEST',3),
('BEAUTY_MASTER', 'Beauty Master', 1,'RIBBON','CONTEST',1),
('CUTENESS_MASTER', 'Cuteness Master', 1,'RIBBON','CONTEST',4),
('CLEVERNESS_MASTER', 'Cleverness Master', 1,'RIBBON','CONTEST',2),
('TOUGHNESS_MASTER', 'Toughness Master', 1,'RIBBON','CONTEST',5),
('CONTEST_STAR', 'Contest Star', 1,'RIBBON','CONTEST',6),
('TWINKLING_STAR', 'Twinkling Star', 1,'RIBBON','CONTEST',7),
('BATTLE_ROYAL_MASTER', 'Battle Royal Master', 1,'RIBBON','BATTLE',3),
('EFFORT', 'Effort', 1,'RIBBON','STATS',1),
('ALERT', 'Alert', 1,'RIBBON','JULIE',1),
('SHOCK', 'Shock', 1,'RIBBON','JULIE',2),
('DOWNCAST', 'Downcast', 1,'RIBBON','JULIE',3),
('CARELESS', 'Careless', 1,'RIBBON','JULIE',4),
('RELAX', 'Relax', 1,'RIBBON','JULIE',5),
('SNOOZE', 'Snooze', 1,'RIBBON','JULIE',6),
('SMILE', 'Smile', 1,'RIBBON','JULIE',7),
('GORGEOUS', 'Gorgeous', 1,'RIBBON','SHOPPING',1),
('ROYAL', 'Royal', 1,'RIBBON','SHOPPING',2),
('GORGEOUS_ROYAL', 'Gorgeous Royal', 1,'RIBBON','SHOPPING',3),
('FOOTPRINT', 'Footprint', 1,'RIBBON','STATS',2),
('BEST_FRIENDS', 'Best Friends', 1,'RIBBON','STATS',3),
('TRAINING', 'Training', 1,'RIBBON','STATS',4),
('HISUI', 'Hisui', 1,'RIBBON','SHOPPING',4),
('DESTINY', 'Destiny', 1,'MARK','STATS',5),
('ITEMFINDER', 'Itemfinder', 1,'MARK','STATS',6),
('GOURMAND', 'Gourmand', 1,'MARK','STATS',7),
('PARTNER', 'Partner', 1,'MARK','STATS',8);

CREATE TABLE IF NOT EXISTS game_ribbons (
  game_key VARCHAR(255) REFERENCES games(game_key),
  ribbon_key VARCHAR(255) REFERENCES ribbons(ribbon_key),
  PRIMARY KEY (game_key, ribbon_key)
);

INSERT INTO game_ribbons (ribbon_key, game_key) VALUES
('SINNOH_CHAMPION', 'BRILLIANT_DIAMOND'),
('KALOS_CHAMPION', 'Y'),
('HOENN_CHAMPION', 'OMEGA_RUBY'),
('ALOLA_CHAMPION', 'ULTRA_SUN'),
('GALAR_CHAMPION', 'SHIELD'),
('PALDEA_CHAMPION', 'SCARLET'),
('COOLNESS_MASTER', 'OMEGA_RUBY'),
('COOLNESS_MASTER', 'BRILLIANT_DIAMOND'),
('BEAUTY_MASTER', 'OMEGA_RUBY'),
('BEAUTY_MASTER', 'BRILLIANT_DIAMOND'),
('CUTENESS_MASTER', 'OMEGA_RUBY'),
('CUTENESS_MASTER', 'BRILLIANT_DIAMOND'),
('CLEVERNESS_MASTER', 'OMEGA_RUBY'),
('CLEVERNESS_MASTER', 'BRILLIANT_DIAMOND'),
('TOUGHNESS_MASTER', 'OMEGA_RUBY'),
('TOUGHNESS_MASTER', 'BRILLIANT_DIAMOND'),
('CONTEST_STAR', 'OMEGA_RUBY'),
('CONTEST_STAR', 'BRILLIANT_DIAMOND'),
('TWINKLING_STAR', 'BRILLIANT_DIAMOND'),
('BATTLE_ROYAL_MASTER', 'ULTRA_SUN'),
('TOWER_MASTER', 'SHIELD'),
('TOWER_MASTER', 'BRILLIANT_DIAMOND'),
('EFFORT', 'Y'),
('EFFORT', 'OMEGA_RUBY'),
('EFFORT', 'ULTRA_SUN'),
('EFFORT', 'SHIELD'),
('EFFORT', 'BRILLIANT_DIAMOND'),
('EFFORT', 'SCARLET'),
('ALERT', 'OMEGA_RUBY'),
('ALERT', 'BRILLIANT_DIAMOND'),
('SHOCK', 'OMEGA_RUBY'),
('SHOCK', 'BRILLIANT_DIAMOND'),
('DOWNCAST', 'OMEGA_RUBY'),
('DOWNCAST', 'BRILLIANT_DIAMOND'),
('CARELESS', 'OMEGA_RUBY'),
('CARELESS', 'BRILLIANT_DIAMOND'),
('RELAX', 'OMEGA_RUBY'),
('RELAX', 'BRILLIANT_DIAMOND'),
('SNOOZE', 'OMEGA_RUBY'),
('SNOOZE', 'BRILLIANT_DIAMOND'),
('SMILE', 'OMEGA_RUBY'),
('SMILE', 'BRILLIANT_DIAMOND'),
('GORGEOUS', 'OMEGA_RUBY'),
('GORGEOUS', 'BRILLIANT_DIAMOND'),
('ROYAL', 'OMEGA_RUBY'),
('ROYAL', 'BRILLIANT_DIAMOND'),
('GORGEOUS_ROYAL', 'OMEGA_RUBY'),
('GORGEOUS_ROYAL', 'BRILLIANT_DIAMOND'),
('FOOTPRINT', 'Y'),
('FOOTPRINT', 'OMEGA_RUBY'),
('FOOTPRINT', 'ULTRA_SUN'),
('FOOTPRINT', 'BRILLIANT_DIAMOND'),
('BEST_FRIENDS', 'Y'),
('BEST_FRIENDS', 'OMEGA_RUBY'),
('BEST_FRIENDS', 'ULTRA_SUN'),
('BEST_FRIENDS', 'SHIELD'),
('BEST_FRIENDS', 'BRILLIANT_DIAMOND'),
('BEST_FRIENDS', 'SCARLET'),
('TRAINING', 'Y'),
('TRAINING', 'OMEGA_RUBY'),
('MASTER_RANK', 'BRILLIANT_DIAMOND'),
('MASTER_RANK', 'SCARLET'),
('HISUI', 'LEGENDS_ARCEUS'),
('DESTINY', 'SCARLET'),
('ITEMFINDER', 'SCARLET'),
('GOURMAND', 'SCARLET'),
('PARTNER', 'SCARLET');

CREATE TABLE IF NOT EXISTS pokemon_ribbons (
  pokemon VARCHAR(255) REFERENCES pokemon(pokemon),
  ribbon_key VARCHAR(255) REFERENCES ribbons(ribbon_key),
  PRIMARY KEY (pokemon, ribbon_key)
);

ALTER TABLE pokemon_ribbons ADD COLUMN achieved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE pokemon ADD COLUMN caught_at TIMESTAMP NULL;
ALTER TABLE pokemon ADD COLUMN shiny BOOL DEFAULT FALSE;
ALTER TABLE pokemon ADD COLUMN nature VARCHAR(255) NULL;
ALTER TABLE pokemon ADD COLUMN characteristic VARCHAR(255) NULL;