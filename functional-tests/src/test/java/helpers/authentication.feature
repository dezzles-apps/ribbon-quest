Feature: Authenticate a functional test user

  Scenario:
    * url baseUrl
    * def username = 'functional-' + karate.uuid()
    * def credentials = { username: '#(username)', password: 'FunctionalTestPass!1' }

    Given path 'api', 'auth', 'v1', 'register'
    And request credentials
    When method post
    Then status 200

    Given path 'api', 'auth', 'v1', 'login'
    And request credentials
    When method post
    Then status 200
    * def result = { authHeader: '#("Bearer " + response.token)' }
