package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"golang.org/x/oauth2"
)

type User struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	URL, err := url.Parse(oauthConfigGoogle.Endpoint.AuthURL)
	panicIfError(err)

	parameters := url.Values{}
	parameters.Add("client_id", oauthConfigGoogle.ClientID)
	parameters.Add("scope", strings.Join(oauthConfigGoogle.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfigGoogle.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthState)

	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func googleOauthCbHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

	if state != oauthState {
		fmt.Println("invalid oauth state, expected ", oauthState, " but got", state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		reason := r.FormValue("error_reason")
		fmt.Println(reason)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfigGoogle.Exchange(oauth2.NoContext, code)
		redirectIfErr(err, "/", w, r)

		res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		redirectIfErr(err, "/", w, r)
		defer res.Body.Close()

		response, err := ioutil.ReadAll(res.Body)
		panicIfError(err)

		var user User
		err = json.Unmarshal(response, &user)
		panicIfError(err)

		pathFile := path.Join("views", "success.html")
		tmpl, err := template.ParseFiles(pathFile)
		panicIfError(err)

		err = tmpl.Execute(w, user)
		panicIfError(err)
		return
	}
}
