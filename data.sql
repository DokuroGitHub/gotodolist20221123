-- USE test_partition;

-- todo_items ---------------------------------------------------------------------
DROP TABLE `todo_items`;

CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(150) CHARACTER SET utf8 NOT NULL,
  `status` enum('Doing','Finished') DEFAULT 'Doing',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- wallets ---------------------------------------------------------------------
DROP TABLE `wallets`;

CREATE TABLE `wallets` (
	`id` int not null AUTO_INCREMENT, primary key (`id`, `user_id`), -- user_id để partition
    `user_id` int unsigned NOT NULL,
    `bank_id` int unsigned NOT NULL,
    `account_id` int unsigned NOT NULL,
    `amount` int NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
	UNIQUE multiple_index (user_id, bank_id, account_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1
PARTITION BY HASH (user_id)
PARTITIONS 50;

INSERT INTO `wallets`(`user_id`,`bank_id`, `account_id`,`amount`)
VALUES
(0, 0, 0, 100),
(1, 1, 1, 101),
(1, 1, 2, 103),
(1, 2, 3, 104),
(2, 1, 4, 1006),
(2, 5, 5, 1003),
(3, 4, 6, 1040),
(4, 2, 7, 1050),
(5, 0, 8, 1070),
(5, 7, 9, 1030),
(6, 3, 10, 1200),
(7, 5, 11, 1020),
(7, 8, 12, 100),
(8, 3, 13, 1030),
(9, 9, 14, 1080),
(10, 4, 15, 1900),
(10, 4, 16, 9100),
(11, 3, 17, 1300),
(11, 9, 18, 1005),
(11, 2, 19, 1040),
(12, 4, 20, 1002),
(13, 1, 21, 1008),
(14, 5, 22, 1500),
(15, 1, 23, 7100),
(15, 6, 24, 1400),
(16, 2, 25, 4100);

-- members ---------------------------------------------------------------------

DROP TABLE `members`;

CREATE TABLE `members` (
    `username` varchar(16) NOT NULL,
    `firstname` varchar(25) NOT NULL,
    `lastname` varchar(25) NOT NULL,
    `email` varchar(35),
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL
)
PARTITION BY RANGE( UNIX_TIMESTAMP(created_at) ) (
    PARTITION p0 VALUES LESS THAN ( UNIX_TIMESTAMP('2000/01/01') ),
    PARTITION p1 VALUES LESS THAN ( UNIX_TIMESTAMP('2010/01/01') ),
    PARTITION p2 VALUES LESS THAN ( UNIX_TIMESTAMP('2020/01/01') ),
    PARTITION p3 VALUES LESS THAN ( UNIX_TIMESTAMP('2030/01/01') ),
    PARTITION p4 VALUES LESS THAN (MAXVALUE)
);

INSERT INTO members(firstname, lastname, username, email, created_at)
VALUES
("firstname0", "lastname0", "username0","email0", "2022/01/01"),
("firstname1", "lastname1", "username1","email1", "2000/01/01"),
("firstname2", "lastname2", "username2","email2", "1999/01/01"),
("firstname3", "lastname3", "username3","email3", "1978/01/01"),
("firstname4", "lastname4", "username4","email4", "2000/01/01"),
("firstname5", "lastname5", "username5","email5", "1977/01/01"),
("firstname6", "lastname6", "username6","email6", "2009/01/01"),
("firstname7", "lastname7", "username7","email7", "2025/01/01"),
("firstname8", "lastname8", "username8","email8", "1999/01/01"),
("firstname9", "lastname9", "username9","email9", "1989/01/01");

ALTER TABLE members DROP PARTITION p4;
ALTER TABLE members ADD PARTITION (
    PARTITION p4 VALUES LESS THAN ( UNIX_TIMESTAMP('2040/01/01') ),
    PARTITION p5 VALUES LESS THAN (MAXVALUE)
);

INSERT INTO members(firstname, lastname, username, email, created_at)
VALUES
("firstname10", "lastname10", "username10","email10", "2022/01/01"),
("firstname11", "lastname11", "username11","email11", "2030/01/01"),
("firstname12", "lastname12", "username12","email12", "1989/01/01"),
("firstname13", "lastname13", "username13","email13", "2033/01/01"),
("firstname14", "lastname14", "username14","email14", "2022/01/01");

-- employees -------------------------------------------------------------------
DROP TABLE employees;

CREATE TABLE employees (
    id INT NOT NULL,
    fname VARCHAR(30),
    lname VARCHAR(30),
    store_id INT
)
PARTITION BY LIST(store_id) (
    PARTITION pNorth VALUES IN (3,5,6,9,17),
    PARTITION pEast VALUES IN (1,2,10,11,19,20),
    PARTITION pWest VALUES IN (4,12,13,14,18),
    PARTITION pCentral VALUES IN (7,8,15,16)
);

----------------------------------------------------------------------

 SELECT PARTITION_NAME,TABLE_ROWS
      FROM INFORMATION_SCHEMA.PARTITIONS
      WHERE TABLE_NAME = 'wallets';

 SELECT PARTITION_NAME,TABLE_ROWS
      FROM INFORMATION_SCHEMA.PARTITIONS
      WHERE TABLE_NAME = 'members';

EXPLAIN SELECT * FROM wallets WHERE  user_id = 11 AND account_id = 19 AND bank_id = 2 ;

EXPLAIN SELECT * FROM members partition(p0) WHERE username = 'username01';

EXPLAIN SELECT * FROM members WHERE created_at>'1999/01/01';
