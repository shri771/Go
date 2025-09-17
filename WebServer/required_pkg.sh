#!/bin/bash

echo("Installing required pkgs")
echo("-----------------------------------------------------")
## Gosse
go install github.com/pressly/goose/v3/cmd/goose@latest

## SQLC
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## Google uuid pkg
go get github.com/google/uuid

## postgress drives
go get github.com/lib/pq
echo("-----------------------------------------------------")
echo("Installing pkgs completed")
