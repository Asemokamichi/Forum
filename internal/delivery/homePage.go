package delivery

import (
	"log"
	"net/http"
)

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.servErrors(w, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
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
