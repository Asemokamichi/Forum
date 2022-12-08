package delivery

import (
	"fmt"
	"net/http"
)

func (h *Handler) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("*")
		h.servErrors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == http.MethodGet {
		if err := h.tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			h.servErrors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
}
