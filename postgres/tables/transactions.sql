DROP TABLE IF EXISTS transactions;

CREATE TABLE transactions(
    id VARCHAR PRIMARY KEY,
    amount FLOAT NOT NULL,
    source TEXT NOT NULL
)