package addressbook

import (
	"context"

	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"

	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"
)

func MakeEndpoints(srv Service) clientapi.Endpoints {
	return clientapi.Endpoints{
		CreateContactEP: makeCreateContactEP(srv),
		ReadContactEP:   makeReadContactEP(srv),
		UpdateContactEP: makeUpdateContactEP(srv),
		DeleteContactEP: makeDeleteContactEP(srv),
	}
}

func makeCreateContactEP(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		log.Info().Msg("CreateContact Endpoint: Enter")
		// correlationID := CtxGetCorrelationID(ctx)
		// data := []any{}
		// TODO logger?

		requ := request.(*clientapi.ContactRequest)
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

		requ := request.(*clientapi.ReadContactRequest)
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

		requ := request.(*clientapi.ContactRequest)
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
		requ := request.(*clientapi.DeleteContactRequest)
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
