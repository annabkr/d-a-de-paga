DROP TABLE IF EXISTS transactions;

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    amount FLOAT NOT NULL,
    source TEXT NOT NULL
)