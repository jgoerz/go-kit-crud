package grpc

import (
	"context"

	"github.com/jgoerz/go-kit-crud/internal/addressbook"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jgoerz/go-kit-crud/pb"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) addressbook.Service {
	var createContact = grpctransport.NewClient(
		conn, "pb.AddressBook", "CreateContact",
		EncodeGRPCContactRequest,
		DecodeGRPCContactResponse,
		pb.ContactResponse{},
	).Endpoint()

	var readContact = grpctransport.NewClient(
		conn, "pb.AddressBook", "ReadContact",
		EncodeGRPCReadContactRequest,
		DecodeGRPCContactResponse,
		pb.ContactResponse{},
	).Endpoint()

	var updateContact = grpctransport.NewClient(
		conn, "pb.AddressBook", "UpdateContact",
		EncodeGRPCContactRequest,
		DecodeGRPCContactResponse,
		pb.ContactResponse{},
	).Endpoint()

	var deleteContact = grpctransport.NewClient(
		conn, "pb.AddressBook", "DeleteContact",
		EncodeGRPCDeleteContactRequest,
		DecodeGRPCContactResponse,
		pb.ContactResponse{},
	).Endpoint()

	return addressbook.Endpoints{
		CreateContactEP: createContact,
		ReadContactEP:   readContact,
		UpdateContactEP: updateContact,
		DeleteContactEP: deleteContact,
	}
}

// gRPC client method
// Encode from domain object to protobuf
func EncodeGRPCContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*addressbook.ContactRequest)
	return &pb.ContactRequest{
		Id:         request.ID,
		TenantId:   request.TenantID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Active:     request.Active,
		Address:    request.Address,
		SomeSecret: request.SomeSecret,
	}, nil
}

// gRPC client method
// Decode from protobuf to domain object
func DecodeGRPCContactResponse(ctx context.Context, r any) (any, error) {
	response := r.(*pb.ContactResponse)
	return &addressbook.ContactResponse{
		ID:         response.Id,
		TenantID:   response.TenantId,
		FirstName:  response.FirstName,
		LastName:   response.LastName,
		Active:     response.Active,
		Address:    response.Address,
		SomeSecret: response.SomeSecret,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	}, nil
}

// gRPC client method
// Encode from domain object to protobuf
func EncodeGRPCReadContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*addressbook.ReadContactRequest)
	return &pb.ReadContactRequest{
		Id: request.ID,
	}, nil
}

// gRPC client method
// Encode from domain object to protobuf
func EncodeGRPCDeleteContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*addressbook.DeleteContactRequest)
	return &pb.DeleteContactRequest{
		Id: request.ID,
	}, nil
}
