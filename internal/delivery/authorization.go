package delivery

import (
	"log"
	"net/http"

	"github.com/Asemokamichi/Forum/internal/model"
)

var answer string

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signUp" {
		log.Println("Error Url auth signUp", r.URL.Path)
		h.servErrors(w, http.StatusNotFound)
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signUp.html", answer); err != nil {
			log.Println("ExecuteTemplate auth signUp", err)
			h.servErrors(w, http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm auth signUp", err)
			h.servErrors(w, http.StatusInternalServerError)
		}

		username, ok := r.Form["username"]
		if !ok {
			answer = "Sign Up: Parse Form: username field not found"
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}

		email, ok := r.Form["email"]
		if !ok || len(email[0]) == 0 {
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
			log.Println(err)
			h.servErrors(w, http.StatusInternalServerError)
		}
		answer = ""
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signIn" {
		h.servErrors(w, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signIn.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm auth signIn", err)
			h.servErrors(w, http.StatusInternalServerError)
		}

		password, ok := r.Form["password"]
		if !ok {
			log.Println("ParseForm auth signIn password")
		}

		username, ok := r.Form["username"]
		if !ok {
			log.Println("ParseForm auth signIn username")
		}

		user := model.User{
			Username: username[0],
			Password: password[0],
		}

		user, err := h.Service.GetUser(user)
		if err != nil {
			log.Println(err)
			h.servErrors(w, http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
	}
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		h.servErrors(w, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "register.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError)
		}
	}
}
