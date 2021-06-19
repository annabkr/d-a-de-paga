CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    amount FLOAT NOT NULL,
    source TEXT NOT NULL
)