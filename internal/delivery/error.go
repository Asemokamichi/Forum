package delivery

import "net/http"

func (h *Handler) errorPage(w http.ResponseWriter, code int) {
	w.WriteHeader(code)

	data := struct {
		Status  int
		Message string
	}{
		Status:  code,
		Message: http.StatusText(code),
	}

	if err := h.tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, http.StatusText(code), http.StatusInternalServerError)
	}
}
