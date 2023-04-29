package application_configurations

import (
	logging_behaviour "musical-tables-api/services/room-admin/internal/application/behaviours/logging"
	validation_behaviour "musical-tables-api/services/room-admin/internal/application/behaviours/validation"
	features "musical-tables-api/services/room-admin/internal/application/features/creating_room"
	"musical-tables-api/services/room-admin/internal/domain"

	"github.com/mehdihadeli/go-mediatr"
)

func ConfigMediatr(repository domain.RoomRepository) error {
	err := mediatr.RegisterRequestHandler[*features.CreateRoom, *features.CreateRoomResponseDto](features.NewCreateRoomHandler(repository))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestPipelineBehaviors(&logging_behaviour.RequestLoggerBehaviour{})
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestPipelineBehaviors(&validation_behaviour.RequestValidationBehaviour{})
	if err != nil {
		return err
	}

	return nil
}
