package main

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfigGoogle = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  getOauthRedirectUrl(),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthState      = ""
	devRedirectUri  = "http://localhost:8000/oauth2/google"
	prodRedirectUri = "https://nekonako.space/oauth2/google"
)

func getEnv(key string) string {
	return os.Getenv(key)
}

func loadEnv(filename ...string) {
	err := godotenv.Load(filename...)
	panicIfError(err)
}

func initOauthGoogle() {
	oauthConfigGoogle.ClientID = getEnv("GOOGLE_OAUTH_CLIENT_ID")
	oauthConfigGoogle.ClientSecret = getEnv("GOOGLE_OAUTH_CLIENT_SECRET")
	oauthState = getEnv("OAUTH_STATE")
}

func getOauthRedirectUrl() string {
	if getEnv("ENV") == "production" {
		return prodRedirectUri
	} else {
		return devRedirectUri
	}
}
