CREATE TABLE IF NOT EXISTS partner_info (
    partner_id SERIAL PRIMARY KEY,
    name VARCHAR(128),
    access_token VARCHAR(64),
    registration_date TIMESTAMPTZ DEFAULT NOW()
)
