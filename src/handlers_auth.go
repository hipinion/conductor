package conductor

import (
	"fmt"
	"net/http"
)

func authLogoutHandler(w http.ResponseWriter, r *http.Request) {
	Sess, _ := GetSessionUser(w, r)
	Sess.Remove()

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func authRegisterHandler(w http.ResponseWriter, r *http.Request) {

	RefreshTemplates()
	site, err := GetSite(r)
	if err != nil {

	}

	Sess, AuthUser := GetSessionUser(w, r)
	fmt.Println(Sess, AuthUser)
	Page := struct {
		Username string
		Email    string
	}{
		"",
		"",
	}

	PL := Payload{
		Title: "Register",
		Page:  Page,
		Site:  site,
		User:  AuthUser,
	}

	if site.InvitationsActive && !site.RegistrationOpen {
		Templates.ExecuteTemplate(w, "register_invitation.html", PL)
		return
	}
	Templates.ExecuteTemplate(w, "register.html", PL)
}

func authRegisterProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	site, err := GetSite(r)
	if err != nil {

	}

	fmt.Println(site)
	var hasErrors bool

	Page := struct {
		Username string
		Email    string
	}{
		r.PostFormValue("user_name"),
		r.PostFormValue("user_email"),
	}

	PL := Payload{
		Title: "Register",
		Page:  Page,
	}

	p := r.PostFormValue("password")
	pr := r.PostFormValue("password_repeat")

	if p != pr {
		hasErrors = true
		PL.Errors = append(PL.Errors, "Passwords do not match")
	} else {

		hash, salt, _ := Hash(p, "")

		U := NewUser()
		U.Name = r.PostFormValue("user_name")
		U.Phone = ""
		U.Email = r.PostFormValue("user_email")
		U.Salt = salt
		U.Hash = hash
		err := U.Set()
		if err != nil {
			hasErrors = true
			PL.Errors = append(PL.Errors, err.Error())
		}
	}

	if hasErrors {
		Templates.ExecuteTemplate(w, "register.html", PL)
	} else {
		http.Redirect(w, r, "/forums", http.StatusSeeOther)
	}
}

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	site, _ := GetSite(r)
	PL := Payload{
		Title: "Login",
		Site:  site,
	}
	Templates.ExecuteTemplate(w, "login.html", PL)
}

func authLoginProcessHandler(w http.ResponseWriter, r *http.Request) {

	Sess, _ := GetSessionUser(w, r)

	username := r.PostFormValue("user_name")
	password := r.PostFormValue("password")

	U := NewUsers()
	U.Name = username
	U.Get()

	if len(U.Users) > 0 {

		valid, _ := ValidatePassword(U.Users[0].Password, (password + U.Users[0].Salt))
		if valid {
			Sess.SetUser(U.Users[0].ID, Sess.ID)
			http.Redirect(w, r, "/forums", http.StatusSeeOther)
		} else {
			fmt.Fprintln(w, "No", U)
		}
	}

}

func authRecoverHandler(w http.ResponseWriter, r *http.Request) {

}

func authRecoverProcessHandler(w http.ResponseWriter, r *http.Request) {

}
