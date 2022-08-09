#!/bin/bash

# create new database "employee.db"
sqlite3 employee.db < init_employee.sql

# select all from the new created "employee.db"
sqlite3 employee.db <<EOF
    select * from employee;
EOF
