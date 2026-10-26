Feature: Authentication APIs

  Background:
    * url baseUrl

  @write
  Scenario: Register a user
    * def credentials = { username: '#("functional-" + karate.uuid())', password: 'FunctionalTestPass!1' }
    Given path 'api', 'auth', 'v1', 'register'
    And request credentials
    When method post
    Then status 200
    And match response.message == 'user registered successfully'

  @write
  Scenario: Log in a registered user
    * def credentials = { username: '#("functional-" + karate.uuid())', password: 'FunctionalTestPass!1' }
    Given path 'api', 'auth', 'v1', 'register'
    And request credentials
    When method post
    Then status 200

    Given path 'api', 'auth', 'v1', 'login'
    And request credentials
    When method post
    Then status 200
    And match response.token == '#string'
