#!/bin/bash

# Install some common packages
go get -u github.com/pkg/errors
go get -u gopkg.in/yaml.v2
go get -u github.com/spf13/afero
go get -u github.com/sirupsen/logrus
go get -u github.com/davecgh/go-spew/spew
go get -u github.com/aws/aws-lambda-go/events
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
go get -u gopkg.in/gormigrate.v1
go get -u github.com/joho/godotenv/autoload
go get -u github.com/alecthomas/template
go get -u github.com/lib/pq
go get -u github.com/pressly/goose/cmd/goose
go get -u github.com/vektra/mockery/v2/.../