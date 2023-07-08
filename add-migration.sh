#!/bin/bash

[ ! -d "$MIGRATIONS_PATH" ] && mkdir $MIGRATIONS_PATH
[ -z "$MIGRATION_NAME" ] && echo "migration name is not set, please use MIGRATION_NAME" && exit 1

touch $MIGRATIONS_PATH/$(date +%s)_$MIGRATION_NAME.up.sql
touch $MIGRATIONS_PATH/$(date +%s)_$MIGRATION_NAME.down.sql
