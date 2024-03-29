This is a simple note taking app demostrating MVC pattern.

- It uses the `net/http` as default
- No persistence is implemented

## routes
|route           | http method | payload |
|----------------|-------------|---------|
|:8080/notes     | GET         | -       |
|:8080/notes/add | POST        |`{"title": "string","content":"string"}`|
---