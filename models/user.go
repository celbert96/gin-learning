package models

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) IsValid() error {
	var validationErrors []string
	const MISSING_REQUIRED_FIELD_MSG = "missing required field %s"

	if u.Username == "" {
		validationErrors = append(validationErrors, fmt.Sprintf(MISSING_REQUIRED_FIELD_MSG, "username"))
	}
	if len(strings.Trim(u.Password, " ")) == 0 {
		validationErrors = append(validationErrors, fmt.Sprintf(MISSING_REQUIRED_FIELD_MSG, "password"))
	}

	if len(validationErrors) > 0 {
		for _, e := range validationErrors {
			fmt.Println(e)
		}

		errorMessage := fmt.Sprintf("The following errors were encountered: %s", strings.Join(validationErrors, "; "))
		return errors.New(errorMessage)
	}

	return nil
}
