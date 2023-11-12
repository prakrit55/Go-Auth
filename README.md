# Go-Auth

It is a simple authentication appplication made up of Gin, JWT, and Postgress at the database. It creates user in database by providing input in JSON. We are able to login as the user we created. Also, we can log the user out by providing the phone number and password. It uses bcrypt to encrypt the password in `Util` directory.

It uses `JWT` for authentication. When we login to the user as we created before, a `JWT` token gets generated.

# Pre-Requisite

* Docker Desktop or Docker
* Make command installed
* Golang
* Postman

# How to use it

With docker in machine, we run `make createpostgres` which download a docker image and create a container out of it. We run `make createdb` to create a database named go-auth.

We are able to see the database created by `make postgres` and then `\l`.

We need to create the migrations for a table for database model. We run `make migrate-addtable` and create two migrate tables, which will be created in `migrations` directory. One causes to implement the table and other to destroy the table.

In `_add_user_table.up.sql` we write the structure given below.

CREATE TABLE "users" (

    "id" bigserial PRIMARY KEY,
    
    "username" varchar NOT NULL,
    
    "email" varchar NOT NULL,
    
    "phone" varchar NOT NULL,
    
    "password" varchar NOT NULL
    
)

Here we create users with the following details.

And in `_add_user_table.down.sql` we write `DROP TABLE IF EXISTS users;` 

Then, we run `make migrate-up` to forward the table to database.

--> Now we run `go run main.go`

In Postman, we select `post` method, provide user details with url `localhost:8000/signup` and
username, email, phone, password to create the user. This creates the user and saves the information in the database.

We use `post` method and provide the details like phone and password of saved user with url `localhost:8000/login` to login.

To logout, we select `get` method with url `localhost:8000/logout` and provide user phone and password to logout.

To delete user, we select `get` method with url `localhost:8000/deleteuser` and provide user phone and password to de;lete the specific user.




