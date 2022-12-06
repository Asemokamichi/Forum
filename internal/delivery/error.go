package delivery

import (
	"net/http"
)

func (h *Handler) servErrors(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	if err := h.tmpl.ExecuteTemplate(w, "error.html", struct {
		code     int
		codeText string
	}{
		code:     code,
		codeText: http.StatusText(code),
	}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
