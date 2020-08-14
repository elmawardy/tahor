#!/usr/bin/env bash
echo "Initializing database"
psql -U postgres -c "CREATE DATABASE tahor"
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE tahor TO postgres"
echo "Database initialization succeeded"