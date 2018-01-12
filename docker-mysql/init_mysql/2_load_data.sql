USE employees;
SELECT 'LOADING departments' as 'INFO';
source /test_db/load_departments.dump;

SELECT 'LOADING employees' as 'INFO';
source /test_db/load_employees.dump;
SELECT 'LOADING dept_emp' as 'INFO';
source /test_db/load_dept_emp.dump;

SELECT 'LOADING dept_manager' as 'INFO';
source /test_db/load_dept_manager.dump;

SELECT 'LOADING titles' as 'INFO';
source /test_db/load_titles.dump;

SELECT 'LOADING salaries' as 'INFO';
source /test_db/load_salaries1.dump;
source /test_db/load_salaries2.dump;
source /test_db/load_salaries3.dump;

source /test_db/show_elapsed.sql ;
