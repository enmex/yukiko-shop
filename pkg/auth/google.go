package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthConfig struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
}

type GoogleAuth struct {
	OauthState string
	Oauth2     *oauth2.Config
}

type GoogleProfile struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

func NewGoogleAuth(googleAuthConfig *GoogleAuthConfig) *GoogleAuth {
	return &GoogleAuth{
		OauthState: "pseudo-random",
		Oauth2: &oauth2.Config{
			ClientID:     googleAuthConfig.ClientID,
			ClientSecret: googleAuthConfig.ClientSecret,
			RedirectURL:  googleAuthConfig.RedirectURL,
			Endpoint:     google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		},
	}
}

func (a *GoogleAuth) GetToken(code string) (*oauth2.Token, error) {
	token, err := a.Oauth2.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *GoogleAuth) GetProfileFromGoogle(token *oauth2.Token) (*GoogleProfile, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var resp GoogleProfile

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
