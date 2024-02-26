package main

import (
	"encoding/json"
	"fmt"
	"gading-go-challenge/database"
	"html/template"
	"net/http"
	"path"
)

func main() {

	PORT := ":9090"

	http.HandleFunc("/", HandleMainPage)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/login-error", HandleLoginError)
	http.HandleFunc("/login-success", HandleLoginSuccess)

	http.ListenAndServe(PORT, nil)
}

func HandleLoginError(w http.ResponseWriter, r *http.Request) {

	errorTargetQuery := r.URL.Query().Get("errorTarget")
	errorMessage := ""

	for keyValue, value := range database.ListOfErrors {
		if keyValue == errorTargetQuery {
			errorMessage = value
			break
		}
	}

	filepath := path.Join("views", "errorpage.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, errorMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleLoginSuccess(w http.ResponseWriter, r *http.Request) {

	emailParam := r.URL.Query().Get("email")
	biodata := map[string]any{}

	for _, value := range database.ListAccounts {
		if emailParam == value.Email {
			biodata = map[string]any{
				"email":   value.Email,
				"address": value.Address,
				"job":     value.Job,
				"reason":  value.Reason,
			}
			break
		}
	}

	filepath := path.Join("views", "biodata.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, biodata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleMainPage(w http.ResponseWriter, r *http.Request) {

	filepath := path.Join("views", "index.html")

	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, database.ListAccounts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		isEmailFound := false

		var loginRequest database.LoginRequest

		errDecode := json.NewDecoder(r.Body).Decode(&loginRequest)

		if errDecode != nil {
			http.Error(w, errDecode.Error(), http.StatusInternalServerError)
		}

		for _, value := range database.ListAccounts {
			if value.Email == loginRequest.Email {
				isEmailFound = true
				if value.Password == loginRequest.Password {
					successUrl := fmt.Sprintf("/login-success?email=%s", value.Email)
					http.Redirect(w, r, successUrl, http.StatusSeeOther)
					return
				} else {
					http.Redirect(w, r, "/login-error?errorTarget=ErrorPassword", http.StatusSeeOther)
					return
				}
			}
		}

		if !isEmailFound {
			http.Redirect(w, r, "/login-error?errorTarget=ErrorEmail", http.StatusSeeOther)
			return
		}

	}

	http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
}
