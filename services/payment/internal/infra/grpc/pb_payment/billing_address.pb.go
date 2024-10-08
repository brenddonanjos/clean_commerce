// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/billing_address.proto

package pb_payment

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type BillingAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Street    string               `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	Number    string               `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	ZipCode   string               `protobuf:"bytes,4,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City      string               `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Country   string               `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	State     string               `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	UserId    int32                `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BillingAddressResponse) Reset() {
	*x = BillingAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_billing_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingAddressResponse) ProtoMessage() {}

func (x *BillingAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_billing_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingAddressResponse.ProtoReflect.Descriptor instead.
func (*BillingAddressResponse) Descriptor() ([]byte, []int) {
	return file_proto_billing_address_proto_rawDescGZIP(), []int{0}
}

func (x *BillingAddressResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BillingAddressResponse) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *BillingAddressResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *BillingAddressResponse) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *BillingAddressResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *BillingAddressResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *BillingAddressResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *BillingAddressResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BillingAddressResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BillingAddressResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BillingAddressResponse) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type CreateBillingAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Number  string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	ZipCode string `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	State   string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	UserId  int32  `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateBillingAddressRequest) Reset() {
	*x = CreateBillingAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_billing_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBillingAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBillingAddressRequest) ProtoMessage() {}

func (x *CreateBillingAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_billing_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBillingAddressRequest.ProtoReflect.Descriptor instead.
func (*CreateBillingAddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_billing_address_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBillingAddressRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CreateBillingAddressRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BillingAddressesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingAddresses []*BillingAddressResponse `protobuf:"bytes,1,rep,name=billing_addresses,json=billingAddresses,proto3" json:"billing_addresses,omitempty"`
}

func (x *BillingAddressesResponse) Reset() {
	*x = BillingAddressesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_billing_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingAddressesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingAddressesResponse) ProtoMessage() {}

func (x *BillingAddressesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_billing_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingAddressesResponse.ProtoReflect.Descriptor instead.
func (*BillingAddressesResponse) Descriptor() ([]byte, []int) {
	return file_proto_billing_address_proto_rawDescGZIP(), []int{2}
}

func (x *BillingAddressesResponse) GetBillingAddresses() []*BillingAddressResponse {
	if x != nil {
		return x.BillingAddresses
	}
	return nil
}

var File_proto_billing_address_proto protoreflect.FileDescriptor

var file_proto_billing_address_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x16, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc5, 0x01, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x18, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4f, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x62,
	0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x32, 0x9a, 0x01, 0x0a, 0x15, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x20,
	0x5a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_billing_address_proto_rawDescOnce sync.Once
	file_proto_billing_address_proto_rawDescData = file_proto_billing_address_proto_rawDesc
)

func file_proto_billing_address_proto_rawDescGZIP() []byte {
	file_proto_billing_address_proto_rawDescOnce.Do(func() {
		file_proto_billing_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_billing_address_proto_rawDescData)
	})
	return file_proto_billing_address_proto_rawDescData
}

var file_proto_billing_address_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_billing_address_proto_goTypes = []any{
	(*BillingAddressResponse)(nil),      // 0: pb_payment.BillingAddressResponse
	(*CreateBillingAddressRequest)(nil), // 1: pb_payment.CreateBillingAddressRequest
	(*BillingAddressesResponse)(nil),    // 2: pb_payment.BillingAddressesResponse
	(*timestamp.Timestamp)(nil),         // 3: google.protobuf.Timestamp
}
var file_proto_billing_address_proto_depIdxs = []int32{
	3, // 0: pb_payment.BillingAddressResponse.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: pb_payment.BillingAddressResponse.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: pb_payment.BillingAddressResponse.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 3: pb_payment.BillingAddressesResponse.billing_addresses:type_name -> pb_payment.BillingAddressResponse
	1, // 4: pb_payment.BillingAddressService.CreateBillingAddress:input_type -> pb_payment.CreateBillingAddressRequest
	0, // 5: pb_payment.BillingAddressService.CreateBillingAddress:output_type -> pb_payment.BillingAddressResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_billing_address_proto_init() }
func file_proto_billing_address_proto_init() {
	if File_proto_billing_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_billing_address_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BillingAddressResponse); i {
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
		file_proto_billing_address_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBillingAddressRequest); i {
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
		file_proto_billing_address_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BillingAddressesResponse); i {
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
			RawDescriptor: file_proto_billing_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_billing_address_proto_goTypes,
		DependencyIndexes: file_proto_billing_address_proto_depIdxs,
		MessageInfos:      file_proto_billing_address_proto_msgTypes,
	}.Build()
	File_proto_billing_address_proto = out.File
	file_proto_billing_address_proto_rawDesc = nil
	file_proto_billing_address_proto_goTypes = nil
	file_proto_billing_address_proto_depIdxs = nil
}
