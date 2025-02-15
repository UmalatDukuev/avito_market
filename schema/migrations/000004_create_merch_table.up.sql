CREATE TABLE merch (
    id serial not null unique,
    name varchar(255) not null,
    price int not null
);
INSERT INTO merch (name, price)
VALUES ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);