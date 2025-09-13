CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    username VARCHAR(64),
    product_ids INT[],
    total_price DECIMAL
)
