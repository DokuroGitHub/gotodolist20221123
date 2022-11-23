# MySQL
CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(150) CHARACTER SET utf8 NOT NULL,
  `status` enum('Doing','Finished') DEFAULT 'Doing',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# Docker MySQL
docker run -d --name demo-mysql -p 8001:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=todo_db mysql:8.0
# Login MySQL
docker exec -it demo-mysql mysql -u root -p todo_db

# Packages
go mod init github.com/DokuroGitHub/gotodolist20221123.git
# installed packages
godoc --http :6060
# Get dependencies
go get .
# Run
go run .