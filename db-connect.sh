#!/bin/sh

docker exec -it GE-postgres bash -c 'psql -h localhost -U postgres GE'
