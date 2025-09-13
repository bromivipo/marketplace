CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL,
    left_in_stock INT,
    provider_id INT,
    category VARCHAR(255)
);

INSERT INTO products (name, price, left_in_stock, provider_id, category) VALUES 
    ('doll', '15.50', 1, 1, 'toys'),
    ('ball', '10.00', 1, 2, 'toys');
