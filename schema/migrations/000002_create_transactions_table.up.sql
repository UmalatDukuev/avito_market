CREATE TABLE transactions (
    id serial not null unique,
    from_user_id int not null,
    to_user_id int not null,
    amount int not null,
    transaction_time timestamp default current_timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (to_user_id) REFERENCES users(id) ON DELETE CASCADE
);