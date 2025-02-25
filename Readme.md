# E-commerce Project - Go BACKEND

This is an e-commerce project built with Go. The project aims to provide a robust and scalable platform for online shopping.

## Project Architecture and stack
For a complete rundown of the bigger picture as well as the stack used visit [Architecture](/docs/architecture.md)


## Project Structure

```plaintext
ecommerce/
├── bin/                    # Compiled binaries
├── cmd/                    # Main applications for the project
│   └── ecommerceBackend/   # Backend application
│       └── main.go         # Entry point of the application
├── docs/                   # Documentation files
├── internal/               # Private application and library code
│   ├── api/                # Routes and database connection pool
│   ├── app/                # Main application
│   │   └── server.go       # Code for creating and running the server
│   ├── database/           # Database related files
│   ├── helpers/            # Helper functions and utilities
│   └── types/              # Type definitions and database models
├── .gitignore              # Git ignore file
├── go.mod                  # Go module file
├── README.md               # Project README file
├── .env                    # Environment variables file
└── LICENSE                 # Project license file
```

## Getting Started

To get started with the project, clone the repository and run the following commands:

```sh
git clone https://github.com/yourusername/ecommerce.git
cd ecommerce
go mod tidy
go run cmd/ecommerceBackend/main.go
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the [MIT LICENSE](/LICENSE).
