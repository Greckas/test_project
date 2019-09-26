package handler

import (
	"awesomeProject/model"
	"awesomeProject/response"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	adminAccessLvl = "admin"
)

type Admin struct {
	store AdminStore
}

type AdminStore interface {
	Login(username string) (*model.User, error)
}

func NewAdmin(s AdminStore) *Admin {
	return &Admin{store: s}
}

func (h *Admin) Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.tmpl.html", nil)
}

func (h *Admin) LoginGo(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.store.Login(username)
	if err != nil {
		//todo add log
		return response.HandleModelError(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// TODO: Properly handle error
		log.Error(err)
		return response.ErrUnauthorized
	}

	if user.AccessLevel != adminAccessLvl {
		//todo add log
		return response.ErrUnauthorized.SetMessage("admin access needed")
	}

	return c.Redirect(http.StatusFound, "/router/v1/admin/home")
}

func (h *Admin) Logout(c echo.Context) error {
	c.Set("token", "")
	return c.Redirect(http.StatusFound, "login")
}

func (h *Admin) Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.tmpl.html", nil)
}

func (h *Admin) Dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.tmpl.html", nil)
}
