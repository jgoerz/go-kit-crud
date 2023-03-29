package addressbook_test

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"

	"github.com/jgoerz/go-kit-crud/internal/addressbook"
	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var TestPKCS8PrivateKey string = `
-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCbwr4dnhiQg+g0
6Gzq/Dzdeufwai8V6ioqRJt4I2T+7FD16mBEQuFxrsgR8Tby+8k3vU6QFdTbw0M0
KcELWgOpLQf4zmHPFowGTnqq0WmtF62wueDmNCXiJQTWc8SArXHwdFugzcZU+v9A
q/IQnwqBUL0eLtsH3wsE/4Yp9RTmApaaB0gEQKwTDYfoXUbS8eOwsIitRIB7G9km
3YfmqsWgETcbEPkBZJYqEbh4EkhBrFK46Emmlj6rO/n/qRY/j3ulKIojosCIy+q+
4AQhSw9UVXjaQGT8/sW6rzdYmrAV7WQ7rLS7WTzZ4JpSMyCHe+oOvoqmbqvWBOb1
JrS1b9V3AgMBAAECggEATcKlQwgAX3Va4b7/UTjn8fJABKgeJaSntD5YF1wkOZgk
uwBtWubAwg5t13dC5X+J7wfVRt1/TM9op4wI0e/4T+cpSa9O6SHqeeOIHes6BK4D
imzhuEtkdkfg6GcXPN5aOZ79d4QDdb1w3Qp6aI3dor16DK17r6xMRgyDIEM4UbNF
tl6tRmbrEjr5pDTlsmOzmAeQ4BsYfgfCDPmCwnc2tqLPRsUFvqBOUoSJ3EID84XM
fnT+8NzR5ZsZDTiJxVdRpTEaruQABxNA43LPDy/6zQGAVlLsp3nD9gZLIHEpKufu
B/IzgHnBpyorW/sIRLeoUwdTxLTkbQZH5g1BxLlQoQKBgQDHJUZzjyl8Z0NTIl7X
gkjONOcuwhvbZapRXnEhywCO75ObrmKeR0Gk+GXk04aI7+5KJ52Guvb7nU09G/Xw
uGtC+F0XKfc2ptChfL6oD9VPMjp3bMuB65egKpzdV+K/ZEct8MI1EsB9Tc5yGWWv
pqwrMIxKj8Ci2k3ohd/GkF8F6QKBgQDIOqVab0SGb2KEhsiqgTVARQmlOTszIq4z
UfcAyCGi+xezywqdvQP3Ml5gogcDDR+6slyzIbzhGtHIQ4FkG9QfaefSTut5BT5/
lfK01CO4HQBoWgLEiKzl76MsRWIv+jxfExg6oEb7WZbK/BTwTxb4m2DYsjl5L7Xk
KFjph00EXwKBgQC/0cG4kY8uSvDoZNTh1JZ4OTDtMv9OJvEVC1kBad4Rz+ZoMGLB
fnVWiATtkmmmASWPu/TZz8ESv4OkdwhAZAK9MSnJpByBQdD3m4axrv6SGBmE6wBj
FiCooCMUeRDptZdyQtNt96/9gjJ2aMwvkuWHfG3FbA3rT0d3z2uqgWll8QKBgEre
Kt/ixPOjiGnXYAbpIzkx10ZxXOJk8E/+MOaY7oLbcmRm4kRS3b27lrB5RTft21Ra
xvCwB8j/1zsTirkc8rcASY9ItSFeRZ09OzBENkrshS9/oJNOK6Aad5/hHbKk1ZgT
MrcRIRlwyUKC+W1VlVhF+PNtyLG4lkGGmKBRWAnvAoGAMyTmCt+51qDF3P8ERX+U
C3FLysZqnEtaGB2rzOxqPI3tGT9R/PI8/zvJBjaenfF5Xs3Swy7t6tjYgsLUFHfB
H6wxyu8SdoVa0YRXBJzrbzqAuxKe35Jee9LjJZzxu2rgPLmY7WU9+KtaC0NNqdhj
+Sy/oU/bw4W4BVvBXr+PF1o=
-----END RSA PRIVATE KEY-----`

type ServiceTestSuite struct {
	suite.Suite
	srv        addressbook.Service
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func TestRunServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupSuite() {
	// Modify this for debugging
	zerolog.SetGlobalLevel(zerolog.Disabled)
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// func (s *ServiceTestSuite) TearDownSuite() {}

func (s *ServiceTestSuite) SetupTest() {
	assert := assert.New(s.T())

	block, _ := pem.Decode([]byte(TestPKCS8PrivateKey))
	if block == nil {
		panic("Couldn't decode private key")
	}

	pKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	privateKey, ok := pKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("could not type assert private key to *rsa.PrivateKey, got: %T\n", pKey)
		panic("could not type assert private key to *rsa.PrivateKey")
	}

	// This is for the Encryption/Decryption tests
	s.PublicKey = &privateKey.PublicKey
	s.PrivateKey = privateKey

	repo := addressbook.NewInMemoryRepository()
	assert.NotNil(repo)

	s.srv = addressbook.NewService(repo, privateKey)
}

// func (s *ServiceTestSuite) TearDownTest() {}

func (s *ServiceTestSuite) TestStringEncryptionDecryption() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var data = []struct {
		given    string
		expected string
	}{
		{
			given:    "secret123",
			expected: "secret123",
		},
		{
			given:    "helloThereWorldOK",
			expected: "helloThereWorldOK",
		},
	}

	for _, item := range data {
		encrypted, err := addressbook.EncryptString(item.given, s.PublicKey)
		assert.NoError(err)
		assert.NotNil(encrypted)

		decrypted, err := addressbook.DecryptString(encrypted, s.PrivateKey)
		assert.NoError(err)
		assert.NotEqual(encrypted, decrypted)
		assert.Equal(item.expected, decrypted)
	}
}

func (s *ServiceTestSuite) TestStructDecryption() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	type TestStruct struct {
		PlainItem     string
		EncryptedItem string `encryption:"true"`
	}

	var data = []struct {
		given    *TestStruct
		expected *TestStruct
	}{
		{
			given: &TestStruct{
				PlainItem:     "secret123",
				EncryptedItem: "bN31V2fTEaSiHaYBvSJXYUsq7GCYklmO580OLVzpkYfxlWQaIkMNa1lwpN1083QNTXiRYQ2Sl/jGthDDjvXutb0mXy99AX5h/xBDqnV1uD/qSAOvN0zIYp07E4VytkDpyUYREAWbjtUBUSW+4iMHlk+o/J9WyN4GP4HKDMLz88Sw53+yGQJ3zzBZ+oWJrEwSBkbxYoHxRUZu/UJQfCp0hkeMvT4lwGY7ZaeKorO5viFL3MPE/KwCmx8Xo2Nz1OXcqa5G+nPlipA3E5NynWt5+kJvMHH5C3+mj51ITy8zYniVR/yRgJxdogoBWj80/ec+/v5eg11OWyFY0dCIkwJwvg==",
			},
			expected: &TestStruct{
				PlainItem:     "secret123",
				EncryptedItem: "secret123"}, // decrypted expectation
		},
		{
			given: &TestStruct{
				PlainItem:     "helloThereWorldOK",
				EncryptedItem: "MQPtt3X5zbW/rlCuCH3qmR9k4orYovlI6jYxVa/Q4p9+sDYvP142vNmmbDR1rJslcR6gorrMvhve9ttIsCkJJLqOd7anTn9yZj86jMRhEgoI2Yr4mxjWDCbtAm9J6BwFrT8HQACHOdtjzPZxPiWNSQsbyCAegzTe93eCX5sehXTkxd3b6eewl0ko+XVTMwS/I0JuqbDFv+56nEX4cLPns8cphGmH2ta6sVOYLPRfPokiUeQz2T0/9AFviH2nr4ASlzii7cGKimc4M2XLDR0qCYwiM3oTrK2z0xAgvHyX9Qm5oh/rxUDadVbBUhbroFJiVrTwKMmZKEFEQbE08yap+w=="},
			expected: &TestStruct{
				PlainItem:     "helloThereWorldOK",
				EncryptedItem: "helloThereWorldOK", // decrypted expectation
			},
		},
	}

	for _, item := range data {
		decrypted, err := addressbook.Decrypt(item.given, s.PrivateKey)
		assert.NoError(err)
		assert.Equal(item.expected.PlainItem, decrypted.PlainItem)
		assert.Equal(item.expected.EncryptedItem, decrypted.EncryptedItem)
	}
}

func (s *ServiceTestSuite) TestStructDecryptionErrors() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	// Non struct argument
	failString := "fail"
	expectedErr := errors.New("argument must be a struct")
	_, err := addressbook.Decrypt(&failString, s.PrivateKey)
	assert.Equal(expectedErr, err)

	// struct argument with badly formatted cipher text
	type TestStruct struct {
		Text string `encryption:"true"`
	}
	input := &TestStruct{"not-base64-encoded"}

	expectedErr = errors.New("Service.DecryptString.DecodeString: illegal base64 data at input byte 3")
	_, err = addressbook.Decrypt(input, s.PrivateKey)
	assert.Equal(expectedErr, err)

	expectedErr = errors.New("Service.DecryptString.DecryptOAEP: crypto/rsa: decryption error")
	_, err = addressbook.DecryptString(failString, s.PrivateKey)
	assert.Equal(expectedErr, err)
}

func (s *ServiceTestSuite) TestStructEncryption() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	type TestStruct struct {
		PlainItem     string
		EncryptedItem string `encryption:"true"`
	}

	var data = []struct {
		given    *TestStruct
		expected *TestStruct
	}{
		{
			given: &TestStruct{
				PlainItem:     "secret123",
				EncryptedItem: "secret123",
			},
			expected: &TestStruct{
				PlainItem:     "secret123",
				EncryptedItem: "secret123",
			},
		},
		{
			given: &TestStruct{
				PlainItem:     "helloThereWorldOK",
				EncryptedItem: "helloThereWorldOK",
			},
			expected: &TestStruct{
				PlainItem:     "helloThereWorldOK",
				EncryptedItem: "helloThereWorldOK",
			},
		},
	}

	for _, item := range data {
		encrypted, err := addressbook.Encrypt(item.given, s.PublicKey)
		assert.NoError(err)
		assert.Equal(item.expected.PlainItem, encrypted.PlainItem)
		assert.NotEqual(item.expected.EncryptedItem, encrypted.EncryptedItem)

		// encrypted should be a base64 encoded string
		rawMsg, err := base64.StdEncoding.DecodeString(encrypted.EncryptedItem)
		assert.NoError(err)
		assert.NotNil(rawMsg)

		decrypted, err := addressbook.Decrypt(encrypted, s.PrivateKey)
		assert.NoError(err)
		assert.Equal(item.expected.PlainItem, decrypted.PlainItem)
		assert.Equal(item.expected.EncryptedItem, decrypted.EncryptedItem)
	}
}

func (s *ServiceTestSuite) TestStructEncryptionErrors() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	// Non struct argument
	failString := "fail"
	expectedErr := errors.New("argument must be a struct")
	_, err := addressbook.Encrypt(&failString, s.PublicKey)
	assert.Equal(expectedErr, err)

	// string exceeds maximum size (190 bytes) for 2048 bit key with SHA256
	type TestStruct struct {
		Text string `encryption:"true"`
	}
	input := &TestStruct{
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901",
	}

	expectedErr = errors.New("Service.EncryptString.EncryptOAEP: crypto/rsa: message too long for RSA key size")
	_, err = addressbook.Encrypt(input, s.PublicKey)
	assert.Equal(expectedErr, err)
}

func (s *ServiceTestSuite) TestServiceCreateContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())
	var testCases = []struct {
		name        string
		given       *clientapi.ContactRequest
		expected    *clientapi.ContactResponse
		expectedErr error
	}{
		{
			name: "create contact id: 1",
			given: &clientapi.ContactRequest{
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expectedErr: nil,
		},
		{
			name: "create contact id: 2",
			given: &clientapi.ContactRequest{
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
			},
			expectedErr: nil,
		},
		{
			name: "create contact, ignore ID ",
			given: &clientapi.ContactRequest{
				ID:         9999,
				TenantID:   123,
				FirstName:  "firstName-123",
				LastName:   "lastName-123",
				Address:    "t123",
				SomeSecret: "secret-123",
			},
			expected: &clientapi.ContactResponse{
				ID:         3,
				TenantID:   123,
				FirstName:  "firstName-123",
				LastName:   "lastName-123",
				Address:    "t123",
				SomeSecret: "secret-123",
			},
			expectedErr: nil,
		},
		// {
		// 	name: "create contact, duplicate entry",
		// 	given: &clientapi.ContactRequest{
		// 		TenantID:   111,
		// 		FirstName:  "firstName-111",
		// 		LastName:   "lastName-111",
		// 		Address:    "address-111",
		// 		SomeSecret: "secret-111",
		// 	},
		// 	expected:    (*clientapi.ContactResponse)(nil),
		// 	expectedErr: addressbook.ErrBadRequest,
		// },
		{
			name: "create contact, negative tenantID",
			given: &clientapi.ContactRequest{
				TenantID:   -1,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
			},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
		{
			name:        "create contact, input nil",
			given:       (*clientapi.ContactRequest)(nil),
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.srv.CreateContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)

		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *ServiceTestSuite) TestServiceReadContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *clientapi.ContactRequest
		expected *clientapi.ContactResponse
	}{
		{
			given: &clientapi.ContactRequest{
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &clientapi.ContactRequest{
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.srv.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *clientapi.ReadContactRequest
		expected    *clientapi.ContactResponse
		expectedErr error
	}{
		{
			name: "read contact tenant 111",
			given: &clientapi.ReadContactRequest{
				ID: 1,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "read contact tenant 222",
			given: &clientapi.ReadContactRequest{
				ID: 2,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "read contact non-existent",
			given: &clientapi.ReadContactRequest{
				ID: 9999,
			},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "read contact input is nil",
			given:       (*clientapi.ReadContactRequest)(nil),
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
		{
			name:        "read contact tenantID is invalid",
			given:       &clientapi.ReadContactRequest{ID: 0},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.srv.ReadContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *ServiceTestSuite) TestServiceUpdateContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *clientapi.ContactRequest
		expected *clientapi.ContactResponse
	}{
		{
			given: &clientapi.ContactRequest{
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &clientapi.ContactRequest{
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.srv.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *clientapi.ContactRequest
		expected    *clientapi.ContactResponse
		expectedErr error
	}{
		{
			name: "update contact tenant 111",
			given: &clientapi.ContactRequest{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111-updated",
				LastName:   "lastName-111-updated",
				Address:    "address-111-updated",
				SomeSecret: "secret-111-updated",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111-updated",
				LastName:   "lastName-111-updated",
				Address:    "address-111-updated",
				SomeSecret: "secret-111-updated",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "update contact tenant 222",
			given: &clientapi.ContactRequest{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222-update",
				LastName:   "lastName-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222-update",
				LastName:   "lastName-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "update contact non-existent",
			given: &clientapi.ContactRequest{
				ID:         9999,
				TenantID:   9999,
				FirstName:  "firstName-222-update",
				LastName:   "lastName-222-update",
				Address:    "address-222-update",
				SomeSecret: "secret-222-update",
				Active:     true,
			},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "update contact input is nil",
			given:       (*clientapi.ContactRequest)(nil),
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
		{
			name:        "update contact ID is invalid",
			given:       &clientapi.ContactRequest{ID: 0},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.srv.UpdateContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)
	}
}

func (s *ServiceTestSuite) TestServiceDeleteContact() {
	// s.T().Skip("Pending")
	assert := assert.New(s.T())

	var seeds = []struct {
		given    *clientapi.ContactRequest
		expected *clientapi.ContactResponse
	}{
		{
			given: &clientapi.ContactRequest{
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
		},
		{
			given: &clientapi.ContactRequest{
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
		},
	}
	for _, seed := range seeds {
		ctx := context.Background()
		contact, err := s.srv.CreateContact(ctx, seed.given)
		assert.NoError(err)
		assert.NotNil(contact)
	}

	var testCases = []struct {
		name        string
		given       *clientapi.DeleteContactRequest
		expected    *clientapi.ContactResponse
		expectedErr error
	}{
		{
			name: "delete contact 1",
			given: &clientapi.DeleteContactRequest{
				ID: 1,
			},
			expected: &clientapi.ContactResponse{
				ID:         1,
				TenantID:   111,
				FirstName:  "firstName-111",
				LastName:   "lastName-111",
				Address:    "address-111",
				SomeSecret: "secret-111",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "delete contact 2",
			given: &clientapi.DeleteContactRequest{
				ID: 2,
			},
			expected: &clientapi.ContactResponse{
				ID:         2,
				TenantID:   222,
				FirstName:  "firstName-222",
				LastName:   "lastName-222",
				Address:    "address-222",
				SomeSecret: "secret-222",
				Active:     true,
			},
			expectedErr: nil,
		},
		{
			name: "delete contact non-existent",
			given: &clientapi.DeleteContactRequest{
				ID: 9999,
			},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrNotFound,
		},
		{
			name:        "delete contact input is nil",
			given:       (*clientapi.DeleteContactRequest)(nil),
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
		{
			name:        "delete contact invalid ID",
			given:       &clientapi.DeleteContactRequest{ID: 0},
			expected:    (*clientapi.ContactResponse)(nil),
			expectedErr: addressbook.ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		contact, err := s.srv.DeleteContact(ctx, tc.given)
		assert.ErrorIs(err, tc.expectedErr, tc.name)
		if contact != nil {
			contact.CreatedAt = ""
			contact.UpdatedAt = ""
		}
		assert.Equal(tc.expected, contact, tc.name)

		if tc.expected != nil {
			c, err := s.srv.ReadContact(ctx, &clientapi.ReadContactRequest{ID: tc.given.ID})
			assert.Nil(c, tc.name)
			assert.ErrorIs(err, addressbook.ErrNotFound, tc.name)
		}

	}
}
