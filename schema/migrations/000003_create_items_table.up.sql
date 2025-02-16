CREATE TABLE items (
    id serial not null unique,
    name varchar(255) not null,
    description text,
    price int not null,
    -- inventory_count int default 0,
    PRIMARY KEY (id)
);