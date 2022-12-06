package delivery

import (
	"fmt"
	"net/http"

	"github.com/Asemokamichi/Forum/internal/model"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signUp" {
		h.servErrors(w, http.StatusNotFound, "Error Url auth signUp")
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signUp.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError, "ExecuteTemplate auth signUp")
			return
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			h.servErrors(w, http.StatusInternalServerError, "ParseForm auth signUp")
			return
		}

		username, ok := r.Form["username"]
		if !ok {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: username field not found")
			return
		}

		email, ok := r.Form["email"]
		if !ok || len(email[0]) == 0 {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: username field not found")
			return
		}

		password, ok := r.Form["password"]
		if !ok {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: password field not found")
			return
		}

		confpassword, ok := r.Form["confirmPassword"]
		if !ok {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: password field not found")
			return
		}

		if confpassword[0] != password[0] {
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		}

		user := model.User{
			Username: username[0],
			Email:    email[0],
			Password: password[0],
		}

		if err := h.Service.CreateUser(user); err != nil {
			h.servErrors(w, http.StatusInternalServerError, fmt.Sprint(err))
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signIn" {
		h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "signIn.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			h.servErrors(w, http.StatusInternalServerError, "ParseForm auth signIn")
			return
		}

		username, ok := r.Form["username"]
		if !ok {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: username field not found")
			return
		}

		password, ok := r.Form["password"]
		if !ok {
			h.servErrors(w, http.StatusBadRequest, "Sign Up: Parse Form: password field not found")
			return
		}

		user, err := h.Service.CreateSession(username[0], password[0])
		if err != nil {
			h.servErrors(w, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session-token",
			Value:   user.UUID,
			Expires: user.ExpDate,
			Path:    "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *Handler) registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		h.servErrors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "register.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
}
