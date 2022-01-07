package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {

	loadEnv()
	initOauthGoogle()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pathFile := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(pathFile)
		panicIfError(err)

		err = tmpl.Execute(w, false)

	})

	http.HandleFunc("/auth/login", loginHandler)
	http.HandleFunc("/oauth2/google", googleOauthCbHandler)

	fmt.Println("server is run at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
