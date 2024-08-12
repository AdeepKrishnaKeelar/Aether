mysql-server-start:
	docker run --name=CDB_MySQL -e MYSQL_ROOT_PASSWORD=12345 -d mysql:latest
	docker exec -it CDB_MySQL mysql -uroot -p12345 -e "CREATE DATABASE Aether_DB;"

binary-build:
	go build -o main go/main.go