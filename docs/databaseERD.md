```mermaid
erDiagram
    USERS {
        UUID id PK "Primary Key"
        TEXT name "User's Name"
        TEXT email "Unique Email Address"
        TEXT password "Hashed Password"
        TIMESTAMP created_at "Account Creation Timestamp"
        TIMESTAMP updated_at "Last Update Timestamp"
    }
    PRODUCTS {
        UUID id PK "Primary Key"
        TEXT name "Product Name"
        TEXT description "Product Description"
        NUMERIC price "Price of the Product"
        INT stock_quantity "Available Stock"
        TEXT image_url "URL to Product Image"
        TIMESTAMP created_at "Creation Timestamp"
        TIMESTAMP updated_at "Last Update Timestamp"
    }
    CATEGORIES {
        UUID id PK "Primary Key"
        TEXT name "Category Name"
        TEXT description "Category Description"
        TIMESTAMP created_at "Creation Timestamp"
    }
    PRODUCT_CATEGORIES {
        UUID product_id FK "Foreign Key to Products"
        UUID category_id FK "Foreign Key to Categories"
    }
    ORDERS {
        UUID id PK "Primary Key"
        UUID user_id FK "Foreign Key to Users"
        TEXT status "Order Status"
        NUMERIC total_price "Total Order Price"
        TIMESTAMP created_at "Creation Timestamp"
    }
    ORDER_ITEMS {
        UUID id PK "Primary Key"
        UUID order_id FK "Foreign Key to Orders"
        UUID product_id FK "Foreign Key to Products"
        INT quantity "Quantity of Product"
        NUMERIC price "Price per Unit"
        TIMESTAMP created_at "Creation Timestamp"
    }
    PAYMENTS {
        UUID id PK "Primary Key"
        UUID order_id FK "Foreign Key to Orders"
        TEXT payment_method "Method of Payment"
        TEXT status "Payment Status"
        TIMESTAMP created_at "Payment Timestamp"
    }

    USERS ||--o{ ORDERS : "places"
    ORDERS ||--|{ ORDER_ITEMS : "contains"
    PRODUCTS ||--o{ ORDER_ITEMS : "includes"
    PRODUCTS }o--o{ PRODUCT_CATEGORIES : "classified as"
    CATEGORIES }o--o{ PRODUCT_CATEGORIES : "categorizes"
    ORDERS ||--o{ PAYMENTS : "processed by"
