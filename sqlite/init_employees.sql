CREATE TABLE employees (
    employee_id INTEGER PRIMARY KEY,
    department VARCHAR(36),
    salary INTEGER
);

INSERT INTO employees (department, salary)
VALUES
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
