package sqlqueries

const INSERT_NEW_PARTNER = "INSERT INTO partner_info (name, access_token) VALUES ($1, $2)"

const SELECT_PARTNER_BY_TOKEN = "SELECT partner_id FROM partner_info WHERE access_token = $1"