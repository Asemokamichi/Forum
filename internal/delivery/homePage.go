package delivery

import (
	"net/http"
)

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.servErrors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	// token, err := r.Cookie("session-token")
	// var user model.User
	// fmt.Println(token.Value, err)
	// if err == nil {
	// 	user, err = h.Service.GetUserSession(token.Value)
	// 	if err != nil {
	// 		h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	// 		return
	// 	}
	// } else if err != http.ErrNoCookie {
	// 	h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	// 	return
	// }

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			h.servErrors(w, http.StatusInternalServerError, "ParseForm auth signUp")
			return
		}

		if _, ok := r.Form["signIn"]; ok {
			http.Redirect(w, r, "/signIn", http.StatusSeeOther)
			return
		} else if _, ok := r.Form["signUp"]; ok {
			http.Redirect(w, r, "/signUp", http.StatusSeeOther)
			return
		}
	}
}
