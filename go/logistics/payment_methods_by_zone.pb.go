// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: proto/logistics/payment_methods_by_zone.proto

package logistics

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentMethodByZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32                                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ZoneId             int32                                    `protobuf:"varint,2,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	DeliveryMethodId   int32                                    `protobuf:"varint,3,opt,name=delivery_method_id,json=deliveryMethodId,proto3" json:"delivery_method_id,omitempty"`
	PaymentMethodId    int32                                    `protobuf:"varint,4,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	Created            string                                   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated            string                                   `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	TransportCompanies []*PaymentMethodByZoneToTransportCompany `protobuf:"bytes,7,rep,name=transport_companies,json=transportCompanies,proto3" json:"transport_companies,omitempty"`
}

func (x *PaymentMethodByZone) Reset() {
	*x = PaymentMethodByZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodByZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodByZone) ProtoMessage() {}

func (x *PaymentMethodByZone) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodByZone.ProtoReflect.Descriptor instead.
func (*PaymentMethodByZone) Descriptor() ([]byte, []int) {
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentMethodByZone) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentMethodByZone) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *PaymentMethodByZone) GetDeliveryMethodId() int32 {
	if x != nil {
		return x.DeliveryMethodId
	}
	return 0
}

func (x *PaymentMethodByZone) GetPaymentMethodId() int32 {
	if x != nil {
		return x.PaymentMethodId
	}
	return 0
}

func (x *PaymentMethodByZone) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *PaymentMethodByZone) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *PaymentMethodByZone) GetTransportCompanies() []*PaymentMethodByZoneToTransportCompany {
	if x != nil {
		return x.TransportCompanies
	}
	return nil
}

type PaymentMethodByZoneToTransportCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodByZoneId int32 `protobuf:"varint,1,opt,name=payment_method_by_zone_id,json=paymentMethodByZoneId,proto3" json:"payment_method_by_zone_id,omitempty"`
	TransportCompanyId    int32 `protobuf:"varint,2,opt,name=transport_company_id,json=transportCompanyId,proto3" json:"transport_company_id,omitempty"`
}

func (x *PaymentMethodByZoneToTransportCompany) Reset() {
	*x = PaymentMethodByZoneToTransportCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodByZoneToTransportCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodByZoneToTransportCompany) ProtoMessage() {}

func (x *PaymentMethodByZoneToTransportCompany) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodByZoneToTransportCompany.ProtoReflect.Descriptor instead.
func (*PaymentMethodByZoneToTransportCompany) Descriptor() ([]byte, []int) {
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethodByZoneToTransportCompany) GetPaymentMethodByZoneId() int32 {
	if x != nil {
		return x.PaymentMethodByZoneId
	}
	return 0
}

func (x *PaymentMethodByZoneToTransportCompany) GetTransportCompanyId() int32 {
	if x != nil {
		return x.TransportCompanyId
	}
	return 0
}

type PaymentMethodByZoneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaymentMethodByZoneId) Reset() {
	*x = PaymentMethodByZoneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethodByZoneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethodByZoneId) ProtoMessage() {}

func (x *PaymentMethodByZoneId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethodByZoneId.ProtoReflect.Descriptor instead.
func (*PaymentMethodByZoneId) Descriptor() ([]byte, []int) {
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentMethodByZoneId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListPaymentMethodsByZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListPaymentMethodsByZoneRequest) Reset() {
	*x = ListPaymentMethodsByZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentMethodsByZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentMethodsByZoneRequest) ProtoMessage() {}

func (x *ListPaymentMethodsByZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentMethodsByZoneRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentMethodsByZoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP(), []int{3}
}

func (x *ListPaymentMethodsByZoneRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPaymentMethodsByZoneRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListPaymentMethodsByZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PaymentMethodByZone `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListPaymentMethodsByZoneResponse) Reset() {
	*x = ListPaymentMethodsByZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentMethodsByZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentMethodsByZoneResponse) ProtoMessage() {}

func (x *ListPaymentMethodsByZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_payment_methods_by_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentMethodsByZoneResponse.ProtoReflect.Descriptor instead.
func (*ListPaymentMethodsByZoneResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP(), []int{4}
}

func (x *ListPaymentMethodsByZoneResponse) GetResults() []*PaymentMethodByZone {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListPaymentMethodsByZoneResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_logistics_payment_methods_by_zone_proto protoreflect.FileDescriptor

var file_proto_logistics_payment_methods_by_zone_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x5f, 0x62, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x61, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e,
	0x65, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x25, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x54,
	0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x38, 0x0a, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x27, 0x0a,
	0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x72, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xdb, 0x05, 0x0a, 0x14,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x79,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x72, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x20,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2d, 0x62, 0x79,
	0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x64, 0x1a, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e,
	0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2d, 0x62,
	0x79, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42,
	0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2d, 0x62, 0x79, 0x2d, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x75, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x1e, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x32, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2d, 0x62, 0x79, 0x2d, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x06, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a,
	0x6f, 0x6e, 0x65, 0x1a, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x79, 0x5a,
	0x6f, 0x6e, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x1a, 0x1b, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x2d, 0x62, 0x79, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42,
	0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2d, 0x62, 0x79, 0x2d,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_logistics_payment_methods_by_zone_proto_rawDescOnce sync.Once
	file_proto_logistics_payment_methods_by_zone_proto_rawDescData = file_proto_logistics_payment_methods_by_zone_proto_rawDesc
)

func file_proto_logistics_payment_methods_by_zone_proto_rawDescGZIP() []byte {
	file_proto_logistics_payment_methods_by_zone_proto_rawDescOnce.Do(func() {
		file_proto_logistics_payment_methods_by_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_logistics_payment_methods_by_zone_proto_rawDescData)
	})
	return file_proto_logistics_payment_methods_by_zone_proto_rawDescData
}

var file_proto_logistics_payment_methods_by_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_logistics_payment_methods_by_zone_proto_goTypes = []interface{}{
	(*PaymentMethodByZone)(nil),                   // 0: logistics.PaymentMethodByZone
	(*PaymentMethodByZoneToTransportCompany)(nil), // 1: logistics.PaymentMethodByZoneToTransportCompany
	(*PaymentMethodByZoneId)(nil),                 // 2: logistics.PaymentMethodByZoneId
	(*ListPaymentMethodsByZoneRequest)(nil),       // 3: logistics.ListPaymentMethodsByZoneRequest
	(*ListPaymentMethodsByZoneResponse)(nil),      // 4: logistics.ListPaymentMethodsByZoneResponse
	(*emptypb.Empty)(nil),                         // 5: google.protobuf.Empty
}
var file_proto_logistics_payment_methods_by_zone_proto_depIdxs = []int32{
	1, // 0: logistics.PaymentMethodByZone.transport_companies:type_name -> logistics.PaymentMethodByZoneToTransportCompany
	0, // 1: logistics.ListPaymentMethodsByZoneResponse.results:type_name -> logistics.PaymentMethodByZone
	0, // 2: logistics.PaymentMethodsByZone.Create:input_type -> logistics.PaymentMethodByZone
	2, // 3: logistics.PaymentMethodsByZone.Get:input_type -> logistics.PaymentMethodByZoneId
	3, // 4: logistics.PaymentMethodsByZone.List:input_type -> logistics.ListPaymentMethodsByZoneRequest
	0, // 5: logistics.PaymentMethodsByZone.Update:input_type -> logistics.PaymentMethodByZone
	0, // 6: logistics.PaymentMethodsByZone.Upsert:input_type -> logistics.PaymentMethodByZone
	2, // 7: logistics.PaymentMethodsByZone.Delete:input_type -> logistics.PaymentMethodByZoneId
	2, // 8: logistics.PaymentMethodsByZone.Create:output_type -> logistics.PaymentMethodByZoneId
	0, // 9: logistics.PaymentMethodsByZone.Get:output_type -> logistics.PaymentMethodByZone
	4, // 10: logistics.PaymentMethodsByZone.List:output_type -> logistics.ListPaymentMethodsByZoneResponse
	0, // 11: logistics.PaymentMethodsByZone.Update:output_type -> logistics.PaymentMethodByZone
	0, // 12: logistics.PaymentMethodsByZone.Upsert:output_type -> logistics.PaymentMethodByZone
	5, // 13: logistics.PaymentMethodsByZone.Delete:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_logistics_payment_methods_by_zone_proto_init() }
func file_proto_logistics_payment_methods_by_zone_proto_init() {
	if File_proto_logistics_payment_methods_by_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_logistics_payment_methods_by_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodByZone); i {
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
		file_proto_logistics_payment_methods_by_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodByZoneToTransportCompany); i {
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
		file_proto_logistics_payment_methods_by_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethodByZoneId); i {
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
		file_proto_logistics_payment_methods_by_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentMethodsByZoneRequest); i {
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
		file_proto_logistics_payment_methods_by_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentMethodsByZoneResponse); i {
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
			RawDescriptor: file_proto_logistics_payment_methods_by_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_logistics_payment_methods_by_zone_proto_goTypes,
		DependencyIndexes: file_proto_logistics_payment_methods_by_zone_proto_depIdxs,
		MessageInfos:      file_proto_logistics_payment_methods_by_zone_proto_msgTypes,
	}.Build()
	File_proto_logistics_payment_methods_by_zone_proto = out.File
	file_proto_logistics_payment_methods_by_zone_proto_rawDesc = nil
	file_proto_logistics_payment_methods_by_zone_proto_goTypes = nil
	file_proto_logistics_payment_methods_by_zone_proto_depIdxs = nil
}
