CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL,
    category VARCHAR(255)
);
