package addressbook_test

import (
	"context"
	"testing"

	"github.com/jgoerz/go-kit-crud/internal/addressbook"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	repo addressbook.Repository
}

func TestRunRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupSuite() {
	// Modify this for debugging
	zerolog.SetGlobalLevel(zerolog.Disabled)
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// func (s *RepositoryTestSuite) TearDownSuite() {}

func (s *RepositoryTestSuite) SetupTest() {
	assert := assert.New(s.T())

	s.repo = addressbook.NewInMemoryRepository()
	assert.NotNil(s.repo)
}

// func (s *RepositoryTestSuite) TearDownTest() {}

func (s *RepositoryTestSuite) TestCreateContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var testCases = []struct {
		name        string
		given       *addressbook.ContactRequest
		expected    *addressbook.ContactResponse
		expectedErr error
	}{
		{
			name: "create contact id: 1",
			given: &addressbook.ContactRequest{
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expectedErr: nil,
		},
		{
			name: "create contact id: 2",
			given: &addressbook.ContactRequest{
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
			},
			expectedErr: nil,
		},
		{
			name: "create contact, ignore ID",
			given: &addressbook.ContactRequest{
				ID:         9999,
				TenantID:   123,
				FirstName:  "first-name-123",
				LastName:   "last-name-123",
				Address:    "t123",
				SomeSecret: "secret-123",
			},
			expected: &addressbook.ContactResponse{
				ID:         3,
				TenantID:   123,
				FirstName:  "first-name-123",
				LastName:   "last-name-123",
				Address:    "t123",
				SomeSecret: "secret-123",
			},
			expectedErr: nil,
		},
		{
			name: "create contact, duplicate entry",
			given: &addressbook.ContactRequest{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.repo.CreateContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)

		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *RepositoryTestSuite) TestUpdateContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *addressbook.ContactRequest
		expected *addressbook.ContactResponse
	}{
		{
			given: &addressbook.ContactRequest{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &addressbook.ContactRequest{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.repo.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *addressbook.ContactRequest
		expected    *addressbook.ContactResponse
		expectedErr error
	}{
		{
			name: "update contact tenant 111",
			given: &addressbook.ContactRequest{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111-updated",
				LastName:   "last-name-111-updated",
				Address:    "address-111-updated",
				SomeSecret: "secret-111-updated",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111-updated",
				LastName:   "last-name-111-updated",
				Address:    "address-111-updated",
				SomeSecret: "secret-111-updated",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "update contact tenant 222",
			given: &addressbook.ContactRequest{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222-update",
				LastName:   "last-name-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222-update",
				LastName:   "last-name-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "update contact non-existent tenant",
			given: &addressbook.ContactRequest{
				TenantID:   9999,
				FirstName:  "first-name-222-update",
				LastName:   "last-name-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "update contact input is nil",
			given:       (*addressbook.ContactRequest)(nil),
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.repo.UpdateContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *RepositoryTestSuite) TestReadContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *addressbook.ContactRequest
		expected *addressbook.ContactResponse
	}{
		{
			given: &addressbook.ContactRequest{
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &addressbook.ContactRequest{
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.repo.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *addressbook.ReadContactRequest
		expected    *addressbook.ContactResponse
		expectedErr error
	}{
		{
			name: "read contact tenant 111",
			given: &addressbook.ReadContactRequest{
				ID: 1,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "read contact tenant 222",
			given: &addressbook.ReadContactRequest{
				ID: 2,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "read contact non-existent tenant",
			given: &addressbook.ReadContactRequest{
				ID: 9999,
			},
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "update contact input is nil",
			given:       (*addressbook.ReadContactRequest)(nil),
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.repo.ReadContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *RepositoryTestSuite) TestDeleteContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *addressbook.ContactRequest
		expected *addressbook.ContactResponse
	}{
		{
			given: &addressbook.ContactRequest{
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &addressbook.ContactRequest{
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.repo.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *addressbook.DeleteContactRequest
		expected    *addressbook.ContactResponse
		expectedErr error
	}{
		{
			name: "delete contact tenant 111",
			given: &addressbook.DeleteContactRequest{
				ID: 1,
			},
			expected: &addressbook.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "first-name-111",
				LastName:   "last-name-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "delete contact tenant 222",
			given: &addressbook.DeleteContactRequest{
				ID: 2,
			},
			expected: &addressbook.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "first-name-222",
				LastName:   "last-name-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "delete contact non-existent tenant",
			given: &addressbook.DeleteContactRequest{
				ID: 9999,
			},
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "delete contact input is nil",
			given:       (*addressbook.DeleteContactRequest)(nil),
			expected:    (*addressbook.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.repo.DeleteContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)

		if tc.given != nil {
			c, err := s.repo.ReadContact(ctx, &addressbook.ReadContactRequest{ID: tc.given.ID})
			assert.Nil(c)
			assert.ErrorIs(err, addressbook.ErrNotFound)
		}

	}
}
