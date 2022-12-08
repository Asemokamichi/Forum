package delivery

import (
	"fmt"
	"net/http"

	"github.com/Asemokamichi/Forum/internal/model"
)

func (h *Handler) setCookie(w http.ResponseWriter, r *http.Request, session model.Session) {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   session.UUID,
		Expires: session.ExpDate,
		Path:    "/",
	}

	http.SetCookie(w, cookie)
}

func (h *Handler) checkCookie(w http.ResponseWriter, r *http.Request, session *model.Session) (*http.Cookie, error) {
	cookie, err := r.Cookie("token")

	if err == http.ErrNoCookie {
		return nil, err
	}

	if cookie.Value == "" {
		return nil, fmt.Errorf("Check cookie: cookie.Value is empty")
	}

	if session != nil && cookie.Value != session.UUID {
		return nil, fmt.Errorf("Check cookie: UUID doesn't match")
	}

	return cookie, nil
}

func (h *Handler) getUserIDBySession(w http.ResponseWriter, r *http.Request, session model.Session) (*model.User, error) {
	cookie, err := h.checkCookie(w, r, &session)
	if err != nil {
		return nil, err
	}

	user, err := h.Service.GetUserSession(cookie.Value)
	if err != nil {
		return nil, err
	}

	return user, err
}
