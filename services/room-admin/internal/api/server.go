package api

import (
	"errors"
	"fmt"
	application_configurations "musical-tables-api/services/room-admin/internal/application/configurations"
	features "musical-tables-api/services/room-admin/internal/application/features/creating_room"
	"musical-tables-api/services/room-admin/internal/mongo_persistence"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	repo, err := mongo_persistence.NewMongoDbRepository()

	if err != nil {
		panic(err)
	}

	application_configurations.ConfigMediatr(repo)
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
		return hadnleError(err, c)
	}

	return c.JSON(http.StatusCreated, result)
}

func hadnleError(err error, c echo.Context) error {
	for {
		unwrapErr := errors.Unwrap(err)
		if unwrapErr == nil {
			break
		}
		err = unwrapErr
	}

	validationErr, ok := err.(validator.ValidationErrors)
	if ok {
		var failedValidations []string
		for _, vErr := range validationErr {
			failedValidations = append(failedValidations, fmt.Sprintf("'%s' has a value of '%v' which does not satisfy '%s'.\n", vErr.Field(), vErr.Value(), vErr.Tag()))
		}
		return c.JSON(http.StatusBadRequest, failedValidations)
	}

	return err
}
