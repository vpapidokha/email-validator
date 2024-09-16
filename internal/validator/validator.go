package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/badoux/checkmail"
)

var (
	smtpServerHostName    = os.Getenv("SMTP_HOST_NAME")
	smtpServerMailAddress = os.Getenv("SMTP_MAIL_ADDRESS")
)

type Email struct {
	Email string `json:"Email"`
}

type httpError struct {
	status  int
	message string
}

func handleExit(w http.ResponseWriter) {
	if r := recover(); r != nil {
		if he, ok := r.(httpError); ok {
			http.Error(w, he.message, he.status)
		} else {
			panic(r)
		}
	}
}

func exit(status int, message string) {
	panic(httpError{status: status, message: message})
}

// NewRouter generates the router used in the HTTP Server
func NewRouter() *http.ServeMux {
	// Create router and define routes and return that router
	router := http.NewServeMux()

	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	router.HandleFunc("POST /validate-email", ValidateEmailHandler)

	return router
}

func ValidateEmailHandler(w http.ResponseWriter, r *http.Request) {
	defer handleExit(w)
	var targetEmail Email

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &targetEmail)
	if err != nil {
		fmt.Println(err)
		exit(422, "Can't unmarshall the body.")
	}

	ValidateEmail(targetEmail.Email)

	encoder := json.NewEncoder(w)
	encoder.Encode("Email address is real and valid")
}

func ValidateEmail(targetEmail string) {
	err := checkmail.ValidateFormat(targetEmail)
	if err != nil {
		fmt.Println(err)
		exit(200, "Email address is not valid.")
	}
	fmt.Printf("Email address format is valid for %v.\n", targetEmail)

	err = checkmail.ValidateHost(targetEmail)
	if err != nil {
		fmt.Println(err)
		exit(200, "Email address host is not valid.")
	}
	fmt.Printf("Email address host is valid for %v.\n", targetEmail)

	err = checkmail.ValidateHostAndUser(smtpServerHostName, smtpServerMailAddress, targetEmail)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		exit(200, "Email address host is not real.")
	}
	fmt.Printf("Email address is real: %v.\n", targetEmail)
}
