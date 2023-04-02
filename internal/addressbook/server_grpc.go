package addressbook

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	clientapi "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"
	"github.com/jgoerz/go-kit-crud/pkg/client/pb"
)

func NewGRPCServer(ctx context.Context, endpoints clientapi.Endpoints) pb.AddressBookServer {
	return &grpcServer{
		createContact: grpctransport.NewServer(
			endpoints.CreateContactEP,
			DecodeGRPCContactRequest,
			EncodeGRPCContactResponse,
		),
		readContact: grpctransport.NewServer(
			endpoints.ReadContactEP,
			DecodeGRPCReadContactRequest,
			EncodeGRPCContactResponse,
		),
		listContacts: grpctransport.NewServer(
			endpoints.ListContactsEP,
			DecodeGRPCListContactsRequest,
			EncodeGRPCListContactsResponse,
		),
		updateContact: grpctransport.NewServer(
			endpoints.UpdateContactEP,
			DecodeGRPCContactRequest,
			EncodeGRPCContactResponse,
		),
		deleteContact: grpctransport.NewServer(
			endpoints.DeleteContactEP,
			DecodeGRPCDeleteContactRequest,
			EncodeGRPCContactResponse,
		),
	}
}

type grpcServer struct {
	pb.UnimplementedAddressBookServer
	createContact grpctransport.Handler
	readContact   grpctransport.Handler
	listContacts  grpctransport.Handler
	updateContact grpctransport.Handler
	deleteContact grpctransport.Handler
}

func (s *grpcServer) CreateContact(ctx context.Context, r *pb.ContactRequest) (*pb.ContactResponse, error) {
	_, resp, err := s.createContact.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ContactResponse), nil
}

func (s *grpcServer) ReadContact(ctx context.Context, r *pb.ReadContactRequest) (*pb.ContactResponse, error) {
	_, resp, err := s.readContact.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ContactResponse), nil
}

func (s *grpcServer) ListContacts(ctx context.Context, r *pb.ListContactsRequest) (*pb.ListContactsResponse, error) {
	_, resp, err := s.listContacts.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListContactsResponse), nil
}

func (s *grpcServer) UpdateContact(ctx context.Context, r *pb.ContactRequest) (*pb.ContactResponse, error) {
	_, resp, err := s.updateContact.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ContactResponse), nil
}

func (s *grpcServer) DeleteContact(ctx context.Context, r *pb.DeleteContactRequest) (*pb.ContactResponse, error) {
	_, resp, err := s.deleteContact.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ContactResponse), nil
}

// gRPC server method
// Decode from protobuf to domain object
func DecodeGRPCContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*pb.ContactRequest)
	return &clientapi.ContactRequest{
		ID:         request.Id,
		TenantID:   request.TenantId,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Active:     request.Active,
		Address:    request.Address,
		SomeSecret: request.SomeSecret,
	}, nil
}

// gRPC server method
// Encode from domain object to protobuf
func EncodeGRPCContactResponse(ctx context.Context, r any) (any, error) {
	resp := r.(*clientapi.ContactResponse)
	return &pb.ContactResponse{
		Id:         resp.ID,
		TenantId:   resp.TenantID,
		FirstName:  resp.FirstName,
		LastName:   resp.LastName,
		Active:     resp.Active,
		Address:    resp.Address,
		SomeSecret: resp.SomeSecret,
		CreatedAt:  resp.CreatedAt,
		UpdatedAt:  resp.UpdatedAt,
	}, nil
}

// gRPC server method
// Decode from protobuf to domain object
func DecodeGRPCReadContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*pb.ReadContactRequest)
	return &clientapi.ReadContactRequest{
		ID: request.Id,
	}, nil
}

// gRPC server method
// Decode from protobuf to domain object
func DecodeGRPCDeleteContactRequest(ctx context.Context, r any) (any, error) {
	request := r.(*pb.DeleteContactRequest)
	return &clientapi.DeleteContactRequest{
		ID: request.Id,
	}, nil
}

// gRPC server method
// Decode from protobuf to domain object
func DecodeGRPCListContactsRequest(ctx context.Context, r any) (any, error) {
	request := r.(*pb.ListContactsRequest)

	return &clientapi.ListContactsRequest{
		PageToken: request.PageToken,
		PageSize:  request.PageSize,
	}, nil
}

// gRPC server method
// Encode from domain object to protobuf
func EncodeGRPCListContactsResponse(ctx context.Context, r any) (any, error) {
	resp := r.(*clientapi.ListContactsResponse)

	contactResponses := []*pb.ContactResponse{}
	for _, contact := range resp.ContactResponses {
		cr := &pb.ContactResponse{
			Id:         contact.ID,
			TenantId:   contact.TenantID,
			FirstName:  contact.FirstName,
			LastName:   contact.LastName,
			Active:     contact.Active,
			Address:    contact.Address,
			SomeSecret: contact.SomeSecret,
			CreatedAt:  contact.CreatedAt,
			UpdatedAt:  contact.UpdatedAt,
		}
		contactResponses = append(contactResponses, cr)
	}

	return &pb.ListContactsResponse{
		ContactResponses: contactResponses,
		NextPageToken:    resp.NextPageToken,
	}, nil
}
