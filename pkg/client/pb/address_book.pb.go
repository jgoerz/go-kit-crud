// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: pb/address_book.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId   int64  `protobuf:"varint,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Active     bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	Address    string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	SomeSecret string `protobuf:"bytes,7,opt,name=some_secret,json=someSecret,proto3" json:"some_secret,omitempty"`
}

func (x *ContactRequest) Reset() {
	*x = ContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_address_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRequest) ProtoMessage() {}

func (x *ContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRequest.ProtoReflect.Descriptor instead.
func (*ContactRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_book_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *ContactRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ContactRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ContactRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *ContactRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactRequest) GetSomeSecret() string {
	if x != nil {
		return x.SomeSecret
	}
	return ""
}

type ReadContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadContactRequest) Reset() {
	*x = ReadContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_address_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadContactRequest) ProtoMessage() {}

func (x *ReadContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadContactRequest.ProtoReflect.Descriptor instead.
func (*ReadContactRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_book_proto_rawDescGZIP(), []int{1}
}

func (x *ReadContactRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteContactRequest) Reset() {
	*x = DeleteContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_address_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContactRequest) ProtoMessage() {}

func (x *DeleteContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContactRequest.ProtoReflect.Descriptor instead.
func (*DeleteContactRequest) Descriptor() ([]byte, []int) {
	return file_pb_address_book_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteContactRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId   int64  `protobuf:"varint,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Active     bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	Address    string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	SomeSecret string `protobuf:"bytes,7,opt,name=some_secret,json=someSecret,proto3" json:"some_secret,omitempty"`
	CreatedAt  string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Error      string `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ContactResponse) Reset() {
	*x = ContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_address_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactResponse) ProtoMessage() {}

func (x *ContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_address_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactResponse.ProtoReflect.Descriptor instead.
func (*ContactResponse) Descriptor() ([]byte, []int) {
	return file_pb_address_book_proto_rawDescGZIP(), []int{3}
}

func (x *ContactResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactResponse) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *ContactResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ContactResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ContactResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *ContactResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactResponse) GetSomeSecret() string {
	if x != nil {
		return x.SomeSecret
	}
	return ""
}

func (x *ContactResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ContactResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ContactResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pb_address_book_proto protoreflect.FileDescriptor

var file_pb_address_book_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xcc, 0x01, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa1, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x6d, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x85, 0x02, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3a, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x67, 0x6f, 0x65, 0x72, 0x7a, 0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x69, 0x74,
	0x2d, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_address_book_proto_rawDescOnce sync.Once
	file_pb_address_book_proto_rawDescData = file_pb_address_book_proto_rawDesc
)

func file_pb_address_book_proto_rawDescGZIP() []byte {
	file_pb_address_book_proto_rawDescOnce.Do(func() {
		file_pb_address_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_address_book_proto_rawDescData)
	})
	return file_pb_address_book_proto_rawDescData
}

var file_pb_address_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_address_book_proto_goTypes = []interface{}{
	(*ContactRequest)(nil),       // 0: pb.ContactRequest
	(*ReadContactRequest)(nil),   // 1: pb.ReadContactRequest
	(*DeleteContactRequest)(nil), // 2: pb.DeleteContactRequest
	(*ContactResponse)(nil),      // 3: pb.ContactResponse
}
var file_pb_address_book_proto_depIdxs = []int32{
	0, // 0: pb.AddressBook.CreateContact:input_type -> pb.ContactRequest
	1, // 1: pb.AddressBook.ReadContact:input_type -> pb.ReadContactRequest
	0, // 2: pb.AddressBook.UpdateContact:input_type -> pb.ContactRequest
	2, // 3: pb.AddressBook.DeleteContact:input_type -> pb.DeleteContactRequest
	3, // 4: pb.AddressBook.CreateContact:output_type -> pb.ContactResponse
	3, // 5: pb.AddressBook.ReadContact:output_type -> pb.ContactResponse
	3, // 6: pb.AddressBook.UpdateContact:output_type -> pb.ContactResponse
	3, // 7: pb.AddressBook.DeleteContact:output_type -> pb.ContactResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_address_book_proto_init() }
func file_pb_address_book_proto_init() {
	if File_pb_address_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_address_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_address_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadContactRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_address_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteContactRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_address_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_address_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_address_book_proto_goTypes,
		DependencyIndexes: file_pb_address_book_proto_depIdxs,
		MessageInfos:      file_pb_address_book_proto_msgTypes,
	}.Build()
	File_pb_address_book_proto = out.File
	file_pb_address_book_proto_rawDesc = nil
	file_pb_address_book_proto_goTypes = nil
	file_pb_address_book_proto_depIdxs = nil
}