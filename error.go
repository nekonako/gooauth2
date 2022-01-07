package main

import (
	"fmt"
	"net/http"
)

func panicIfError(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func redirectIfErr(err interface{}, path string, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, path, http.StatusTemporaryRedirect)
		return
	}
}
