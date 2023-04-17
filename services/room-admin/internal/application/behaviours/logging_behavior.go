package behaviours

import (
	"context"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

type RequestLoggerBehaviour struct {
}

func (r *RequestLoggerBehaviour) Handle(ctx context.Context, request interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	log.Printf("received request: %v", request)

	response, err := next()
	if err != nil {
		return nil, err
	}

	log.Println("request executed")

	return response, nil
}
