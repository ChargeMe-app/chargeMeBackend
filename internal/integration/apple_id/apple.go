package apple_id

import (
	"context"
	"fmt"
	"github.com/Timothylock/go-signin-with-apple/apple"
	authDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/auth"
)

type AppleAuthenticator struct {
	config authDomain.AppleSignInConfig
}

func (a AppleAuthenticator) ValidateAuthorizationToken(
	ctx context.Context,
	token string,
	isAndroid bool,
) (*authDomain.AuthenticatedAppleUser, error) {
	var clientId string

	if isAndroid {
		clientId = a.config.AndroidClientId
	} else {
		clientId = a.config.ClientId
	}

	secret, err := apple.GenerateClientSecret(a.config.PrivateKey, a.config.TeamId, clientId, a.config.KeyId)
	if err != nil {
		return nil, err
	}

	client := apple.New()

	req := apple.AppValidationTokenRequest{
		ClientID:     clientId,
		ClientSecret: secret,
		Code:         token,
	}

	var resp apple.ValidationResponse

	err = client.VerifyAppToken(ctx, req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		fmt.Printf("apple returned an error: %s - %s\n", resp.Error, resp.ErrorDescription)
		if err != nil {
			return nil, err
		}
	}

	userId, err := apple.GetUniqueID(resp.IDToken)
	if err != nil {
		return nil, err
	}

	claim, err := apple.GetClaims(resp.IDToken)
	if err != nil {
		return nil, err
	}

	email := (*claim)["email"].(string)

	return &authDomain.AuthenticatedAppleUser{
		AppleUserId: userId,
		Email:       email,
	}, nil
}
