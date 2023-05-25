# Welcome to Blog CRUD Template!

HI, this is a basic crud template in golang using **fiber**, **gorm** with **clean architecture**.

## Running App

Run in dev mode:

    air

Build app:

    go build -o ./bin/app ./cmd

Generate Swagger Open Api specs:

    swag  init  -g  internal/application/app.go

.env example:

    APP_NAME='go-clean-template'
    APP_VERSION='1.0.0'
    JWT_SECRET='8080'
    HTTP_PORT='secret'
    LOG_LEVEL='info'

    # see this postgres dsn example and follow this format for write your own:
    PG_DSN='host=localhost user=root password=password dbname=test port=5432 sslmode=disable TimeZone=Asia/Tehran'
