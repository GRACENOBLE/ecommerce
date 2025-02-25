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

To get started with the project, follow these steps:

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/ecommerce.git
    cd ecommerce
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the application:**

    ```sh
    go run cmd/ecommerceBackend/main.go
    ```

4. **Build binaries for multiple platforms:**

    You can use the provided script to build binaries for Linux, macOS, and Windows. Create a file named `build.sh` in the root of your project and add the following content:

    ```sh
    #!/bin/bash
    set -e

    # Define target architectures in the format "GOOS/GOARCH"
    targets=("linux/amd64" "darwin/amd64" "windows/amd64")

    # Loop through each target
    for target in "${targets[@]}"; do
        # Split the target into OS and architecture
        IFS="/" read -r goos goarch <<< "$target"
        
        # Define output binary name (adjust path and name as needed)
        output="bin/ecommerceBackend-${goos}-${goarch}"
        
        echo "Building for $goos/$goarch..."
        
        # Set environment variables and build
        GOOS=$goos GOARCH=$goarch go build -o "$output" ./cmd/ecommerceBackend
    done

    echo "Builds completed!"
    ```

    Make the script executable:

    ```sh
    chmod +x build.sh
    ```

    Run the script to build the binaries:

    ```sh
    ./build.sh
    ```

5. **Environment variables:**

    Ensure you have a `.env` file in the root of your project with the necessary environment variables configured.

6. **Documentation:**

    For more details on the project architecture and stack, visit the [Architecture](/docs/architecture.md) documentation.

You are now ready to start developing and contributing to the project!

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the [MIT LICENSE](/LICENSE).
