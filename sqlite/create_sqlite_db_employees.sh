#!/bin/bash

# create new database "employees.db"
sqlite3 employees.db < init_employees.sql

# select all from the new created "employees.db"
sqlite3 employees.db <<EOF
    select * from employees;
EOF
