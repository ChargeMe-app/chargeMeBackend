package user

type AppleCredentials struct {
	userId            UserId
	authorizationCode string
	identityToken     string
}

func NewAppleCredentials(
	userId UserId,
	authorizationCode string,
	identityToken string,
) AppleCredentials {
	return AppleCredentials{
		userId:            userId,
		authorizationCode: authorizationCode,
		identityToken:     identityToken,
	}
}

func (a *AppleCredentials) GetUserId() UserId {
	return a.userId
}

func (a *AppleCredentials) GetAuthorizationCode() string {
	return a.authorizationCode
}

func (a *AppleCredentials) GetIdentityToken() string {
	return a.identityToken
}

type GoogleCredentials struct {
	userId      UserId
	idToken     string
	accessToken string
}

func NewGoogleCredentials(
	userId UserId,
	IdToken string,
	AccessToken string,
) GoogleCredentials {
	return GoogleCredentials{
		userId:      userId,
		idToken:     IdToken,
		accessToken: AccessToken,
	}
}

func (g *GoogleCredentials) GetUserId() UserId {
	return g.userId
}

func (g *GoogleCredentials) GetIdToken() string {
	return g.idToken
}

func (g *GoogleCredentials) GetAccessToken() string {
	return g.accessToken
}
