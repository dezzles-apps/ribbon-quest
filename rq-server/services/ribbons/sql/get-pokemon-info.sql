SELECT p.pokemon, data.pokedex_id, data.species, form.form_name, form.sprite_ref, p.nickname, p.caught_at, p.nature, p.characteristic, p.notes, p.shiny
FROM pokemon p, pokemon_data data, pokemon_forms form
WHERE p.pokemon = ?
AND p.pokedex_id = data.pokedex_id 
AND form.form_id = p.form_id
