package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {
	uName, email, pwd, pwdConfirm := "", "", "", ""
	mux := http.NewServeMux()
	//Sign up
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10:There is empty data")
			return
		}
		if pwd == pwdConfirm {
			fmt.Fprintf(w, "Registration succesfull.")
		} else {
			fmt.Fprintf(w, "Password information must be same.")
		}
	})
	//Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		dbemail := "cumhurmesuttokalak@gmail.com"
		dbpwd := "123456"
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10:There is empty data")
			return
		}
		if email == dbemail && pwd == dbpwd {
			fmt.Fprintf(w, "Login succesful.")
		} else {
			fmt.Fprintf(w, "Login failed")
		}

	})
	http.ListenAndServe(":8080", mux)
}

/*
for key, value := range r.Form {
			fmt.Printf("%s =%s\n", key, value)
		}
*/
