package validator

import (
	"errors"
	"regexp"
)

func ValidateURL(url string) error {
	if url == "" {
		return errors.New("empty body request")
	}

	pattern := `^(https?://|www.)?[a-zA-Z0-9-]{1,256}([.][a-zA-Z-]{1,256})?([.][a-zA-Z]{1,30})([/]?[a-zA-Z0-9/?=%&#_.-]+)`

	valid, err := regexp.Match(pattern, []byte(url))
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("url is invalid")
	}

	return nil
}

func ValidateToken(token string) error {
	if token == "" {
		return errors.New("empty body request")
	}

	pattern := `^[a-zA-Z0-9_]{10}$`

	valid, err := regexp.Match(pattern, []byte(token))
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("token is invalid")
	}

	return nil
}
