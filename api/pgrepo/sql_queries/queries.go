package sqlqueries

const SELECT_PRODUCTS = "SELECT * FROM products"
const SELECT_PRODUCT_BY_ID = "SELECT * FROM products WHERE id = $1"
const INSERT_NEW_USER = "INSERT INTO users VALUES ($1, $2)"
const SELECT_USER = "SELECT password FROM users WHERE username = $1"