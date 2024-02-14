# go-rest-api-tutorial



# user-service

# REST API



GET /users --list of users                status Code --200,404,500
GET /users/:id --user by id               status Code --200,404,500 
POST /users --create user                 status Code --201,4xx, Header Location:url
PUT /users/:id --fully update user        status Code --204/200,404,,400,500
PUTCH /users/:id --partially update user  status Code --204/200,404,,400,500
DELETE /users/:id --delete user by id     status Code --200,404,400 
