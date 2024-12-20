package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}

	api.server = echo.New()
	api.server.POST("/reg", api.signUp)
	api.server.POST("/aui", api.signIn)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
