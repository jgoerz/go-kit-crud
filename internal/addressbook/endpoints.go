package addressbook

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"
)

func MakeEndpoints(srv Service) Endpoints {
	return Endpoints{
		CreateContactEP: makeCreateContactEP(srv),
		ReadContactEP:   makeReadContactEP(srv),
		UpdateContactEP: makeUpdateContactEP(srv),
		DeleteContactEP: makeDeleteContactEP(srv),
	}
}

type Endpoints struct {
	CreateContactEP endpoint.Endpoint
	ReadContactEP   endpoint.Endpoint
	UpdateContactEP endpoint.Endpoint
	DeleteContactEP endpoint.Endpoint
}

func (e Endpoints) CreateContact(ctx context.Context, requ *ContactRequest) (resp *ContactResponse, err error) {
	raw, err := e.CreateContactEP(ctx, requ)
	if err != nil {
		return nil, err
	}
	resp, ok := raw.(*ContactResponse)
	if !ok {
		err = fmt.Errorf("failed interface conversion;  expected *ContactResponse, got '%T'", raw)
		log.Err(err).Msg("")
		return nil, err
	}
	return resp, err
}

func (e Endpoints) ReadContact(ctx context.Context, requ *ReadContactRequest) (resp *ContactResponse, err error) {
	raw, err := e.ReadContactEP(ctx, requ)
	if err != nil {
		return nil, err
	}
	resp, ok := raw.(*ContactResponse)
	if !ok {
		err = fmt.Errorf("failed interface conversion;  expected *ContactResponse, got '%T'", raw)
		log.Err(err).Msg("")
		return nil, err
	}
	return resp, err
}

func (e Endpoints) UpdateContact(ctx context.Context, requ *ContactRequest) (resp *ContactResponse, err error) {
	raw, err := e.UpdateContactEP(ctx, requ)
	if err != nil {
		return nil, err
	}
	resp, ok := raw.(*ContactResponse)
	if !ok {
		err = fmt.Errorf("failed interface conversion;  expected *ContactResponse, got '%T'", raw)
		log.Err(err).Msg("")
		return nil, err
	}
	return resp, err
}

func (e Endpoints) DeleteContact(ctx context.Context, requ *DeleteContactRequest) (resp *ContactResponse, err error) {
	raw, err := e.DeleteContactEP(ctx, requ)
	if err != nil {
		return nil, err
	}
	resp, ok := raw.(*ContactResponse)
	if !ok {
		err = fmt.Errorf("failed interface conversion;  expected *ContactResponse, got '%T'", raw)
		log.Err(err).Msg("")
		return nil, err
	}
	return resp, err
}

func makeCreateContactEP(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		log.Info().Msg("CreateContact Endpoint: Enter")
		// correlationID := CtxGetCorrelationID(ctx)
		// data := []any{}
		// TODO logger?

		requ := request.(*ContactRequest)
		contact, err := srv.CreateContact(ctx, requ)
		if err != nil {
			log.Err(err).Msg("")
			// https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L14
			// Handled by errorHandler and errorEncoder
			return nil, err
		}

		log.Info().Msg("CreateContact Endpoint: Exit")
		return contact, err
	}
}

func makeReadContactEP(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		correlationID := CtxGetCorrelationID(ctx)
		log := log.With().Str("correlation_id", correlationID).Logger() // Needs to be in a GetLogger method somewhere
		log.Info().Msg("ReadContact Endpoint: Enter")

		requ := request.(*ReadContactRequest)
		contact, err := srv.ReadContact(ctx, requ)
		if err != nil {
			log.Err(err).Msg("")
			// https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L14
			// Handled by errorHandler and errorEncoder
			return nil, err
		}

		log.Info().Msg("ReadContact Endpoint: Exit")
		return contact, err
	}
}

func makeUpdateContactEP(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		correlationID := CtxGetCorrelationID(ctx)
		log := log.With().Str("correlation_id", correlationID).Logger() // Needs to be in a GetLogger method somewhere
		log.Info().Msg("UpdateContact Endpoint: Enter")
		// TODO logger?

		requ := request.(*ContactRequest)
		contact, err := srv.UpdateContact(ctx, requ)
		if err != nil {
			log.Err(err).Msg("")
			// https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L14
			// Handled by errorHandler and errorEncoder
			return nil, err
		}

		log.Info().Msg("UpdateContact Endpoint: Exit")
		return contact, err
	}
}

func makeDeleteContactEP(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		correlationID := CtxGetCorrelationID(ctx)
		log := log.With().Str("correlation_id", correlationID).Logger() // Needs to be in a GetLogger method somewhere
		log.Info().Msg("DeleteContact Endpoint: Enter")

		// TODO logger?
		requ := request.(*DeleteContactRequest)
		contact, err := srv.DeleteContact(ctx, requ)
		if err != nil {
			log.Err(err).Msg("")
			// https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L14
			// Handled by errorHandler and errorEncoder
			return nil, err
		}

		log.Info().Msg("DeleteContact Endpoint: Exit")
		return contact, err
	}
}
