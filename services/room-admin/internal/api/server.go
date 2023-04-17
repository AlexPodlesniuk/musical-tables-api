package api

import (
	application_configurations "musical-tables-api/services/room-admin/internal/application/configurations"
	features "musical-tables-api/services/room-admin/internal/application/features/creating_room"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

type echoHttpServer struct {
	echo *echo.Echo
}

func NewServer() *echoHttpServer {
	return &echoHttpServer{echo: echo.New()}
}

func (server *echoHttpServer) ConfigureEndpoints() {
	application_configurations.ConfigMediatr()
	server.echo.POST("/api/v1/rooms", handleCreateRoom)
}

func (server *echoHttpServer) Start() {
	server.echo.Logger.Fatal(server.echo.Start(":9090"))
}

func handleCreateRoom(c echo.Context) error {
	ctx := c.Request().Context()
	request := &features.CreateRoomRequestDto{}

	if err := c.Bind(request); err != nil {
		return err
	}

	command := features.NewCreateRoom(request.Name)
	result, err := mediatr.Send[*features.CreateRoom, *features.CreateRoomResponseDto](ctx, command)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
