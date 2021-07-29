// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: proto/logistics/delivery_manual_priority.proto

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

type DeliveryManualToTransportCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryManualPriorityId int32   `protobuf:"varint,1,opt,name=delivery_manual_priority_id,json=deliveryManualPriorityId,proto3" json:"delivery_manual_priority_id,omitempty"`
	TransportCompanyId       int32   `protobuf:"varint,2,opt,name=transport_company_id,json=transportCompanyId,proto3" json:"transport_company_id,omitempty"`
	Value                    float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DeliveryManualToTransportCompany) Reset() {
	*x = DeliveryManualToTransportCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryManualToTransportCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryManualToTransportCompany) ProtoMessage() {}

func (x *DeliveryManualToTransportCompany) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryManualToTransportCompany.ProtoReflect.Descriptor instead.
func (*DeliveryManualToTransportCompany) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryManualToTransportCompany) GetDeliveryManualPriorityId() int32 {
	if x != nil {
		return x.DeliveryManualPriorityId
	}
	return 0
}

func (x *DeliveryManualToTransportCompany) GetTransportCompanyId() int32 {
	if x != nil {
		return x.TransportCompanyId
	}
	return 0
}

func (x *DeliveryManualToTransportCompany) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DeliveryManualPriorityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeliveryManualPriorityId) Reset() {
	*x = DeliveryManualPriorityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryManualPriorityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryManualPriorityId) ProtoMessage() {}

func (x *DeliveryManualPriorityId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryManualPriorityId.ProtoReflect.Descriptor instead.
func (*DeliveryManualPriorityId) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryManualPriorityId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeliveryManualPriority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ZoneId                   int32                    `protobuf:"varint,2,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	DeliveryMethodId         int32                    `protobuf:"varint,3,opt,name=delivery_method_id,json=deliveryMethodId,proto3" json:"delivery_method_id,omitempty"`
	MaxDays                  float64                  `protobuf:"fixed64,4,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	Tariff                   float64                  `protobuf:"fixed64,5,opt,name=tariff,proto3" json:"tariff,omitempty"`
	Created                  string                   `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated                  string                   `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	TransportCompaniesValues []*TransportCompanyValue `protobuf:"bytes,8,rep,name=transport_companies_values,json=transportCompaniesValues,proto3" json:"transport_companies_values,omitempty"`
}

func (x *DeliveryManualPriority) Reset() {
	*x = DeliveryManualPriority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryManualPriority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryManualPriority) ProtoMessage() {}

func (x *DeliveryManualPriority) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryManualPriority.ProtoReflect.Descriptor instead.
func (*DeliveryManualPriority) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryManualPriority) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryManualPriority) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *DeliveryManualPriority) GetDeliveryMethodId() int32 {
	if x != nil {
		return x.DeliveryMethodId
	}
	return 0
}

func (x *DeliveryManualPriority) GetMaxDays() float64 {
	if x != nil {
		return x.MaxDays
	}
	return 0
}

func (x *DeliveryManualPriority) GetTariff() float64 {
	if x != nil {
		return x.Tariff
	}
	return 0
}

func (x *DeliveryManualPriority) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *DeliveryManualPriority) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *DeliveryManualPriority) GetTransportCompaniesValues() []*TransportCompanyValue {
	if x != nil {
		return x.TransportCompaniesValues
	}
	return nil
}

type TransportCompanyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransportCompanyId int32   `protobuf:"varint,1,opt,name=transport_company_id,json=transportCompanyId,proto3" json:"transport_company_id,omitempty"`
	Value              float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TransportCompanyValue) Reset() {
	*x = TransportCompanyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportCompanyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportCompanyValue) ProtoMessage() {}

func (x *TransportCompanyValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportCompanyValue.ProtoReflect.Descriptor instead.
func (*TransportCompanyValue) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{3}
}

func (x *TransportCompanyValue) GetTransportCompanyId() int32 {
	if x != nil {
		return x.TransportCompanyId
	}
	return 0
}

func (x *TransportCompanyValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ListDeliveriesManualPrioritiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListDeliveriesManualPrioritiesRequest) Reset() {
	*x = ListDeliveriesManualPrioritiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeliveriesManualPrioritiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeliveriesManualPrioritiesRequest) ProtoMessage() {}

func (x *ListDeliveriesManualPrioritiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeliveriesManualPrioritiesRequest.ProtoReflect.Descriptor instead.
func (*ListDeliveriesManualPrioritiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{4}
}

func (x *ListDeliveriesManualPrioritiesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDeliveriesManualPrioritiesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListDeliveriesManualPrioritiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DeliveryManualPriority `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListDeliveriesManualPrioritiesResponse) Reset() {
	*x = ListDeliveriesManualPrioritiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeliveriesManualPrioritiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeliveriesManualPrioritiesResponse) ProtoMessage() {}

func (x *ListDeliveriesManualPrioritiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_delivery_manual_priority_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeliveriesManualPrioritiesResponse.ProtoReflect.Descriptor instead.
func (*ListDeliveriesManualPrioritiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP(), []int{5}
}

func (x *ListDeliveriesManualPrioritiesResponse) GetResults() []*DeliveryManualPriority {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListDeliveriesManualPrioritiesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_logistics_delivery_manual_priority_proto protoreflect.FileDescriptor

var file_proto_logistics_delivery_manual_priority_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3d, 0x0a, 0x1b, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x18, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb6,
	0x02, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x5e, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x30, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x25, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x7b, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xd9, 0x09, 0x0a,
	0x1a, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x7d, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x2b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x23, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x95, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x30, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x75,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x80, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x21,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x32, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x27, 0x2a, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a,
	0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x1a, 0x25, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xce, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x2b, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x5d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x57, 0x22, 0x52,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xd0, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x71, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x6b, 0x2a, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x7b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_logistics_delivery_manual_priority_proto_rawDescOnce sync.Once
	file_proto_logistics_delivery_manual_priority_proto_rawDescData = file_proto_logistics_delivery_manual_priority_proto_rawDesc
)

func file_proto_logistics_delivery_manual_priority_proto_rawDescGZIP() []byte {
	file_proto_logistics_delivery_manual_priority_proto_rawDescOnce.Do(func() {
		file_proto_logistics_delivery_manual_priority_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_logistics_delivery_manual_priority_proto_rawDescData)
	})
	return file_proto_logistics_delivery_manual_priority_proto_rawDescData
}

var file_proto_logistics_delivery_manual_priority_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_logistics_delivery_manual_priority_proto_goTypes = []interface{}{
	(*DeliveryManualToTransportCompany)(nil),       // 0: logistics.DeliveryManualToTransportCompany
	(*DeliveryManualPriorityId)(nil),               // 1: logistics.DeliveryManualPriorityId
	(*DeliveryManualPriority)(nil),                 // 2: logistics.DeliveryManualPriority
	(*TransportCompanyValue)(nil),                  // 3: logistics.TransportCompanyValue
	(*ListDeliveriesManualPrioritiesRequest)(nil),  // 4: logistics.ListDeliveriesManualPrioritiesRequest
	(*ListDeliveriesManualPrioritiesResponse)(nil), // 5: logistics.ListDeliveriesManualPrioritiesResponse
	(*emptypb.Empty)(nil),                          // 6: google.protobuf.Empty
}
var file_proto_logistics_delivery_manual_priority_proto_depIdxs = []int32{
	3,  // 0: logistics.DeliveryManualPriority.transport_companies_values:type_name -> logistics.TransportCompanyValue
	2,  // 1: logistics.ListDeliveriesManualPrioritiesResponse.results:type_name -> logistics.DeliveryManualPriority
	2,  // 2: logistics.DeliveriesManualPriorities.Create:input_type -> logistics.DeliveryManualPriority
	1,  // 3: logistics.DeliveriesManualPriorities.Get:input_type -> logistics.DeliveryManualPriorityId
	4,  // 4: logistics.DeliveriesManualPriorities.List:input_type -> logistics.ListDeliveriesManualPrioritiesRequest
	2,  // 5: logistics.DeliveriesManualPriorities.Update:input_type -> logistics.DeliveryManualPriority
	1,  // 6: logistics.DeliveriesManualPriorities.Delete:input_type -> logistics.DeliveryManualPriorityId
	2,  // 7: logistics.DeliveriesManualPriorities.CreateOrUpdate:input_type -> logistics.DeliveryManualPriority
	0,  // 8: logistics.DeliveriesManualPriorities.AddTransportCompany:input_type -> logistics.DeliveryManualToTransportCompany
	0,  // 9: logistics.DeliveriesManualPriorities.DeleteTransportCompany:input_type -> logistics.DeliveryManualToTransportCompany
	1,  // 10: logistics.DeliveriesManualPriorities.Create:output_type -> logistics.DeliveryManualPriorityId
	2,  // 11: logistics.DeliveriesManualPriorities.Get:output_type -> logistics.DeliveryManualPriority
	5,  // 12: logistics.DeliveriesManualPriorities.List:output_type -> logistics.ListDeliveriesManualPrioritiesResponse
	2,  // 13: logistics.DeliveriesManualPriorities.Update:output_type -> logistics.DeliveryManualPriority
	6,  // 14: logistics.DeliveriesManualPriorities.Delete:output_type -> google.protobuf.Empty
	2,  // 15: logistics.DeliveriesManualPriorities.CreateOrUpdate:output_type -> logistics.DeliveryManualPriority
	0,  // 16: logistics.DeliveriesManualPriorities.AddTransportCompany:output_type -> logistics.DeliveryManualToTransportCompany
	6,  // 17: logistics.DeliveriesManualPriorities.DeleteTransportCompany:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_logistics_delivery_manual_priority_proto_init() }
func file_proto_logistics_delivery_manual_priority_proto_init() {
	if File_proto_logistics_delivery_manual_priority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryManualToTransportCompany); i {
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
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryManualPriorityId); i {
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
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryManualPriority); i {
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
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportCompanyValue); i {
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
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeliveriesManualPrioritiesRequest); i {
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
		file_proto_logistics_delivery_manual_priority_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeliveriesManualPrioritiesResponse); i {
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
			RawDescriptor: file_proto_logistics_delivery_manual_priority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_logistics_delivery_manual_priority_proto_goTypes,
		DependencyIndexes: file_proto_logistics_delivery_manual_priority_proto_depIdxs,
		MessageInfos:      file_proto_logistics_delivery_manual_priority_proto_msgTypes,
	}.Build()
	File_proto_logistics_delivery_manual_priority_proto = out.File
	file_proto_logistics_delivery_manual_priority_proto_rawDesc = nil
	file_proto_logistics_delivery_manual_priority_proto_goTypes = nil
	file_proto_logistics_delivery_manual_priority_proto_depIdxs = nil
}
