@read
Feature: Event APIs

  Background:
    * url baseUrl

  Scenario: Get events
    Given path 'api', 'events', 'v1'
    When method get
    Then status 200
    And match response.data contains { items: '#[]' }
