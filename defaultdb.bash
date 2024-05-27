#!/bin/bash

docker exec -i goweb-mysql-1 sh -c 'exec mysql -uroot' < ./db.sql
