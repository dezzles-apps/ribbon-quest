Feature: Pokedex APIs

  Background:
    * url baseUrl

  @read
  Scenario: Get all Pokedexes
    Given path 'api', 'pokedexes', 'v1'
    When method get
    Then status 200
    And match response.data == '#[]'
    And match response.data[0] contains { name: '#string', stats: '#object' }

  @read
  Scenario: Get a Pokedex
    Given path 'api', 'pokedexes', 'v1'
    When method get
    Then status 200
    * def pokedex = response.data[0].name

    Given path 'api', 'pokedexes', 'v1', pokedex
    When method get
    Then status 200
    And match response.data.name == pokedex
    And match response.data.entries == '#[]'

  @write
  Scenario: Catch a Pokedex entry
    * def auth = call read('classpath:helpers/authentication.feature')
    Given path 'api', 'pokedexes', 'v1'
    When method get
    Then status 200
    * def pokedex = response.data[0].name

    Given path 'api', 'pokedexes', 'v1', pokedex
    When method get
    Then status 200
    * def pokedexNo = response.data.entries[0].pokedexNo.text

    Given path 'api', 'pokedexes', 'v1', pokedex, 'entries', pokedexNo
    And header Authorization = auth.result.authHeader
    When method post
    Then status 200
    And match response.data contains { pokedexNo: '#object', caught: true }

  @write
  Scenario: Uncatch a Pokedex entry
    * def auth = call read('classpath:helpers/authentication.feature')
    Given path 'api', 'pokedexes', 'v1'
    When method get
    Then status 200
    * def pokedex = response.data[0].name

    Given path 'api', 'pokedexes', 'v1', pokedex
    When method get
    Then status 200
    * def pokedexNo = response.data.entries[0].pokedexNo.text

    Given path 'api', 'pokedexes', 'v1', pokedex, 'entries', pokedexNo
    And header Authorization = auth.result.authHeader
    When method post
    Then status 200

    Given path 'api', 'pokedexes', 'v1', pokedex, 'entries', pokedexNo
    And header Authorization = auth.result.authHeader
    When method delete
    Then status 200
    And match response.data contains { pokedexNo: '#object', caught: false }
