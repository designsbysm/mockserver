# Go Server

HTTP server prototype written in Go.

## Sample ./server.yaml Config

// TODO: update config example

```
server:
  port: :5000
  protocol: HTTP
  https:
    cert: cert.pem
    key: key.pem

db:
  connection: postgres://<username>:<passowrd>@<host>:<port>/<database>?sslmode=disable

gin:
  release: true

gorm:
  level: 3

timber:
  cli:
    level: 9
    timestamp: "[15:04:05]"
  file:
    level: 9
    timestamp: "[15:04:05]"
    path: logs/2006-01-02.log
```
