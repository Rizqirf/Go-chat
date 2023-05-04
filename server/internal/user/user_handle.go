package user

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
)

var (lock = sync.Mutex{})

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func(h *Handler) CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var user CreateUserReq

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := h.Service.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) Login(c echo.Context) error{
	lock.Lock()
	defer lock.Unlock()

	var user LoginUserReq

	if err := c.Bind(&user); err != nil {
		
		return c.JSON(http.StatusBadRequest, err)
	}

	u, err := h.Service.Login(c.Request().Context(), &user)
	if err != nil {
		
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = u.accessToken
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Logout(c echo.Context) error{
	lock.Lock()
	defer lock.Unlock()

	res := map[string]string{}

	res["message"] = "Logout Successful"



	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, res)
}