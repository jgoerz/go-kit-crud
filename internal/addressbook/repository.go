package addressbook

import (
	"context"

	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"
)

type Repository interface {
	CreateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error)
	ReadContact(ctx context.Context, input *clientapi.ReadContactRequest) (output *clientapi.ContactResponse, err error)
	UpdateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error)
	DeleteContact(ctx context.Context, input *clientapi.DeleteContactRequest) (output *clientapi.ContactResponse, err error)
}
