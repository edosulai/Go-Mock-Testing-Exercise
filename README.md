# Go-Mock-Testing-Exercise

# README.md

This repository is an exercise for mock testing in Golang using the "product" entity as an example. The project is structured as follows:

```
Go-Mock-Testing-Exercise/
├── controllers
├── database
├── helpers
├── middlewares
├── models
├── repository
│   ├── product.go
│   └── product_mock.go
├── router
├── service
│   ├── product.go
│   └── product_test.go
├── admin.http
├── go.mod
├── go.sum
├── main.go
└── user.http
```

## Dependencies

The project uses the following Go modules:

- github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
- github.com/gin-gonic/gin v1.9.0
- github.com/golang-jwt/jwt v3.2.2+incompatible
- golang.org/x/crypto v0.7.0
- gorm.io/driver/postgres v1.5.0
- gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11

As well as several indirect dependencies, which can be found in the `go.mod` file.

## Description

This project is focused on implementing mock tests for the `Product` entity. The `product_test.go` file contains the test cases, while the `product_mock.go` file contains the mock implementation used in those tests.

The rest of the project structure is organized as follows:

- `controllers`: Contains the Gin HTTP controllers.
- `database`: Contains the database configuration and migration scripts.
- `helpers`: Contains helper functions used throughout the project.
- `middlewares`: Contains Gin middlewares.
- `models`: Contains the GORM models for the database tables.
- `repository`: Contains the repository layer implementation, as well as the mock implementation.
- `router`: Contains the Gin router setup.
- `service`: Contains the service layer implementation, as well as the test cases.

The `admin.http` and `user.http` files contain sample HTTP requests for testing purposes.

To run the project, simply run `go run main.go`.