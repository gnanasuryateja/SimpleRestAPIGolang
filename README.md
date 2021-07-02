# THIS PROJECT IS CURRENTLY INPROGRESS

# SimpleRestAPIGolang
This code has the basic RestAPI developed in Golang with a mocked DB.

# Steps to run the project
```
clone or download the repository
```
```
expand the project structure
```
```
open the terminal in the main folder
```
```
use command:
go run main.go
to run the project
```

# Sample Request and Response
Request 1: To get single User

url
```
http://localhost:8080/getUser/1
```
Response
```
{
    "id": 1,
    "fname": "Steve",
    "city": "LA",
    "phone": 1234567890,
    "height": 5.8,
    "Married": true
}
```
Request 2: To get multiple Users

url
```
http://localhost:8080/getUsers/1,2,3
```
Response
```
[
    {
        "id": 1,
        "fname": "Steve",
        "city": "LA",
        "phone": 1234567890,
        "height": 5.8,
        "Married": true
    },
    {
        "id": 2,
        "fname": "Balmer",
        "city": "NY",
        "phone": 9876543210,
        "height": 5.1,
        "Married": true
    },
    {
        "id": 3,
        "fname": "George",
        "city": "NJ",
        "phone": 9012345678,
        "height": 6.2,
        "Married": false
    }
]
```

# Error Message
------------------------
Message | User not found
