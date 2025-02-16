CREATE TABLE inventory (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    merch_id int NOT NULL,
    quantity int NOT NULL DEFAULT 0,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_merch FOREIGN KEY (merch_id) REFERENCES merch(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_merch UNIQUE (user_id, merch_id)
);