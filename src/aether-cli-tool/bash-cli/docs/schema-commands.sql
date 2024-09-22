-- Description: This file contains the schema commands for the Aether_DB database
-- docker run -itd --name CDB_MySQL -e MYSQL_ROOT_PASSWORD=12345 -p3306:3306 mysql:latest
CREATE DATABASE Aether_DB;
-- create
CREATE TABLE Aether_Node (
  Node_Count int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Node_Name varchar(50), 
  Node_IP varchar(20),
  Node_User varchar(50), 
  Node_Pass varchar(100),
  Node_Details JSON,
  Node_Type varchar(20),
  Node_Status varchar(20)
)

-- insert
-- pre determine the schema of the json
-- keys -- Architecture:, CPU(s):, Model_Name:, OS:, Version:, 