# Currency Converter API

## summary
Using converter api from ApiLayer
- https://apilayer.com/marketplace/exchangerates_data-api

### Architecture
Simple and light sources, created for being a templates of golang development.
- Go
  - chin for routing
  - github.com/jackc/pgconn for DB
- DDD desing patern

### API
- POST /login(just a demo)
- GET /converter
  - params: base(base currency)
  - params: symbols(exchange currency with conmma separated)
  - params: amount(amount of base currency to convert)

## develop settings
Set up env file
```
mkdir .env
touch .env/dev.json
```
set your environment variables like below
```json
{
  "DB_HOST": "postgres",
  "PORT": 5432,
  "DB_USER": "postgres",
  "DB_PASSWORD": "postgres",
  "DB_NAME": "users",
  "CONVERTER_API_KEY": "{USER API KEY FOR ApiLayer}",
  "TOKEN_SIGNING_KEY": "test",
  "TOKEN_EXPIRATION": 100
}
```

Everything are working inside Continer without any setup at local environments.
- At VSCode, search command "Open Folder in Continer..." and select the root directory.
- Your local source are mounted to the container and ready to accept request.

## production settings
- build an image from Docker file
