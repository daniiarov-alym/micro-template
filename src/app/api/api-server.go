package apiserver

import (
	"github.com/daniiarov-alym/micro-template/src/app/service"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type apiServer struct {
	router     *echo.Echo
	cmdService service.Service
}

func NewServer(cmdService service.Service) (s *apiServer) {
	if cmdService == nil {
		logger.Fatal("cmdService can not be nil!")
	}
	s = &apiServer{echo.New(), cmdService}
	s.registerApi()
	return
}

func (s *apiServer) Start(port string) error {
	return s.router.Start(":" + port)
}

func (s *apiServer) registerApi() {
	s.router.GET("/health-ping", s.healthPing)
	// configure routing further here
}
