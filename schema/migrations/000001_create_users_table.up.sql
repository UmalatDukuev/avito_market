CREATE TABLE users (
    id serial not null unique,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    coins int DEFAULT 1000
);
/*
 migrate -database 
 "postgres://postgres:03795@localhost:5432/market?sslmode=disable" 
 -path ./schema/migrations up
 */