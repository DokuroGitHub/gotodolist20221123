# MySQL
```sql
CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(150) CHARACTER SET utf8 NOT NULL,
  `status` enum('Doing','Finished') DEFAULT 'Doing',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

# Docker MySQL
```bash
docker run -d --name demo-mysql -p 8001:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=todo_db mysql:8.0
```
# Login MySQL
```bash
docker exec -it demo-mysql mysql -u root -p todo_db
```

# Packages
```go
go mod init github.com/DokuroGitHub/gotodolist20221123.git
```
# installed packages
```go
godoc --http :6060
```
# Get dependencies
```go
go get .
```
# Run
```go
sh
MYSQL_CONNECTION="root:root@tcp(127.0.0.1:8001)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
go run .
```