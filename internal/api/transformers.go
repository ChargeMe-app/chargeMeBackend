package api

import (
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

func convertGoogleCredentials(
	userId userDomain.UserId,
	credentials *schema.GoogleAuthCredentials,
) userDomain.GoogleCredentials {
	return userDomain.NewGoogleCredentials(
		userId,
		credentials.IdToken,
		credentials.AccessToken,
	)
}

func convertAppleCredentials(
	userId userDomain.UserId,
	credentials *schema.AppleAuthCredentials,
) userDomain.AppleCredentials {
	return userDomain.NewAppleCredentials(
		userId,
		credentials.AuthorizationCode,
		credentials.IdentityToken,
	)
}
