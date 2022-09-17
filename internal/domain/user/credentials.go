package user

type AppleCredentials struct {
	userId            UserID
	authorizationCode string
	identityToken     string
}

func NewAppleCredentials(
	userId UserID,
	authorizationCode string,
	identityToken string,
) AppleCredentials {
	return AppleCredentials{
		userId:            userId,
		authorizationCode: authorizationCode,
		identityToken:     identityToken,
	}
}

func (a *AppleCredentials) GetUserId() UserID {
	return a.userId
}

func (a *AppleCredentials) GetAuthorizationCode() string {
	return a.authorizationCode
}

func (a *AppleCredentials) GetIdentityToken() string {
	return a.identityToken
}

type GoogleCredentials struct {
	userId      UserID
	idToken     string
	accessToken string
}

func NewGoogleCredentials(
	userId UserID,
	IdToken string,
	AccessToken string,
) GoogleCredentials {
	return GoogleCredentials{
		userId:      userId,
		idToken:     IdToken,
		accessToken: AccessToken,
	}
}

func (g *GoogleCredentials) GetUserId() UserID {
	return g.userId
}

func (g *GoogleCredentials) GetIdToken() string {
	return g.idToken
}

func (g *GoogleCredentials) GetAccessToken() string {
	return g.accessToken
}
