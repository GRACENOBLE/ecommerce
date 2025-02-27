package queries

import "github.com/GRACENOBLE/ecommerce/internal/types"

var User = types.UserQueries{
	CreateUser: `INSERT INTO users (name, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id, name, role`,
	UpdateUser: `UPDATE users SET name = $1, password = $2, email = $3 WHERE id = $4`,
	DeleteUser: `DELETE FROM users WHERE id = $1`,
    CheckIfUserExists : `SELECT name FROM users where id = $1`,
    GetUserByEmail: `SELECT id, name, password, email FROM users WHERE email = $1`,
}

var Product = types.ProductQueries {
    CreateProduct: `INSERT INTO products (name, description, price, stock_quantity, image_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`,
    GetAllProducts: `SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products`,
    GetProductById: `SELECT id, name, description, price, stock_quantity, image_url, created_at, updated_at FROM products WHERE id = $1`,
    UpdateProduct: `UPDATE products SET name = COALESCE($1, name), description = COALESCE($2, description), price = COALESCE($3, price), stock_quantity = COALESCE($4, stock_quantity), image_url = COALESCE($5, image_url), updated_at = CURRENT_TIMESTAMP WHERE id = $6 RETURNING id, name, description, price, stock_quantity, image_url, updated_at;`,
    DeleteProduct: `DELETE FROM products WHERE id = $1;`,

}