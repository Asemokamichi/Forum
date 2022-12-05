package delivery

import (
	"net/http"
	"text/template"

	"github.com/Asemokamichi/Forum/internal/service"
)

type Handler struct {
	tmpl    *template.Template
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		tmpl:    template.Must(template.ParseGlob("internal/template/html/*.html")),
		Service: service,
	}
}

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.homePage)
	mux.HandleFunc("/signUp", h.signUp)
	mux.HandleFunc("/signIn", h.signIn)
	mux.HandleFunc("/registration", h.registration)
}
