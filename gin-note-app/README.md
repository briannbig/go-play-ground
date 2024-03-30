This is a simple note taking app Using `gin` for routing

- It uses the `gin-gonic/gin` as default router
- No persistence is implemented

## routes
| route         | http method | payload                                  |
| ------------- | ----------- | ---------------------------------------- |
| :8080/welcome | GET         | -                                        |
| :8080/        | GET         | -                                        |
| :8080/        | POST        | `{"title": "string","content":"string"}` |
| :8080/{id}    | GET         | -                                        |
| :8080/{id}    | DELETE      | -                                        |
---