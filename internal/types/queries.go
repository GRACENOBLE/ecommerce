package types

type UserQueries struct {
    CreateUser string
    UpdateUser string
    DeleteUser string
	CheckIfUserExists string
	GetUserByEmail string
	GetUserByID string
}

type ProductQueries struct {
	CreateProduct string
	GetAllProducts string
	GetProductById string
	UpdateProduct string
	DeleteProduct string
}