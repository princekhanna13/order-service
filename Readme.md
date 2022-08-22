## Order Service

This app is for order service.

### Development Environment Requirements

| Tool      | Version      | Usage                        | Install                                                                  |
|-----------|--------------|------------------------------|--------------------------------------------------------------------------|
| `go`      | `>= v1.18.0` | Programming language used    | [golang](https://golang.org/)                                            |

- GoLang 
- Postman

### Project Set up

```bash
cd order-service
touch .env ## Needed in case of local env is different from config/app.env
go run cmd/server/main.go
```

### APIs

- /order: For Creating an Order Entry
  - Input: ``` {
    "bank_id":"test",
    "amount":2000, //In Euros
    "information":"Test Transaction" //Optional Field
    }```
  
  - Output: ```{
      "code": "Reponse Code",
      "message": Response Message"
      }```
