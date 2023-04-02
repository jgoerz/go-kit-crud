package addressbook

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"

	"github.com/rs/zerolog/log"
)

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		lock:          &sync.RWMutex{},
		contacts:      make(map[int64]*Contact, 0),
		nextContactID: 1,
	}
}

type inMemoryRepository struct {
	lock          *sync.RWMutex
	contacts      map[int64]*Contact
	nextContactID int64
}

func (r *inMemoryRepository) CreateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msg("inMemoryRepository: CreateContact: Enter")
	r.lock.Lock()
	defer r.lock.Unlock()

	now := time.Now().UTC().Format(time.RFC3339)
	contact := &Contact{
		ID:         r.nextContactID,
		TenantID:   input.TenantID,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Active:     input.Active,
		Address:    input.Address,
		SomeSecret: input.SomeSecret,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	_, exists := r.contacts[input.ID]
	if exists {
		return nil, fmt.Errorf("duplicate entry: '%v' %w", contact.ID, ErrBadRequest)
	} else {
		r.nextContactID++
	}

	r.contacts[contact.ID] = contact

	log.Debug().Msg("inMemoryRepository: CreateContact: Exit")
	return &clientapi.ContactResponse{
		ID:         contact.ID,
		TenantID:   contact.TenantID,
		FirstName:  contact.FirstName,
		LastName:   contact.LastName,
		Active:     contact.Active,
		Address:    contact.Address,
		SomeSecret: contact.SomeSecret,
		CreatedAt:  contact.CreatedAt,
		UpdatedAt:  contact.UpdatedAt,
	}, nil
}

func (r *inMemoryRepository) ReadContact(ctx context.Context, input *clientapi.ReadContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msg("inMemoryRepository: ReadContact: Enter")
	r.lock.RLock()
	defer r.lock.RUnlock()

	if input == nil {
		return nil, ErrBadRequest
	}
	log.Debug().Msgf("inMemoryRepository: ReadContact: input: %v", input)

	contact, exists := r.contacts[input.ID]
	if !exists || input.ID == 0 {
		log.Error().Msg("inMemoryRepository: ReadContact: not found")
		return nil, ErrNotFound
	}
	log.Debug().Msgf("inMemoryRepository: ReadContact: contact: %v", contact)

	log.Debug().Msg("inMemoryRepository: ReadContact: Exit")
	return &clientapi.ContactResponse{
		ID:         contact.ID,
		TenantID:   contact.TenantID,
		FirstName:  contact.FirstName,
		LastName:   contact.LastName,
		Active:     contact.Active,
		Address:    contact.Address,
		SomeSecret: contact.SomeSecret,
		CreatedAt:  contact.CreatedAt,
		UpdatedAt:  contact.UpdatedAt,
	}, nil
}

func (r *inMemoryRepository) ListContacts(ctx context.Context, input *clientapi.ListContactsRequest) (output *clientapi.ListContactsResponse, err error) {
	log.Debug().Msg("inMemoryRepository: ListContacts: Enter")
	r.lock.RLock()
	defer r.lock.RUnlock()

	if input == nil {
		return nil, ErrBadRequest
	}
	log.Debug().Msgf("inMemoryRepository: ListContacts: input: %v", input)
	log.Debug().Msgf("inMemoryRepository: ListContacts: total: %v", len(r.contacts))

	output = &clientapi.ListContactsResponse{
		ContactResponses: []*clientapi.ContactResponse{},
	}

	var minID int64 = math.MaxInt64
	for _, contact := range r.contacts {
		if minID > contact.ID {
			minID = contact.ID
		}
		cr := &clientapi.ContactResponse{
			ID:         contact.ID,
			TenantID:   contact.TenantID,
			FirstName:  contact.FirstName,
			LastName:   contact.LastName,
			Active:     contact.Active,
			Address:    contact.Address,
			SomeSecret: contact.SomeSecret,
			CreatedAt:  contact.CreatedAt,
			UpdatedAt:  contact.UpdatedAt,
		}
		output.ContactResponses = append(output.ContactResponses, cr)
	}
	output.NextPageToken = minID

	log.Debug().Msgf("inMemoryRepository: ListContacts: NextPageToken: %v", output.NextPageToken)
	log.Debug().Msgf("inMemoryRepository: ListContacts: ContactResponses: %v", output.ContactResponses)
	log.Debug().Msg("inMemoryRepository: ListContacts: Exit")
	return output, nil
}

func (r *inMemoryRepository) UpdateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msg("inMemoryRepository: UpdateContact: Enter")
	r.lock.Lock()
	defer r.lock.Unlock()

	if input == nil {
		return nil, ErrBadRequest
	}

	contact := r.contacts[input.ID]
	if contact == nil || input.ID == 0 {
		return nil, ErrNotFound
	}

	now := time.Now().UTC().Format(time.RFC3339)
	updated := &Contact{
		ID:         contact.ID,
		TenantID:   input.TenantID,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Active:     input.Active,
		Address:    input.Address,
		SomeSecret: input.SomeSecret,
		CreatedAt:  contact.CreatedAt,
		UpdatedAt:  now,
	}

	r.contacts[input.ID] = updated

	log.Debug().Msg("inMemoryRepository: UpdateContact: Exit")
	return &clientapi.ContactResponse{
		ID:         updated.ID,
		TenantID:   updated.TenantID,
		FirstName:  updated.FirstName,
		LastName:   updated.LastName,
		Active:     updated.Active,
		Address:    updated.Address,
		SomeSecret: updated.SomeSecret,
		CreatedAt:  updated.CreatedAt,
		UpdatedAt:  updated.UpdatedAt,
	}, nil
}

func (r *inMemoryRepository) DeleteContact(ctx context.Context, input *clientapi.DeleteContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msg("inMemoryRepository: DeleteContact: Enter")
	r.lock.Lock()
	defer r.lock.Unlock()

	if input == nil {
		return nil, ErrBadRequest
	}

	contact, exists := r.contacts[input.ID]
	if !exists || input.ID == 0 {
		return nil, ErrNotFound
	}

	delete(r.contacts, input.ID)
	log.Debug().Msgf("inMemoryRepository: DeleteContact: contact: %v", contact)

	log.Debug().Msg("inMemoryRepository: DeleteContact: Exit")
	return &clientapi.ContactResponse{
		ID:         contact.ID,
		TenantID:   contact.TenantID,
		FirstName:  contact.FirstName,
		LastName:   contact.LastName,
		Active:     contact.Active,
		Address:    contact.Address,
		SomeSecret: contact.SomeSecret,
		CreatedAt:  contact.CreatedAt,
		UpdatedAt:  contact.UpdatedAt,
	}, nil
}
