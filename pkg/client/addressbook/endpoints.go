package addressbook

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"
)

type Endpoints struct {
	CreateContactEP endpoint.Endpoint
	ReadContactEP   endpoint.Endpoint
	ListContactsEP  endpoint.Endpoint
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

func (e Endpoints) ListContacts(ctx context.Context, requ *ListContactsRequest) (resp *ListContactsResponse, err error) {
	raw, err := e.ListContactsEP(ctx, requ)
	if err != nil {
		return nil, err
	}
	resp, ok := raw.(*ListContactsResponse)
	if !ok {
		err = fmt.Errorf("failed interface conversion;  expected *ListContactsResponse, got '%T'", raw)
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
