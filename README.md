# momen-backend
# Money Management
This app is for help user to make records of their income or expenses
## API Spec

### Register User
	Request:
	- Method: POST
	- Endpoint: "/api/v1/register"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"name" "string",
			"email" "string",
			"password" "string"
		}
	Response:
	{
		"meta": {
        "message": "string",
        "code": "integer",
        "status": "string"
    },
    "data": {
        "id": "integer",
        "name": "string",
        "email": "string",
        "token": "string"
    }
	}

### Login User
	Request:
	- Method: POST
	- Endpoint: "/api/v1/login"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"email" "string",
			"password" "string"
		}
	Response:
	{
		"meta": {
        "message": "string",
        "code": "integer",
        "status": "string"
    },
    "data": {
        "id": "integer",
        "name": "string",
        "email": "string",
        "token": "string"
    }
	}

### Get Transactions
	Request:
	- Method: GET
	- Endpoint: "/api/v1/transactions"
	- Header:
		- Content-Type: application/json
	Response:
	{
		"meta": {
        "message": "string",
        "code": "string",
        "status": "string"
    },
    "total_transaction": "integer",
    "transactions": [
        {
            "id": "integer",
            "user_id": "integer",
            "name": "string",
            "description": "string",
            "category": "string",
            "amount": "integer"
        },
        {
            "id": "integer",
            "user_id": "integer",
            "name": "string",
            "description": "string",
            "category": "string",
            "amount": "integer"
        }
    ]
	}

### Create Transaction
	Request:
	- Method: POST
	- Endpoint: "/api/v1/transaction"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"name": "string",
       "description": "string",
       "category": "string",
       "amount": "integer"
		}
	Response:
	{
		"meta": {
        "message": "string",
        "code": "integer",
        "status": "string"
    },
    "data": {
        "id": "integer",
        "user_id": "integer",
        "name": "string",
        "description": "string",
        "category": "string",
        "amount": "integer"
    }
	}

### Edit Transaction
	Request:
	- Method: PUT
	- Endpoint: "/api/v1/transaction/{id}"
	- Header:
		- Content-Type: application/json
	- Body:
		{
			"name": "string",
       "description": "string",
       "category": "string",
       "amount": "integer"
		}
	Response:
	{
		"meta": {
        "message": "string",
        "code": "integer",
        "status": "string"
    },
    "data": {
        "id": "integer",
        "user_id": "integer",
        "name": "string",
        "description": "string",
        "category": "string",
        "amount": "integer"
    }
	}

### Delete Transaction
	Request:
	- Method: DELETE
	- Endpoint: "/api/transaction/{id}"
	- Header:
		- Content-Type: application/json
	Response:
	{
		"meta": {
        "message": "string",
        "code": "integer",
        "status": "string"
    },
    "data": true
	}
