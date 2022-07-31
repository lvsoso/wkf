
mysql
```
docker run --rm -it -d -p 31306:3306 --name local-mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0.25 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci 

mysql -uroot -p123456 -h127.0.0.1 -P31306
CREATE DATABASE wkf;
USE wkf;

CREATE TABLE IF NOT EXISTS `wkf_tbl`(
   `wkf_id` INT UNSIGNED AUTO_INCREMENT,
   `wkf_title` VARCHAR(100) NOT NULL,
   `wkf_author` VARCHAR(40) NOT NULL,
   `submission_date` DATE,
   PRIMARY KEY ( `wkf_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO wkf_tbl  (wkf_title, wkf_author, submission_date) VALUES ("test1", "test 1", NOW());
INSERT INTO wkf_tbl  (wkf_title, wkf_author, submission_date) VALUES ("test2", "test 2", NOW());
INSERT INTO wkf_tbl  (wkf_title, wkf_author, submission_date) VALUES ("test3", "test 3", NOW());
```

.wkf.yaml
```yaml
DBDSN:
    root:123456@localhost:31306/wkf
```