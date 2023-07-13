# A RealWorld Go lang backend application
The purpose of this codebase is to provide a strong MVC Restful ORM API foundation with CRUD capabilities, integrated with RabbitMQ and mock testing. 

## Frameworks in use

- Fiber
  > Fiber is a Go web framework built on top of Fasthttp, inspired in Express
  
  Documentation: https://docs.gofiber.io/
  
- AMQP
  > An AMQP 0-9-1 Go client maintained by the RabbitMQ

  Documentation: https://pkg.go.dev/github.com/rabbitmq/amqp091-go?utm_source=godoc
  
- Testify
  > A toolkit with common assertions and mocks that plays nicely with the standard library

  Documentation: https://pkg.go.dev/github.com/stretchr/testify?utm_source=godoc

- gorm pgx driver
  > GORM PostgreSQL driver 

  Documentation: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

- gorm
  > The fantastic ORM library for Golang, aims to be developer friendly

  Documentation: https://gorm.io/docs/
  


## Project Initialization 

Clone the repository

```
git clone https://github.com/Valikmeister/go-lang-backend.git
```
& Run

```
go mod download
```

## How to run the project locally 

Inside the /app directory run

```
go run main.go
```
It will start a server on localhost at port 3000
