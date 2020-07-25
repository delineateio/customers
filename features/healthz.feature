Feature: Tests Healthz

Scenario: /healthz
  Given domain api with path /healthz
  When method GET
  Then status 200
