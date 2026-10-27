Feature: Reference data APIs

  Background:
    * url baseUrl

  @read
  Scenario: Get Pokemon reference data
    Given path 'api', 'data', 'v1', 'pokemon'
    When method get
    Then status 200
    And match response.data == '#[]'
    And match response.data[0] contains { pokedexNo: '#object', species: '#string', forms: '#[]' }

  @read
  Scenario: Get common reference data
    Given path 'api', 'data', 'v1', 'common'
    When method get
    Then status 200
    And match response.data contains { games: '#[]', genders: '#[]', characteristics: '#[]', natures: '#[]', methods: '#[]' }
