#!/bin/sh

set -o errexit
set -o pipefail
set -o nounset

# Run migrations
echo "Running migrations..."
migrate -path ./migrations -database $DATABASE_URL up

# Run server
echo "Running server..."
./server