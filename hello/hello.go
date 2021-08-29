package hello

import (
	"fmt"

	"github.com/TeamChii/hello-lambda/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Servicer interface {
	HelloService(c echo.Context) (*HelloResponse, error)
}

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {

	return &Service{logger: logger}
}

func (s *Service) HelloService(c echo.Context, req *HelloRequest) (*HelloResponse, error) {
	s.logger.Info("process HelloService")
	return &HelloResponse{
		HeaderResp: common.HeaderResp{
			Code:    common.StatusCdSuccess,
			Message: common.StatusDescMap[common.StatusCdSuccess],
		}, Msg: fmt.Sprintf("hello %v", req.Name),
	}, nil
}
