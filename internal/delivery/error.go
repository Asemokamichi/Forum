package delivery

import (
	"net/http"
)

func (h *Handler) servErrors(w http.ResponseWriter, code int, ErrorText string) {
	w.WriteHeader(code)
	if err := h.tmpl.ExecuteTemplate(w, "error.html", struct {
		code      int
		codeText  string
		ErrorText string
	}{
		code:     code,
		codeText: http.StatusText(code),
		ErrorText: ErrorText,
	}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
