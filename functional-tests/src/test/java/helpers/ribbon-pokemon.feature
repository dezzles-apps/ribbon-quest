Feature: Create a ribbon Pokemon for a functional test

  Scenario:
    * url baseUrl
    * def auth = call read('classpath:helpers/authentication.feature')

    Given path 'api', 'data', 'v1', 'pokemon'
    When method get
    Then status 200
    * def pokemonData = response.data[0]
    * def form = pokemonData.forms[0]

    Given path 'api', 'ribbons', 'v1', 'games'
    When method get
    Then status 200
    * def games = response.data
    * match games[1] == '#object'
    * def primaryGame = games[0].gameKey
    * def secondaryGame = games[1].gameKey
    * def pokemonId = 'functional-' + karate.uuid()
    * def requestBody = { pokemon: '#(pokemonId)', pokedexId: '#(pokemonData.pokedexNo.number)', form: '#(form.form)', games: ['#(primaryGame)'] }

    Given path 'api', 'ribbons', 'v1', 'pokemon'
    And header Authorization = auth.result.authHeader
    And request requestBody
    When method post
    Then status 200
    * def result = { authHeader: '#(auth.result.authHeader)', pokemonId: '#(pokemonId)', primaryGame: '#(primaryGame)', secondaryGame: '#(secondaryGame)' }
