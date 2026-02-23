-- Create a useres table

CREATE TABLE useres(
    user_id SERIAL primary_key,
    username TEXT,
    email_id TEXT,
    mobile_number DECIMAL(10,0),
);