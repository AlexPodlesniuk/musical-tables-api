package validation_behaviour

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/mehdihadeli/go-mediatr"
)

type RequestValidationBehaviour struct {
}

func (r *RequestValidationBehaviour) Handle(ctx context.Context, request interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {

	v := validator.New()

	err := v.Struct(request)
	if err != nil {
		return nil, err
	}

	return next()
}
