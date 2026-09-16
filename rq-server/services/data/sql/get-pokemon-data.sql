SELECT pd.pokedex_id, pd.species, pf.form_name, pf.sprite_ref
FROM pokemon_data pd, pokemon_forms pf
WHERE pd.pokedex_id = pf.pokedex_id