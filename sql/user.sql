use public;
show tables;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

desc users;
-- +-------+--------------+------+-----+---------+----------------+
-- | Field | Type         | Null | Key | Default | Extra          |
-- +-------+--------------+------+-----+---------+----------------+
-- | id    | int          | NO   | PRI | NULL    | auto_increment |
-- | name  | varchar(255) | NO   |     | NULL    |                |
-- | email | varchar(255) | NO   | UNI | NULL    |                |
-- +-------+--------------+------+-----+---------+----------------+
-- 3 rows in set (0.00 sec)

INSERT INTO users(name, email) VALUES ('John Doe', 'johndoe@example.com');
-- Query OK, 1 row affected (0.02 sec)

select * from users;
-- +----+----------+---------------------+
-- | id | name     | email               |
-- +----+----------+---------------------+
-- |  1 | John Doe | johndoe@example.com |
-- +----+----------+---------------------+
-- 1 row in set (0.00 sec)


