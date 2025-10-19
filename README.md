# GoMathAPI

Simple Go API that implements mathematical operations (no frameworks).

## Introduction

Even if the dev life isn't for you, as a cybersecurity enthusiast you should know at least the basics of web development. This project is a simple Go API that implements mathematical operations without using any frameworks. It also helps to understand how web servers work and how to create API services.

Why Go? It is simple, fast and the built-in `net/http` package makes it easy to create APIs without worrying about any fancy frameworks.

Why math? Because it is simple and easy to understand. You don't need to know anything about databases or complex business logic to understand how this API works.

A big thank you to [@elliotminns](https://github.com/elliottminns) (Dreams of Code) for his [video](https://youtu.be/H7tbjKFSg58?si=hep2W0kGxAUFcZ--) which inspired me to create this project.

## Design Features

-   Middlewares
    -   Logging
    -   Authentication (API Key based)

## Application Features

-   Summation (/sum)
    -   Addition
    -   Subtraction
-   Multiplication (/mul)
-   Division (/div)
-   Exponentiation (/pow)
-   Modulo (/mod)

## Usage

For the mathematical operations, you can send a GET request to the `/math/{operation}` endpoint with the required query parameters. For example, to add two numbers, you can send a GET request to `/math/add`. For all operations you submit operands like so, using comma-separated values:

```
/math/sum/5,-2,3
/math/mul/5,10
/math/div/10,2
/math/pow/2,3
/math/mod/10,3
```

The `sum` and `mul` operations can take multiple operands at a time, while the others take exactly two.

## Bonus Features (WIP or Coming Soon)

### Database Usage

Interaction with a relational Database (SQLite/PostgreSQL).

A simple API is cool and all, but CRUD (CREATE, READ, UPDATE, DELETE) operations are a core part of most web applications. Implementing database interactions helps to understand how to persist data and manage state securely in your applications.
