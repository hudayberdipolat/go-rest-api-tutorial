# go-rest-api-tutorial

# user-service

# REST API


1. GET /users --list of users                status Code --200,404,500
2. GET /users/:id --user by id               status Code --200,404,500 
3. POST /users --create user                 status Code --201,4xx, Header Location:url
4. PUT /users/:id --fully update user        status Code --204/200,404,,400,500
5. PUTCH /users/:id --partially update user  status Code --204/200,404,,400,500
6. DELETE /users/:id --delete user by id     status Code --200,404,400 

# Logger 

github.com/sirupsen/logrus