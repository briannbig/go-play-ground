This is a simple note taking app Using `gin` for routing

- It uses the `gin-gonic/gin` as default router
- Uses mysql for database and  [`go-sql-driver/mysql/`](https://github.com/go-sql-driver/mysql/)

## routes
| route         | http method | payload                                  |
| ------------- | ----------- | ---------------------------------------- |
| :8080/welcome | GET         | -                                        |
| :8080/        | GET         | -                                        |
| :8080/        | POST        | `{"title": "string","content":"string"}` |
| :8080/{id}    | GET         | -                                        |
| :8080/{id}    | DELETE      | -                                        |
---

# Setting up database

make sure you have mysql installed
#### Note that this is a one time process

### log into mysql server
```bash
mysql -u root -p
```

### initialize up database
 ```sql
 CREATE DATABASE gin_note_app;
 USE gin_note_app;
 DROP TABLE IF EXISTS notes;

 CREATE TABLE notes(
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title VARCHAR(255),
    content TEXT,
    time_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );
 ```

 or alternatively create database 
 ```shell 
 source /path/to/project/gin-note-app/create_tables.sql
 ```