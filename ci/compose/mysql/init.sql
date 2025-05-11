-- 事务隔离级别
set session transaction isolation level read committed;
SET GLOBAL tx_isolation='READ-COMMITTED';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'Eisoo.com123' WITH GRANT OPTION;
FLUSH PRIVILEGES;
