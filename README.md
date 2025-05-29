# E-Commerce Backend API in Go

A comprehensive e-commerce backend API built with Go standard packages, demonstrating best practices in API development.

## Features

- User Authentication with JWT
- Product Management
- Shopping Cart Operations
- Order Processing
- Payment Integration
- File Upload for Product Images
- Shipping Integration
- Input Validation
- Error Handling
- Database Operations (PostgreSQL)

## Project Structure

```
├── cmd/
│   └── api/            # Application entry point
├── models/            # Data models and validation
├── handlers/          # HTTP request handlers
├── middleware/        # Authentication and request processing
├── service/          # Business logic layer
├── db/               # Database operations and migrations
└── config/           # Configuration management
```

## Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Git

## Setup Instructions

1. Clone the repository:
```bash
git clone <repository-url>
cd ECOM
```

2. Set up the database:
```bash
# Create PostgreSQL database
createdb ecom_db
```

3. Set up environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=your_user
export DB_PASSWORD=your_password
export DB_NAME=ecom_db
export JWT_SECRET=your_secret_key
```

4. Run the application:
```bash
go run cmd/api/main.go
```

## API Documentation

### Authentication
- POST /api/v1/auth/register - Register new user
- POST /api/v1/auth/login - Login user

### Products
- GET /api/v1/products - List all products
- GET /api/v1/products/{id} - Get single product
- POST /api/v1/products - Create product (admin only)
- PUT /api/v1/products/{id} - Update product (admin only)
- DELETE /api/v1/products/{id} - Delete product (admin only)

### Cart
- GET /api/v1/cart - Get user's cart
- POST /api/v1/cart/items - Add item to cart
- PUT /api/v1/cart/items/{id} - Update cart item
- DELETE /api/v1/cart/items/{id} - Remove item from cart

### Orders
- POST /api/v1/orders - Create order
- GET /api/v1/orders - List user's orders
- GET /api/v1/orders/{id} - Get order details

## Learning Resources

### Go Fundamentals
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)

### API Development
- [Standard Package HTTP Documentation](https://golang.org/pkg/net/http/)
- [REST API Design Best Practices](https://golang.org/doc/effective_go#interfaces)
- [Go Web Examples](https://gowebexamples.com/)

### Authentication
- [JWT Go Implementation](https://github.com/golang-jwt/jwt)
- [Go Security Practices](https://golang.org/security)

### Database
- [Database SQL Package](https://golang.org/pkg/database/sql/)
- [PostgreSQL Driver](https://github.com/lib/pq)

### Testing
- [Testing Package Documentation](https://golang.org/pkg/testing/)
- [HTTP Testing in Go](https://golang.org/pkg/net/http/httptest/)

## Best Practices Implemented

1. **Project Structure**
   - Clear separation of concerns
   - Modular design
   - Interface-based design

2. **Error Handling**
   - Custom error types
   - Proper error propagation
   - Meaningful error messages

3. **Authentication**
   - JWT implementation
   - Secure password hashing
   - Role-based access control

4. **Input Validation**
   - Request validation
   - Data sanitization
   - Type checking

5. **Database**
   - Connection pooling
   - Prepared statements
   - Transaction management

6. **Testing**
   - Unit tests
   - Integration tests
   - Mock interfaces

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 