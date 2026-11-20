-- change test_db dabase --
\c test
 
-- public schema create --
GRANT USAGE, CREATE ON SCHEMA public TO test_db;
