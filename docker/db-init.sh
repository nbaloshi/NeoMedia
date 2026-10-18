#!/bin/bash
set -e

for f in /migrations/*.up.sql; do
    echo "Applying migration: $f"
    psql -v ON_ERROR_STOP=1 \
        --username "$POSTGRES_USER" \
        --dbname "$POSTGRES_DB" \
        -f "$f"
done

echo "Database migrations completed."