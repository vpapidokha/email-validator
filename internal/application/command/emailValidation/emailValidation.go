package emailValidation

import (
	"fmt"
	"log"

	emailverifier "github.com/AfterShip/email-verifier"
)

func ValidateEmail(email Email) (*emailverifier.SMTP, error) {
	verifier := emailverifier.NewVerifier().EnableSMTPCheck()
	splitEmail := verifier.ParseAddress(email.Email)

	result, err := verifier.CheckSMTP(splitEmail.Domain, splitEmail.Username)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return result, err
	}

	log.Print("email validation result", result)

	return result, nil
}
