#!/usr/bin/env bash
source ./config.sh
migrate -database $DSN_URL -path $MIGRATION_DIR down
