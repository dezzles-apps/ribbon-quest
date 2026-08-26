# Ribbon Quest Server

## How to run

1. Configure `config.dev.yml` to connect to the database
2. Ensure that all SQL scripts in `sql` have been ran
3. Inside `rq-server` run `go get`
4. Inside `rq-server` run `ENVIRONMENT=dev go run .`


## Physical Data Model

```mermaid
---
config:
  layout: elk
---
erDiagram
  direction LR
  GAMES {
    string game_key
    int view_order
  }
  GAME_RIBBONS {
    string game_key
    string ribbon_key
  }
  POKEMON {
    string pokemon
    int view_order
  }
  POKEMON_GAMES {
    string pokemon
    string game_key
  }
  POKEMON_RIBBONS {
    string pokemon
    string ribbon_key
  }
  RIBBONS {
    string ribbon_key
    string category
    int view_order
  }
  GAMES ||--|{ GAME_RIBBONS : has
  GAME_RIBBONS }|--|| RIBBONS : has
  POKEMON_GAMES }|--|| GAMES : has
  POKEMON ||--|{ POKEMON_GAMES : has
  POKEMON }|--|| POKEMON_RIBBONS : has
  POKEMON_RIBBONS ||--|{ RIBBONS : has
```