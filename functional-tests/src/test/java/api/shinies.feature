Feature: Shiny APIs

  Background:
    * url baseUrl

  @read
  Scenario: Get all shinies
    Given path 'api', 'shinies', 'v1'
    When method get
    Then status 200
    And match response.data == '#[]'

  @read
  Scenario: Get shiny statistics
    Given path 'api', 'shinies', 'v1', 'stats'
    And param startDate = '2020-01-01'
    And param endDate = '2100-01-01'
    When method get
    Then status 200
    And match response.data == '#object'

  @write
  Scenario: Create a shiny
    * def auth = call read('classpath:helpers/authentication.feature')
    Given path 'api', 'data', 'v1', 'pokemon'
    When method get
    Then status 200
    * def formId = response.data[0].forms[0].formId

    Given path 'api', 'data', 'v1', 'common'
    When method get
    Then status 200
    * def commonData = response.data
    * def requestBody = { form: '#(formId)', level: 1, game: '#(commonData.games[0].key)', method: '#(commonData.methods[0].key)', characteristic: '#(commonData.characteristics[0].key)', nature: '#(commonData.natures[0].key)', gender: '#(commonData.genders[0].key)', notes: 'Created by Karate functional tests' }

    Given path 'api', 'shinies', 'v1'
    And header Authorization = auth.result.authHeader
    And request requestBody
    When method post
    Then status 201
    And match response.data == null
