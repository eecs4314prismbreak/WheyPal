# WheyPal
## Motivation
This project elverages a REST API with users and auth implementation through Javascript Web Tokens (JWT)

## Libraries
### Gin
Gin is used as the router due to its high speeed and similarity to the standard library

### jwt-go
Jwt-go and jwt-go-middleware is used to follow standard practices to create and validate JWTs

## Structure
main.go has the creation of the router with routes.go countaining the API endpoints

### User
user.go conatins the relevant structs such as the user service
routes.go contains the functions of the user service
repo.go contains the data access layer of the user service

### Auth
auth.go contains relevant structs such as the auth service, requests and response structs
routes.go contains the functions of the auth service
repo.go contains the data access layer for the user service
utils.go contains helper functions for auth such as retrieving claims or the jwt middleware