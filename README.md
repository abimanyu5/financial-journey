
# Financial Journey API Documentation
Welcome to the documentation for the Financial Journey API. This API provides endpoints to manage user transactions, master data, and goals related to financial management.

## Base URL
The base URL for all API endpoints is:


## Authentication
To access protected routes, you need to authenticate using JWT (JSON Web Tokens). After successful authentication, include the token in the Authorization header of your requests.

Example:

```makefile
Authorization: Bearer <token>
```
 ## Endpoints

### User Authentication

Register

URL: /register

Method: POST

Description: Register a new user.

Request Body:

```json
{
    "username": "example_user",
    "password": "example_password"
}
```

Response: Returns the newly registered user's details along with a JWT token.

Login

URL: /login

Method: POST

Description: Login to the application.

Request Body:

```json

{
    "username": "example_user",
    "password": "example_password"
}
```

Response: Returns a JWT token upon successful login.

# User Routes

### Get User Profile

URL: /api/user

Method: GET

Description: Get the profile details of the authenticated user.

# Admin Routes

Admin Dashboard

URL: /api/admin

Method: GET

Description: Access the admin dashboard.

Transaction Routes

## Insert Transaction

URL: /api/transactions

Method: POST

Description: Add a new transaction.

Authorization: Required

Request Body: Details of the transaction.

Response: Details of the inserted transaction.

## Update Transaction

URL: /api/transactions/:id

Method: PUT

Description: Update an existing transaction.

Authorization: Required

Request Body: New details of the transaction.

Response: Details of the updated transaction.

## Delete Transaction

URL: /api/transactions/:id

Method: DELETE

Description: Delete a transaction.

Authorization: Required

Response: Success message.

## Get All Transactions

URL: /api/transactions

Method: GET

Description: Get all transactions of the authenticated user.

Authorization: Required

Response: List of transactions.

# Master Data Routes

## Insert Master Data

URL: /api/masters

Method: POST

Description: Add new master data.

Authorization: Required

Request Body: Details of the master data.

Response: Details of the inserted master data.

## Update Master Data

URL: /api/masters/:id

Method: PUT

Description: Update existing master data.

Authorization: Required

Request Body: New details of the master data.

Response: Details of the updated master data.

## Delete Master Data

URL: /api/masters/:id

Method: DELETE

Description: Delete master data.

Authorization: Required

Response: Success message.

### Get All Master Data

URL: /api/masters

Method: GET

Description: Get all master data.

Authorization: Required

Response: List of master data.

### Get Master Data by ID

URL: /api/masters/:id/data

Method: GET

Description: Get master data by ID.

Authorization: Required

Response: Details of the master data.

# Goals Routes

## Insert Goals

URL: /api/goals

Method: POST

Description: Add a new goal.

Authorization: Required

Request Body: Details of the goal.

Response: Details of the inserted goal.

### Update Goals

URL: /api/goals/:id

Method: PUT

Description: Update an existing goal.

Authorization: Required

Request Body: New details of the goal.

Response: Details of the updated goal.

 ### Delete Goals

URL: /api/goals/:id

Method: DELETE

Description: Delete a goal.

Authorization: Required

Response: Success message.

 ### Get All Goals

URL: /api/goals

Method: GET

Description: Get all goals of the authenticated user.

Authorization: Required

Response: List of goals.

  ### Get Goals by ID

URL: /api/goals/:id/data

Method: GET

Description: Get goals by ID.

Authorization: Required

Response: Details of the goals.

## Running the Application

Ensure you have set the environment variable PORT to specify the port on which the server should listen.

Example:

```go
PORT=8080 go run main.go
```
