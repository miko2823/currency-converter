# Currency Converter API

## summary
### Architecture
Simple and light sources, created for being a templates of golang development.
- Go 
  - chin for routing
  - github.com/jackc/pgconn for DB
- DDD patern

### API
- POST /login(just a demo)
- GET /converter
  - params: base(base currency)
  - params: symbols(exchange currency with conmma separated)
  - params: amount(amount of base currency to convert)

## develop settings
Everything are working inside Continer without any setup at local environments.
- At VSCode, search command "Open Folder in Continer..." and select the root directory.
- Your local source are mounted to the container.
- To start the server, run "go run ./cmd/" at VSCode terminal.

## production settings
- build an image from Docker file
