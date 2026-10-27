Feature: Ribbon APIs

  Background:
    * url baseUrl

  @read
  Scenario: Get all ribbon games
    Given path 'api', 'ribbons', 'v1', 'games'
    When method get
    Then status 200
    And match response.data == '#[]'
    And match response.data[0] contains { gameKey: '#string', name: '#string' }

  @read
  Scenario: Get a ribbon game
    Given path 'api', 'ribbons', 'v1', 'games'
    When method get
    Then status 200
    * def gameKey = response.data[0].gameKey

    Given path 'api', 'ribbons', 'v1', 'games', gameKey
    When method get
    Then status 200
    And match response.data.gameKey == gameKey

  @read
  Scenario: Get all ribbon Pokemon
    Given path 'api', 'ribbons', 'v1', 'pokemon'
    When method get
    Then status 200
    And match response.data == '#[]'
    And match response.data[0].pokemon == '#string'

  @read
  Scenario: Get a ribbon Pokemon
    Given path 'api', 'ribbons', 'v1', 'pokemon'
    When method get
    Then status 200
    * def pokemonId = response.data[0].pokemon

    Given path 'api', 'ribbons', 'v1', 'pokemon', pokemonId
    When method get
    Then status 200
    And match response.data contains { pokemon: '#(pokemonId)', species: '#object', details: '#object', ribbons: '#[]', games: '#[]' }

  @read
  Scenario: Get ribbon statistics
    Given path 'api', 'ribbons', 'v1', 'stats'
    When method get
    Then status 200
    And match response.data == '#object'

  @write
  Scenario: Create a ribbon Pokemon
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    And match created.result.pokemonId == '#string'

  @write
  Scenario: Catch a ribbon Pokemon
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'catch'
    And header Authorization = created.result.authHeader
    When method post
    Then status 200
    And match response.data.details.caughtAt == '#string'

  @write
  Scenario: Update a ribbon Pokemon
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    * def nickname = 'Functional ' + karate.uuid()
    * def requestBody = { nickname: '#(nickname)', notes: 'Created by Karate functional tests', shiny: true }
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId
    And header Authorization = created.result.authHeader
    And request requestBody
    When method put
    Then status 200
    And match response.data.details.nickname == nickname
    And match response.data.details.shiny == true

  @write
  Scenario: Add a ribbon to a ribbon Pokemon
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId
    When method get
    Then status 200
    * def ribbonKey = response.data.ribbons[0].ribbonKey

    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'ribbons', ribbonKey
    And header Authorization = created.result.authHeader
    When method post
    Then status 200
    And match response.data contains { ribbonKey: '#(ribbonKey)', achieved: true }

  @write
  Scenario: Remove a ribbon from a ribbon Pokemon
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId
    When method get
    Then status 200
    * def ribbonKey = response.data.ribbons[0].ribbonKey

    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'ribbons', ribbonKey
    And header Authorization = created.result.authHeader
    When method post
    Then status 200

    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'ribbons', ribbonKey
    And header Authorization = created.result.authHeader
    When method delete
    Then status 200
    And match response.data contains { ribbonKey: '#(ribbonKey)', achieved: false }

  @write
  Scenario: Add a ribbon Pokemon to a game
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'games', created.result.secondaryGame
    And header Authorization = created.result.authHeader
    When method put
    Then status 200
    And match response.data.games[*].gameKey contains created.result.secondaryGame

  @write
  Scenario: Remove a ribbon Pokemon from a game
    * def created = call read('classpath:helpers/ribbon-pokemon.feature')
    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'games', created.result.secondaryGame
    And header Authorization = created.result.authHeader
    When method put
    Then status 200

    Given path 'api', 'ribbons', 'v1', 'pokemon', created.result.pokemonId, 'games', created.result.secondaryGame
    And header Authorization = created.result.authHeader
    When method delete
    Then status 200
    And match response.data.games[*].gameKey !contains created.result.secondaryGame
