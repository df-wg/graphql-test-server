# GraphQL Test Server

This project is a sample project, showing how to run the cosmos router with multiple GraphQL servers. 

The main server (in `user_server`) is a simple GraphQL server built using Go and `gqlgen`. It provides an API to manage users with the following fields:
- `id`: Unique identifier for the user.
- `email`: The user's email address.
- `first_name`: The user's first name.
- `last_name`: The user's last name.
- `full_name`: A computed field that concatenates the first name and last name.

The second server (in `user_age_server`) is a super simple GraphQL server that extends the model by adding `age`

The third directory (`cosmos_router`) contains the router configuration for running the server

## Features
- **Query Users**: Fetch a list of users with optional filtering by first name and ordering by email.
- **Create User**: Add a new user to the list.
- **Delete User**: Remove a user by their ID.

All user data is stored in memory, so it will be lost when the server restarts.

## Getting Started
### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [gqlgen](https://gqlgen.com/getting-started/)

### Running

1. Clone the repository:
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Run tests: 
   ```bash
   go test ./...
   ```
3. Run all 3 servers at the same time (in 3 tabs):
    ```bash
    go run user_server/server.go
    go run user_age_server/server.go
    cd cosmos_router && ./router
    ```

## Testing the API
You can test the API using GraphQL Playground, Postman, or any other GraphQL client.

### Example Queries
* Get a list of users:
```graphql
query {
  users {
    id
    email
    full_name
     age
  }
}
```

* Filter users by first name
```graphql
query {
  users(first_name: "John") {
    id
    email
    full_name
  }
}
```
 
* Order users by email
```graphql
query {
  users(orderByEmail: true) {
    id
    email
    full_name
  }
}
```

### Example Mutations
* Create a new user:
```graphql
mutation {
  createUser(input: { email: "john.doe@example.com", first_name: "John", last_name: "Doe" }) {
    id
    full_name
  }
}
```

* Delete a user:
```graphql
mutation {
  deleteUser(id: "U1")
}
```

## Acknowledgments
gqlgen - A Go library for building GraphQL servers.