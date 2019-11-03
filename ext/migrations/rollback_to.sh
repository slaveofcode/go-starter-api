#!/usr/bin/env bash
source ./config.sh
echo -n "Please provide the number of how much migration to rollback [e.g. 1]: "
read digit
migrate -database $DSN_URL -path $MIGRATION_DIR down $digit
