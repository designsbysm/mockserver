# Mock Server

HTTP server to mock API responses for testing.

## Sample ./server.yaml Config

```
api:
  address: :4000

gin:
  debug: false

timber:
  level: LevelAll
  flags: FlagColorful
  timestamp: "[15:04:05]"

```
