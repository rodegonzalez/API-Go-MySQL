# API-Go-MySQL
Simple API using Go+MySQL

# Author
2023-Rode Gonzalez

# Usage
Create a mysql schema to be used as storage. 
Create a mysql user/pass also.
Create a mysql table with columns id, name and description, to be used by this crud example.
Configure mysql credentials and server port in main.go, configuration section.

Execute 'go run .' in a terminal in the folder where this code has been downloaded.

Test API using postman or similar.

## Available EndPoints
* (GET) "/" - It returns a simple message "It works!".
* (GET) "/api/items" - It returns a complete list of items in database.
* (GET) "/api/item/{id}" - It returns the item with id={id}.
* (POST) "/api/item" - It creates a new item. Http Body must have a json element like {"name": "the-new-item-name", "description": "the-new-item-description"} 
* (POST) "/api/item/{id}" - It updates the item with id={id}. Http Body must have a json element like {"name": "the-item-name", "description": "the-item-description"}.
* (DELETE) "/api/item/{id}" - It deletes the item with id={id}. 

## Database spec
Mysql database.
See configuration section in _main.go_ file for modifications.
dbname = test
dbuser = test
dbpass = test

Table used in this exercise: items(
    id int,
    name varchar(45),
    description varchar(45)
)

# License
Distributed under the GNU GENERAL PUBLIC LICENSE. 
See LICENSE.txt for more information.

# Contact
rodegonzalez.com
rode.gonzalez@gmail.com

