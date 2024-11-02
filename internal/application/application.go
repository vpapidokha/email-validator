package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go/log"
	"github.com/vpapidokha/email-validator/internal/application/command/emailValidation"
	"github.com/vpapidokha/email-validator/internal/config"
)

func NewRouter(ctx context.Context, configuration *config.Config) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	router.HandleFunc("POST /api/validate-email", validateEmailHandler(configuration))

	return router
}

func validateEmailHandler(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var targetEmail emailValidation.Email

		body, _ := io.ReadAll(r.Body)
		err = json.Unmarshal(body, &targetEmail)
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		}

		result, err := emailValidation.ValidateEmail(targetEmail)
		if err != nil {
			log.Error(err)
			http.Error(w, "Failed to validate email", http.StatusBadRequest)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(result)
	}
}
