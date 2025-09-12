CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(64) PRIMARY KEY,
    password VARCHAR(32),
    registration_date TIMESTAMPTZ DEFAULT NOW()
)
