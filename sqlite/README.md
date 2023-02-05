# SQLite


## Create Database

- create a new database from commend line
  ```bash
  # create new database "employees.db"
  sqlite3 employees.db

  # check all the usage commend
  .help

  # list names and files of attached databases
  .databases

  # create new table "employees"
  create table employees (
      employees_id INTEGER PRIMARY KEY,
      department VARCHAR(36),
      salary INTEGER
  );

  # list names of tables matching LIKE pattern TABLE
  .tables ?TABLE?

  # show "employees" schema
  .schema employees

  # insert data into table
  insert into employees (department, salary)
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

  # select all data
  select * from employees;

  # exit
  .exit
  ```

- create a new database from a sql script
  - move sql commends to a sql script and create a new database by running the
  script
  ```bash
  sqlite3 employees.db < init_employees.sql
  ```
