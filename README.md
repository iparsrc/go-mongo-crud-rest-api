# go-mongo-CRUD-web
This is a simple CRUD operation on mongodb using golang. It's using http req to operate.  
Endpoints are:
```
/ping
/users/find
/users/create
/users/delete
/users/update
```
## /ping  
Send a GET request.  
Returns string:  
```
pong
```
This is only to check if the server is running.
## /users/find  
Send a GET request(only email). EX)
```
localhost:8080/users/find?email=akbariparsa1209@gmail.com
```
Returns a json:
```
{
  "_id": "<user id>",
  "name": "<user name>",
  "email": "<user email>",
  "password": "",
}
```
