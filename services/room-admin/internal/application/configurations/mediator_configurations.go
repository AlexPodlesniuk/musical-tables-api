package application_configurations

import (
	"musical-tables-api/services/room-admin/internal/application/behaviours"
	features "musical-tables-api/services/room-admin/internal/application/features/creating_room"
	"musical-tables-api/services/room-admin/internal/persistence"

	"github.com/mehdihadeli/go-mediatr"
)

func ConfigMediatr() error {
	err := mediatr.RegisterRequestHandler[*features.CreateRoom, *features.CreateRoomResponseDto](features.NewCreateRoomHandler(persistence.NewInMemoryRoomRepository()))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestPipelineBehaviors(&behaviours.RequestLoggerBehaviour{})
	if err != nil {
		return err
	}

	return nil
}
