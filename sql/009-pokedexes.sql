CREATE TABLE IF NOT EXISTS pokedexes (
  pokedex_id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  prefix VARCHAR(8) NOT NULL,
  padding int NOT NULL DEFAULT 3
);

CREATE TABLE IF NOT EXISTS pokedex_entries (
  pokedex_entry_id INT NOT NULL,
  pokedex_id BIGINT UNSIGNED NOT NULL REFERENCES pokedexes(pokedex_id),
  caught_at TIMESTAMP NULL DEFAULT NULL,
  form_id BIGINT UNSIGNED NOT NULL REFERENCES pokemon_forms(form_id),
  PRIMARY KEY (pokedex_entry_id, pokedex_id)
);

INSERT INTO pokedexes (pokedex_id, name, prefix, padding) VALUES
  (1, 'National', 'NAT', 4),
  (2, 'Shiny', 'SHI', 4);