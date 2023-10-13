#!/bin/sh

set -e

echo "run db migrations"
/app/migarte -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"