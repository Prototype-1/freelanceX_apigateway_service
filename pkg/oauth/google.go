package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
	"encoding/json"
)

var googleOAuth2Config = oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

type UserProfile struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func ExchangeAuthorizationCodeForToken(code string) (*UserProfile, error) {
	token, err := googleOAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("unable to exchange authorization code for token: %v", err)
	}
	client := googleOAuth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("unable to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userProfile UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&userProfile); err != nil {
		return nil, fmt.Errorf("unable to decode user info: %v", err)
	}

	return &userProfile, nil
}
