Feature: Tests Customer

Scenario: /customer
  Given domain api with path /customer
    And header Content-Type as application/json
    And request {"forename": "jonathan","surname": "fenwick"}
  When method POST
  Then status 201
