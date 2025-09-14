package sqlqueries

const INSERT_NEW_USER = "INSERT INTO users VALUES ($1, $2)"
const INSERT_NEW_ORDER = `
INSERT INTO orders (username, product_ids, total_price) VALUES ($1, $2, $3)
`
const SELECT_TOTAL_AMOUNT = "SELECT SUM(price) FROM products WHERE id = ANY($1)"
const SELECT_PRODUCTS = "SELECT * FROM products"
const SELECT_PRODUCT_BY_ID = "SELECT * FROM products WHERE id = $1"
const SELECT_USER = "SELECT password FROM users WHERE username = $1"

const UPDATE_PRODUCT = "UPDATE products SET left_in_stock = left_in_stock - 1 WHERE id = $1 RETURNING left_in_stock"