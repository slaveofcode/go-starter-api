#!/usr/bin/env bash
source ./config.sh
echo -n "Please provide the migration file name [without space chars]: "
read migration_file
migrate create -ext sql -dir $MIGRATION_DIR -seq $migration_file