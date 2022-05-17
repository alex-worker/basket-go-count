#!/bin/bash
set -e

./tcp-port-wait.sh "$DB_HOST" "$DB_PORT"
ls
go test -v ./testing
