package database

var ProductStatements = map[string]string {
	"CreateProduct" : "INSERT INTO products (name) VALUES(?)",
	"GetProductByID" : "SELECT id, name FROM products WHERE id = ? AND is_active = 1",
	"UpdateProduct" : "UPDATE products SET name = ?, updated_at = ? WHERE id = ?",
	"DeleteProduct" : "UPDATE products SET is_active = 0, updated_at = ? WHERE id = ?",
	"GetProductWithVariant" : `SELECT p.name as product_name, v.variant_name , v.quantity FROM products p
	INNER JOIN variants v on p.id = v.product_id
	WHERE p.is_active = 1 and v.is_active = 1`,
}

var VariantStatements = map[string]string {
	"CreateVariant" : "INSERT INTO variants (variant_name, quantity, product_id) VALUES(?, ?, ?)",
	"GetVariantByID" : "SELECT id, variant_name, quantity, product_id FROM variants WHERE id = ? AND is_active = 1",
	"UpdateVariant" : "UPDATE variants SET variant_name = ?, quantity = ?, product_id = ?, updated_at = ? WHERE id = ?",
	"DeleteVariant" : "UPDATE variants SET is_active = 0, updated_at = ? WHERE id = ?",
}