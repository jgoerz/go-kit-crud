package addressbook

import (
	"context"
)

type Repository interface {
	CreateContact(ctx context.Context, input *ContactRequest) (output *ContactResponse, err error)
	ReadContact(ctx context.Context, input *ReadContactRequest) (output *ContactResponse, err error)
	UpdateContact(ctx context.Context, input *ContactRequest) (output *ContactResponse, err error)
	DeleteContact(ctx context.Context, input *DeleteContactRequest) (output *ContactResponse, err error)
}
