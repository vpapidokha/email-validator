To run you need to setup two env variables:
``` bash
export SMTP_HOST_NAME="smtp.gmail.com"
export SMTP_MAIL_ADDRESS="example@gmail.com"
```

After that:
``` bash
go run cmd/cli/main.go
```

To use:
``` bash
curl -X POST "http://localhost:8080/validate-email" -H "Content-Type: application/json" --data '{"email": "test-email@gmail.com"'
```