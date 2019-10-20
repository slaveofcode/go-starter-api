#!/usr/bin/env bash
DSN_URL="postgres://slaveofcode:slaveofcode@localhost:5432/go_starter?sslmode=disable"
MIGRATION_DIR="queries"

if ! [ -x "$(command -v migrate)" ]; then
  echo 'Error: migrate CLI is not installed.' >&2
  echo 'Please refer this source to install: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation'
  exit 1
fi