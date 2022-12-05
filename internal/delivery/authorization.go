package delivery

import (
	"log"
	"net/http"

	"github.com/Asemokamichi/Forum/internal/model"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signUp" {
		log.Fatal("Error Url auth signUp", r.URL.Path)
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signUp.html", nil); err != nil {
			log.Fatal("ExecuteTemplate auth signUp", err)
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatal("ParseForm auth signUp", err)
		}

		username, ok := r.Form["username"]
		if !ok {
			w.Write([]byte("Sign Up: Parse Form: username field not found"))
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		email, ok := r.Form["email"]
		if !ok {
			w.Write([]byte("Sign Up: Parse Form: email field not found"))
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		password, ok := r.Form["password"]
		if !ok {
			w.Write([]byte("Sign Up: Parse Form: password field not found"))
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		confpassword, ok := r.Form["confirmPassword"]
		if !ok {
			w.Write([]byte("Sign Up: Parse Form: password field not found"))
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		if confpassword[0] != password[0] {
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		user := model.User{
			Username: username[0],
			Email:    email[0],
			Password: password[0],
		}

		if err := h.Service.CreateUser(user); err != nil {
			log.Fatal(err)
		}

	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signIn" {
		h.errorPage(w, 500)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signIn.html", nil); err != nil {
			h.errorPage(w, 500)
			return
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatal("ParseForm auth signIn", err)
		}

		password, ok := r.Form["password"]
		if !ok {
			log.Fatal("ParseForm auth signIn password")
		}

		username, ok := r.Form["username"]
		if !ok {
			log.Fatal("ParseForm auth signIn username")
		}

		user := model.User{
			Username: username[0],
			Password: password[0],
		}

		user, err := h.Service.GetUser(user)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
	}
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		h.errorPage(w, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "register.html", nil); err != nil {
			h.errorPage(w, 500)
		}
	}

}
