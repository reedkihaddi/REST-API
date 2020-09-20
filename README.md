# Go REST API

A simple RESTful API to help in future projects. 

Features of the API: 

* RESTful endpoints
* JWT authentication
* CRUD operations
* Configuration
* Logging 
* HealthCheck
* API Docs using SwaggerUI
* Error Handling
* Data migration
* Data seed

The API uses the following Go packages:

* Routing: [Gorilla](https://github.com/gorilla/mux)
* Database Migration: [Goose](https://bitbucket.org/liamstask/goose/src)
* Logging: [Zap](https://github.com/uber-go/zap)
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)

## Getting started

```bash
# Download
git clone https://github.com/reedkihaddi/REST-API/

cd REST-API

# Make sure you Postgres setup.
# Create a new database.

# Change the configuration file in ./config and dbconf.yml in ./db.

# Run goose up and add seed data.

# Run the server 
go run cmd/main.go
```

The application runs as an HTTP server. It provides the following RESTful endpoints:

* `GET /token`: Returns a JWT token
* `GET /product/:id`: Returns a product
* `GET /products`: Returns a list of products (Requires a JWT token in authorization header to access)
* `POST /product`: Creates a new product
* `PUT /product/:id`: Update a product
* `DELETE /product/:id`: Delete a product

For health check, try `/health`, it should return status of the service.

To try out the API, either use some API tool like [Postman](https://www.getpostman.com/) or go to `/docs/` and use the swagger tools.

## Project Layout

```bash
.
├───cmd                     Main applications of the project
│   │   main.go             Initialize the server                            
│   └───router              Contains the gorilla mux router and the routes
│           router.go
│           routes.go
├───config                  Config environments
│       config.go           
│       local.env       
├───db                      For database migration
│   │   dbconf.yml
│   └───migrations
│           20200920231504_product_schema.sql
├───internal                
│   ├───docs                Swagger docs
│   │       docs.go
│   │       swagger.json
│   │       swagger.yaml
│   └───handlers            Handlers to use for the API
│           api.go          API endpoints
│           healthcheck.go  Health check for the API
│           token.go        JWT token middleware
├───pkg                     
│   ├───db                  Database code
│   │       product.go
│   ├───logging             Logging middleware
│   │       logger.go
│   │       server.log
│   └───models              Structs used in the project
│           product.go
└───seed                    Demo data
        testdata.sql
```

### TODO

- [ ] ADD TESTING
- [ ] Use Viper for better handling env and config
- [ ] Add Docker deployment
- [ ] Add Kubernetes deployment
- [ ] Add database transactions
- [ ] Use mocking
- [ ] Follow SOLID & Clean architecture
