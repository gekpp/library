// Code generated by goa v3.0.3, DO NOT EDIT.
//
// auther HTTP client CLI support package
//
// Command:
// $ goa gen library/design

package client

import (
	auther "library/gen/auther"
)

// BuildSigninPayload builds the payload for the auther signin endpoint from
// CLI flags.
func BuildSigninPayload(autherSigninUsername string, autherSigninPassword string) (*auther.SigninPayload, error) {
	var username string
	{
		username = autherSigninUsername
	}
	var password string
	{
		password = autherSigninPassword
	}
	payload := &auther.SigninPayload{
		Username: username,
		Password: password,
	}
	return payload, nil
}
