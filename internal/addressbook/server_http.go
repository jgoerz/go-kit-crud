package addressbook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/zerolog/log"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	createContactHandler := httptransport.NewServer(
		endpoints.CreateContactEP,
		decodeJSONToContactRequest,
		encodeResponseToJSON,
		httptransport.ServerErrorEncoder(createContactErrorEncoder),
	)

	readContactHandler := httptransport.NewServer(
		endpoints.ReadContactEP,
		decodeJSONToReadContactRequest,
		encodeResponseToJSON,
		httptransport.ServerErrorEncoder(readContactErrorEncoder),
	)

	updateContactHandler := httptransport.NewServer(
		endpoints.UpdateContactEP,
		decodeJSONToContactRequest,
		encodeResponseToJSON,
		httptransport.ServerErrorEncoder(updateContactErrorEncoder),
	)

	deleteContactHandler := httptransport.NewServer(
		endpoints.DeleteContactEP,
		decodeJSONToDeleteContactRequest,
		encodeResponseToJSON,
		httptransport.ServerErrorEncoder(deleteContactErrorEncoder),
	)

	router := chi.NewRouter()
	router.Method("POST", "/contacts", createContactHandler)
	router.Method("GET", "/contacts/{id}", readContactHandler)
	router.Method("PUT", "/contacts/{id}", updateContactHandler)
	router.Method("DELETE", "/contacts/{id}", deleteContactHandler)

	root := chi.NewRouter()
	root.Mount("/addressbook/v1", router)

	return root
}

// func contactRouter(service) http.Handler {
// 	r := chi.NewRouter()
// 	r.Post("/", srv.CreateContact)
// 	r.Get("/{id}", h.ReadContact)
// 	// FIXME TODO  Just needs to be added here in the server
// 	// r.Get("/", h.ReadAllContacts)
// 	r.Put("/{id}", h.UpdateContact)
// 	r.Delete("/{id}", h.DeleteContact)
//
// 	return r
// }

// https://github.com/go-kit/kit/blob/v0.12.0/transport/http/server.go#L139
func createContactErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	correlationID := CtxGetCorrelationID(ctx)                       // Needs to be in a GetLogger method somewhere
	log := log.With().Str("correlation_id", correlationID).Logger() // Needs to be in a GetLogger method somewhere
	log.Err(err).Msg("")

	data := []any{}
	payload := &StandardPayloadResponse{
		Data:          data,
		NextPageToken: 0,
		Errors: []StandardPayloadError{
			{
				// FIXME get error code support
				// Code:    "123",
				Message: err.Error(),
			},
		},
		CorrelationID: correlationID,
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func readContactErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	correlationID := CtxGetCorrelationID(ctx)
	log := log.With().Str("correlation_id", correlationID).Logger()
	log.Err(err).Msg("")

	data := []any{}
	payload := &StandardPayloadResponse{
		Data:          data,
		NextPageToken: 0,
		Errors: []StandardPayloadError{
			{
				// FIXME get error code support
				Message: err.Error(),
			},
		},
		CorrelationID: correlationID,
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func updateContactErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	correlationID := CtxGetCorrelationID(ctx)
	log := log.With().Str("correlation_id", correlationID).Logger()
	log.Err(err).Msg("")

	data := []any{}
	payload := &StandardPayloadResponse{
		Data:          data,
		NextPageToken: 0,
		Errors: []StandardPayloadError{
			{
				// FIXME get error code support
				Message: err.Error(),
			},
		},
		CorrelationID: correlationID,
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func deleteContactErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	correlationID := CtxGetCorrelationID(ctx)
	log := log.With().Str("correlation_id", correlationID).Logger()
	log.Err(err).Msg("")

	data := []any{}
	payload := &StandardPayloadResponse{
		Data:          data,
		NextPageToken: 0,
		Errors: []StandardPayloadError{
			{
				// FIXME get error code support
				Message: err.Error(),
			},
		},
		CorrelationID: correlationID,
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func decodeJSONToContactRequest(ctx context.Context, r *http.Request) (any, error) {
	log.Info().Msg("decodeContactRequest: Enter")
	var request ContactRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error().Msgf("decodeContactRequest: %v", err)
		return nil, err
	}
	log.Info().Msg("decodeContactRequest: Exit")
	return &request, nil
}

func decodeJSONToReadContactRequest(ctx context.Context, r *http.Request) (any, error) {
	log.Info().Msg("decodeReadContactRequest: Enter")

	rawID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		msg := fmt.Sprintf("invalid path param, expected integer got '%v'", rawID)
		log.Error().Msg(msg)
		return nil, fmt.Errorf(msg+": %w", ErrBadRequest)
	}
	log.Debug().Msgf("decodeReadContactRequest: id: '%v'", id)

	request := &ReadContactRequest{
		ID: int64(id),
	}
	log.Debug().Msgf("decodeReadContactRequest: request: '%v'", request)

	log.Info().Msg("decodeReadContactRequest: Exit")
	return request, nil
}

func decodeJSONToDeleteContactRequest(ctx context.Context, r *http.Request) (any, error) {
	log.Info().Msg("decodeDeleteContactRequest: Enter")

	rawID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		msg := fmt.Sprintf("invalid path param, expected integer got '%v'", rawID)
		log.Error().Msg(msg)
		return nil, fmt.Errorf(msg+": %w", ErrBadRequest)
	}
	log.Debug().Msgf("decodeDeleteContactRequest: id: '%v'", id)

	request := &DeleteContactRequest{
		ID: int64(id),
	}
	log.Debug().Msgf("decodeDeleteContactRequest: request: '%v'", request)

	log.Info().Msg("decodeDeleteContactRequest: Exit")
	return request, nil
}

func encodeResponseToJSON(ctx context.Context, w http.ResponseWriter, response any) error {
	log.Info().Msg("encodeResponse: Enter and Exit")
	return json.NewEncoder(w).Encode(response)
}
