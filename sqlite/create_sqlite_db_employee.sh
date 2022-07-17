#!/bin/bash

# create new database "employee.db"
sqlite3 employee.db <<EOF

create table employee (
    employee_id INTEGER PRIMARY KEY,
    department VARCHAR(36),
    salary INTEGER
);

insert into employee (department, salary)
values
    ('A', '10'),
    ('A', '20'),
    ('A', '35'),
    ('A', '42'),
    ('B', '12'),
    ('B', '21'),
    ('B', '33'),
    ('C', '14'),
    ('C', '23'),
    ('C', '38'),
    ('C', '41'),
    ('C', '55');

select * from employee;

EOF
