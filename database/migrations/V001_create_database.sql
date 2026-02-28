

CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL primary_key,
    username TEXT,
    email_id TEXT,
    mobile_number DECIMAL(10,0),
);

CREATE TABLE IF NOT EXISTS expences(
    transaction_id SERIAL primary_key,
    transaction_name TEXT,
    transaction_amount BIGINT,
    users_spend JSONB
);

CREATE TABLE IF NOT EXISTS balances(
    transaction_id INT,
    lender_id INT REFERENCES users(user_id),
    borrower INT REFERENCES users(user_id),
    CONSTRAINT fk_transactions
        FOREIGN KEY transaction_id
        REFERENCES expences(transaction_id)
        ON DELETE CASCADE
);
