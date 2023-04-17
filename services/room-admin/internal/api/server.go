package api

import (
	features "musical-tables-api/services/room-admin/internal/application/features/creating_room"
	"musical-tables-api/services/room-admin/internal/persistence"
	"net/http"

	"github.com/labstack/echo/v4"
)

type echoHttpServer struct {
	echo *echo.Echo
}

func NewServer() *echoHttpServer {
	return &echoHttpServer{echo: echo.New()}
}

func (server *echoHttpServer) ConfigureEndpoints() {
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
	handler := features.NewCreateRoomHandler(persistence.NewInMemoryRoomRepository())
	result, err := handler.Handle(ctx, command)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
