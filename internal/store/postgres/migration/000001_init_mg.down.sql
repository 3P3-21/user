
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';

GRANT SELECT, INSERT, UPDATE, DELETE ON database_name.* TO 'username'@'localhost';
