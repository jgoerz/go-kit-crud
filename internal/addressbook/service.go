package addressbook

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"

	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"
	"github.com/rs/zerolog/log"
)

const ContactMaxPageSize = 10

type Service interface {
	CreateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error)
	ReadContact(ctx context.Context, input *clientapi.ReadContactRequest) (output *clientapi.ContactResponse, err error)
	ListContacts(ctx context.Context, input *clientapi.ListContactsRequest) (output *clientapi.ListContactsResponse, err error)
	UpdateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error)
	DeleteContact(ctx context.Context, input *clientapi.DeleteContactRequest) (output *clientapi.ContactResponse, err error)
}

// Look at https://betterstack.com/community/guides/logging/zerolog/#using-zerolog-in-a-web-application
// for setting up a logger you can "Get" that is preconfigured.
// https://betterstack.com/community/guides/logging/zerolog/#creating-a-logging-middleware
func NewService(repo Repository, privateKey *rsa.PrivateKey) Service {
	return &addressBook{
		publicKey:  &privateKey.PublicKey,
		privateKey: privateKey,
		repo:       repo,
	}
}

type addressBook struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	repo       Repository
}

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal error")
)

func (service *addressBook) CreateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msgf("Service: CreateContact: obj: %T, %v", input, input)

	// Input validation
	if input == nil {
		return nil, fmt.Errorf("service.CreateContact: no input: %w", ErrBadRequest)
	}

	if input.ID != 0 {
		input.ID = 0
	}

	if input.TenantID < 0 {
		return nil, fmt.Errorf("service.CreateContact: invalid TenantID: %v %w", input.TenantID, ErrBadRequest)
	}

	// if ctxTenantID != input.TenantID {
	// 	msg := fmt.Sprintf("Unauthorized.  Authorized tenant: %v, accessing for tenant: %v", ctxTenantID, input.TenantID)
	// 	log.Error().Msg(msg)
	// 	return nil, fmt.Errorf("Service: CreateContact: %s: %w", msg, ErrBadRequest)
	// }

	// Do the work
	in, err := Encrypt(input, service.publicKey)
	if err != nil {
		return nil, fmt.Errorf("service.CreateContact.Encrypt: %v: %w", err, ErrInternal)
	}

	contact, err := service.repo.CreateContact(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("service.CreateContact: %w", err)
	}

	decrypted, err := Decrypt(contact, service.privateKey)
	if err != nil {
		return nil, fmt.Errorf("service.CreateContact.decrypt: %v: %w", err, ErrInternal)
	}
	output = decrypted

	return output, nil
}

func (service *addressBook) ReadContact(ctx context.Context, input *clientapi.ReadContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msgf("Service: ReadContact: input: %v", input)

	// Input validation
	if input == nil {
		return nil, fmt.Errorf("service.ReadContact: no input: %w", ErrBadRequest)
	}

	if input.ID <= 0 {
		msg := fmt.Sprintf("service.ReadContact: invalid ID: %v", input.ID)
		log.Error().Msg(msg)
		return nil, fmt.Errorf("%s %w", msg, ErrBadRequest)
	}

	contact, err := service.repo.ReadContact(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("service.ReadContact: %w", err)
	}

	decrypted, err := Decrypt(contact, service.privateKey)
	if err != nil {
		return nil, fmt.Errorf("service.ReadContact.decrypt: %w", err)
	}
	output = decrypted

	return output, nil
}

func (service *addressBook) ListContacts(ctx context.Context, input *clientapi.ListContactsRequest) (output *clientapi.ListContactsResponse, err error) {
	log.Debug().Msgf("Service: ListContacts: input: %v", input)

	output = &clientapi.ListContactsResponse{
		ContactResponses: []*clientapi.ContactResponse{},
		NextPageToken:    0,
	}

	// Input validation
	if input == nil {
		return nil, fmt.Errorf("service.ListContacts: no input: %w", ErrBadRequest)
	}

	if input.PageToken < 0 {
		input.PageToken = 0
	}
	if input.PageSize <= 0 {
		input.PageSize = ContactMaxPageSize
	}

	response, err := service.repo.ListContacts(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("service.ListContacts: %w", err)
	}
	output.NextPageToken = response.NextPageToken

	for _, contact := range response.ContactResponses {
		decrypted, err := Decrypt(contact, service.privateKey)
		if err != nil {
			return nil, fmt.Errorf("service.ListContacts.decrypt: %w", err)
		}
		output.ContactResponses = append(output.ContactResponses, decrypted)
	}

	log.Debug().Msgf("Service: ListContacts: output: %v", output)
	return output, nil
}

func (service *addressBook) UpdateContact(ctx context.Context, input *clientapi.ContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msgf("Service: UpdateContact: input: %T %v", input, input)

	// Input validation
	if input == nil {
		return nil, fmt.Errorf("service.UpdateContact: no input: %w", ErrBadRequest)
	}

	// if ctxTenantID == 0 {
	// 	msg := fmt.Sprintf("Unauthorized.  Authorized tenant: %v, accessing for tenant: %v", ctxTenantID, ctxTenantID)
	// 	log.Error().Msg(msg)
	// 	return nil, fmt.Errorf("Service: UpdateContact: %s: %w", msg, ErrBadRequest)
	// }

	if input.TenantID <= 0 {
		msg := fmt.Sprintf("service.UpdateContact: invalid TenantID: '%v'", input.TenantID)
		log.Error().Msg(msg)
		return nil, fmt.Errorf(msg+" %w", ErrBadRequest)
	}

	input, err = Encrypt(input, service.publicKey)
	if err != nil {
		return nil, fmt.Errorf("service.UpdateContact.encrypt: %w", err)
	}

	contact, err := service.repo.UpdateContact(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("service.UpdateContact: %w", err)
	}

	decrypted, err := Decrypt(contact, service.privateKey)
	if err != nil {
		return nil, fmt.Errorf("service.UpdatedContact.decrypt: %w", err)
	}
	output = decrypted

	return output, nil
}

func (service *addressBook) DeleteContact(ctx context.Context, input *clientapi.DeleteContactRequest) (output *clientapi.ContactResponse, err error) {
	log.Debug().Msgf("Service: DeleteContact: %v", input)

	// Input validation
	if input == nil {
		return nil, fmt.Errorf("service.DeleteContact: no input: %w", ErrBadRequest)
	}

	if input.ID <= 0 {
		msg := fmt.Sprintf("service.DeleteContact: invalid ID: '%v'", input.ID)
		log.Error().Msg(msg)
		return nil, fmt.Errorf(msg+" %w", ErrBadRequest)
	}

	// if ctxTenantID == 0 {
	// 	msg := fmt.Sprintf("Unauthorized.  Authorized tenant: %v, accessing for tenant: %v", ctxTenantID, tenantID)
	// 	log.Error().Msg(msg)
	// 	return fmt.Errorf("Service: DeleteContact: %s: %w", msg, ErrBadRequest)
	// }

	contact, err := service.repo.DeleteContact(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("service.DeleteContact: %w", err)
	}

	decrypted, err := Decrypt(contact, service.privateKey)
	if err != nil {
		return nil, fmt.Errorf("service.ReadContact.decrypt: %w", err)
	}
	output = decrypted

	return output, err
}

// References:
// https://engineering.kablamo.com.au/posts/2022/field-level-data-encryption-in-go
// https://www.sohamkamani.com/golang/rsa-encryption/#encryption
// Limits on size of message you can encrypt:
// https://security.stackexchange.com/questions/112029/should-sha-1-be-used-with-rsa-oaep#answer-112032

// Possible solution for larger message sizes (limit is 190 bytes with 2048 key and SHA256):
// https://www.nsoftware.com/kb/xml/09051401.rst
func Encrypt[T any](obj T, publicKey *rsa.PublicKey) (T, error) {
	log.Debug().Msgf("Service.Encrypt: %T, nil?: %v", obj, reflect.ValueOf(obj).IsZero())
	// v := reflect.ValueOf(&obj).Elem() // using a reflected pointer value allows us to manipulate the underlying field values
	v := reflect.ValueOf(obj).Elem() // using a reflected pointer value allows us to manipulate the underlying field values
	if v.Kind() != reflect.Struct {  // We only want to support struct types
		return obj, fmt.Errorf("argument must be a struct")
	}

	for i := 0; i < v.NumField(); i++ { // iterate through each field of the struct
		f := v.Field(i)
		if f.Kind() == reflect.String { // We only want to support encryption for string types
			if v.Type().Field(i).Tag.Get("encryption") == "true" { // check if the field has encryption:"true"
				plainText := f.String()                                // fetch the plaintext value from the reflected struct field
				cipherText, err := EncryptString(plainText, publicKey) // encrypt the value
				if err != nil {
					return obj, err
				}
				f.SetString(cipherText) // set the new cipherText value back into the reflected struct field
			}
		}
	}

	log.Debug().Msgf("Service.Encrypt: complete")
	return obj, nil
}

// References:
// https://www.sohamkamani.com/golang/rsa-encryption/#decryption
func Decrypt[T any](obj T, privateKey *rsa.PrivateKey) (T, error) {
	log.Debug().Msgf("Service.Decrypt: %T, nil?: %v", obj, reflect.ValueOf(obj).IsZero())

	// v := reflect.ValueOf(&obj).Elem() // using a reflected pointer value allows us to manipulate the underlying field values
	v := reflect.ValueOf(obj).Elem() // using a reflected pointer value allows us to manipulate the underlying field values
	if v.Kind() != reflect.Struct {  // We only want to support struct types
		return obj, fmt.Errorf("argument must be a struct")
	}

	for i := 0; i < v.NumField(); i++ { // iterate through each field of the struct
		f := v.Field(i)
		if f.Kind() == reflect.String { // We only want to support decryption for string types
			if v.Type().Field(i).Tag.Get("encryption") == "true" { // check if the field has encryption:"true"
				cipherText := f.String() // fetch the cipherText value from the reflected struct field
				log.Debug().Msgf("Service.Decrypt: cipherText: %T, empty?: %v", cipherText, cipherText == "")
				plainText, err := DecryptString(cipherText, privateKey) // decrypt the value
				if err != nil {
					return obj, err
				}
				f.SetString(plainText) // set the new plainText value back into the reflected struct field
			}
		}
	}

	log.Debug().Msgf("Service.Decrypt: complete")
	return obj, nil
}

func EncryptString(msg string, publicKey *rsa.PublicKey) (string, error) {
	log.Debug().Msgf("Service.EncryptString: %T, empty?: %v", msg, msg == "")
	log.Debug().Msgf("Service.EncryptString.PublicKey: %T, nil?: %v", publicKey, publicKey == nil)

	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(msg), nil)
	if err != nil {
		return "", fmt.Errorf("Service.EncryptString.EncryptOAEP: %v", err)
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func DecryptString(msg string, privateKey *rsa.PrivateKey) (string, error) {
	log.Debug().Msgf("Service.DecryptString: %T, empty?: %v", msg, msg == "")
	log.Debug().Msgf("Service.DecryptString.PrivateKey: %T, nil?: %v", privateKey, privateKey == nil)

	rawMsg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", fmt.Errorf("Service.DecryptString.DecodeString: %v", err)
	}

	// decryptedBytes, err := privateKey.Decrypt(rand.Reader, []byte(rawMsg), &rsa.OAEPOptions{Hash: crypto.SHA256})
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, rawMsg, nil)
	if err != nil {
		return "", fmt.Errorf("Service.DecryptString.DecryptOAEP: %v", err)
	}

	return string(decrypted), nil
}

// Gets called if there is a decode request error
// Gets called if there is a endpoint error
// Gets called if there is a encode response error
// Used for both HTTP and gRPC, returns in gRPC, calls errorEncoder in HTTP
// func CreateContactErrorHandler(ctx context.Context, err error) {
// }
