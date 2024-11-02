To run app:
``` bash
go run main.go serve
```

To use:
``` bash
curl -X POST "http://localhost:8080/api/validate-email" -H "Content-Type: application/json" --data '{"email": "test-email@gmail.com"}'
```