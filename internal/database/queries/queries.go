package queries

import "github.com/GRACENOBLE/ecommerce/internal/types"

var User = types.UserQueries{
	CreateUser: `INSERT INTO users (name, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id, name, role`,
	UpdateUser: `UPDATE users SET name = $1, password = $2, email = $3 WHERE id = $4`,
	DeleteUser: `DELETE FROM users WHERE id = $1`,
    CheckIfUserExists : `SELECT name FROM users where id = $1`,
    GetUserByEmail: `SELECT id, name, password, email FROM users WHERE email = $1`,
}

var Product = types.ProductQueries