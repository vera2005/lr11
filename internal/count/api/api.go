package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	maxNum int

	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxNum int, uc Usecase) *Server {
	api := Server{
		maxNum: maxNum,
		uc:     uc,
	}

	api.server = echo.New()
	api.server.GET("/count", api.GetCount)
	api.server.POST("/count", api.PostCount)
	api.server.PUT("/count", api.PutCount)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
