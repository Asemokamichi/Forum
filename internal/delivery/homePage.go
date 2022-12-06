package delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Asemokamichi/Forum/internal/model"
)

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.servErrors(w, http.StatusNotFound)
		return
	}
	token, err := r.Cookie("session-token")
	var user model.User
	fmt.Println(token.Value, err)
	if err == nil {
		user, err = h.Service.GetUserSession(token.Value)
		if err != nil {
			h.servErrors(w, http.StatusInternalServerError)
			return
		}
	} else if err != http.ErrNoCookie {
		h.servErrors(w, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "index.html", user); err != nil {
			h.servErrors(w, http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm auth signUp", err)
			h.servErrors(w, http.StatusInternalServerError)
		}

		if _, ok := r.Form["signIn"]; ok {
			http.Redirect(w, r, "/signIn", http.StatusSeeOther)
		} else if _, ok := r.Form["signUp"]; ok {
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
		}
	}
}
