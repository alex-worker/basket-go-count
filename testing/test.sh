#!/bin/bash
set -e

function main() {
  DB_HOST=$1
  DB_PORT=$2
  # shellcheck disable=SC2086
  ./tcp-port-wait.sh $DB_HOST "$DB_PORT"
}

echo "$DB_HOST" "$DB_PORT"
main "$DB_HOST" "$DB_PORT"
