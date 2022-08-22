package google

import (
	"context"
	"errors"
	"google.golang.org/api/idtoken"
	"strings"

	authDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/auth"
)

type GoogleAuthenticator struct {
	config authDomain.GoogleSignInConfig
}

func (g GoogleAuthenticator) ValidateIdToken(ctx context.Context, token string) (*authDomain.AuthenticatedGoogleUser, error) {
	payload, err := idtoken.Validate(ctx, token, "")
	if err != nil {
		return nil, err
	}

	claims := payload.Claims

	user := &authDomain.AuthenticatedGoogleUser{
		UserId: payload.Subject,
	}

	err = validateAudience(g.config.Audiences, (claims)["aud"].(string))

	if err != nil {
		return nil, err
	}

	email, ok := (claims)["email"]

	if ok {
		user.Email = strings.TrimSpace(strings.ToLower(email.(string)))
	} else {
		return nil, errors.New("email not found in claims")
	}

	firstName, ok := claims["given_name"]

	if ok {
		user.FirstName = firstName.(string)
	}

	lastName, ok := claims["family_name"]

	if ok {
		user.LastName = lastName.(string)
	}

	return user, nil
}

// Validate that we have a valid audience, using a list since we have different audiences for iOS and Android.
func validateAudience(valid []string, audience string) error {
	for _, s := range valid {
		if s == audience {
			return nil
		}
	}
	return errors.New("audience provided does not match aud claim in the JWT, audience " + audience)
}
