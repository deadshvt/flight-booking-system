// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: bonus-service/proto/bonus.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPrivilegeWithHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetPrivilegeWithHistoryRequest) Reset() {
	*x = GetPrivilegeWithHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivilegeWithHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivilegeWithHistoryRequest) ProtoMessage() {}

func (x *GetPrivilegeWithHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivilegeWithHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetPrivilegeWithHistoryRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{0}
}

func (x *GetPrivilegeWithHistoryRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPrivilegeWithHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privilege *Privilege   `protobuf:"bytes,1,opt,name=privilege,proto3" json:"privilege,omitempty"`
	History   []*Operation `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *GetPrivilegeWithHistoryResponse) Reset() {
	*x = GetPrivilegeWithHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivilegeWithHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivilegeWithHistoryResponse) ProtoMessage() {}

func (x *GetPrivilegeWithHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivilegeWithHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetPrivilegeWithHistoryResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{1}
}

func (x *GetPrivilegeWithHistoryResponse) GetPrivilege() *Privilege {
	if x != nil {
		return x.Privilege
	}
	return nil
}

func (x *GetPrivilegeWithHistoryResponse) GetHistory() []*Operation {
	if x != nil {
		return x.History
	}
	return nil
}

type Privilege struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Balance  int32  `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Status   string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Privilege) Reset() {
	*x = Privilege{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Privilege) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Privilege) ProtoMessage() {}

func (x *Privilege) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Privilege.ProtoReflect.Descriptor instead.
func (*Privilege) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{2}
}

func (x *Privilege) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Privilege) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Privilege) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Privilege) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PrivilegeID   int32                  `protobuf:"varint,2,opt,name=privilegeID,proto3" json:"privilegeID,omitempty"`
	TicketUid     string                 `protobuf:"bytes,3,opt,name=ticketUid,proto3" json:"ticketUid,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	BalanceDiff   int32                  `protobuf:"varint,5,opt,name=balanceDiff,proto3" json:"balanceDiff,omitempty"`
	OperationType string                 `protobuf:"bytes,6,opt,name=operationType,proto3" json:"operationType,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{3}
}

func (x *Operation) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Operation) GetPrivilegeID() int32 {
	if x != nil {
		return x.PrivilegeID
	}
	return 0
}

func (x *Operation) GetTicketUid() string {
	if x != nil {
		return x.TicketUid
	}
	return ""
}

func (x *Operation) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Operation) GetBalanceDiff() int32 {
	if x != nil {
		return x.BalanceDiff
	}
	return 0
}

func (x *Operation) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

type GetPrivilegeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetPrivilegeRequest) Reset() {
	*x = GetPrivilegeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivilegeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivilegeRequest) ProtoMessage() {}

func (x *GetPrivilegeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivilegeRequest.ProtoReflect.Descriptor instead.
func (*GetPrivilegeRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{4}
}

func (x *GetPrivilegeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPrivilegeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privilege *Privilege `protobuf:"bytes,1,opt,name=privilege,proto3" json:"privilege,omitempty"`
}

func (x *GetPrivilegeResponse) Reset() {
	*x = GetPrivilegeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivilegeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivilegeResponse) ProtoMessage() {}

func (x *GetPrivilegeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivilegeResponse.ProtoReflect.Descriptor instead.
func (*GetPrivilegeResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{5}
}

func (x *GetPrivilegeResponse) GetPrivilege() *Privilege {
	if x != nil {
		return x.Privilege
	}
	return nil
}

type CreatePrivilegeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privilege *Privilege `protobuf:"bytes,1,opt,name=privilege,proto3" json:"privilege,omitempty"`
}

func (x *CreatePrivilegeRequest) Reset() {
	*x = CreatePrivilegeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrivilegeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivilegeRequest) ProtoMessage() {}

func (x *CreatePrivilegeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivilegeRequest.ProtoReflect.Descriptor instead.
func (*CreatePrivilegeRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePrivilegeRequest) GetPrivilege() *Privilege {
	if x != nil {
		return x.Privilege
	}
	return nil
}

type CreatePrivilegeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePrivilegeResponse) Reset() {
	*x = CreatePrivilegeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrivilegeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivilegeResponse) ProtoMessage() {}

func (x *CreatePrivilegeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivilegeResponse.ProtoReflect.Descriptor instead.
func (*CreatePrivilegeResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{7}
}

type UpdatePrivilegeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privilege *Privilege `protobuf:"bytes,1,opt,name=privilege,proto3" json:"privilege,omitempty"`
}

func (x *UpdatePrivilegeRequest) Reset() {
	*x = UpdatePrivilegeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrivilegeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivilegeRequest) ProtoMessage() {}

func (x *UpdatePrivilegeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivilegeRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrivilegeRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePrivilegeRequest) GetPrivilege() *Privilege {
	if x != nil {
		return x.Privilege
	}
	return nil
}

type UpdatePrivilegeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePrivilegeResponse) Reset() {
	*x = UpdatePrivilegeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrivilegeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivilegeResponse) ProtoMessage() {}

func (x *UpdatePrivilegeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivilegeResponse.ProtoReflect.Descriptor instead.
func (*UpdatePrivilegeResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{9}
}

type CreateOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *CreateOperationRequest) Reset() {
	*x = CreateOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOperationRequest) ProtoMessage() {}

func (x *CreateOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOperationRequest.ProtoReflect.Descriptor instead.
func (*CreateOperationRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{10}
}

func (x *CreateOperationRequest) GetOperation() *Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type CreateOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOperationResponse) Reset() {
	*x = CreateOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_bonus_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOperationResponse) ProtoMessage() {}

func (x *CreateOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_bonus_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOperationResponse.ProtoReflect.Descriptor instead.
func (*CreateOperationResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_bonus_proto_rawDescGZIP(), []int{11}
}

var File_bonus_service_proto_bonus_proto protoreflect.FileDescriptor

var file_bonus_service_proto_bonus_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xd3, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x66, 0x66, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x69, 0x66,
	0x66, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xb7, 0x03, 0x0a, 0x0c, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x68, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x2e,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x61, 0x64, 0x73, 0x68,
	0x76, 0x74, 0x2f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bonus_service_proto_bonus_proto_rawDescOnce sync.Once
	file_bonus_service_proto_bonus_proto_rawDescData = file_bonus_service_proto_bonus_proto_rawDesc
)

func file_bonus_service_proto_bonus_proto_rawDescGZIP() []byte {
	file_bonus_service_proto_bonus_proto_rawDescOnce.Do(func() {
		file_bonus_service_proto_bonus_proto_rawDescData = protoimpl.X.CompressGZIP(file_bonus_service_proto_bonus_proto_rawDescData)
	})
	return file_bonus_service_proto_bonus_proto_rawDescData
}

var file_bonus_service_proto_bonus_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_bonus_service_proto_bonus_proto_goTypes = []any{
	(*GetPrivilegeWithHistoryRequest)(nil),  // 0: bonus.GetPrivilegeWithHistoryRequest
	(*GetPrivilegeWithHistoryResponse)(nil), // 1: bonus.GetPrivilegeWithHistoryResponse
	(*Privilege)(nil),                       // 2: bonus.Privilege
	(*Operation)(nil),                       // 3: bonus.Operation
	(*GetPrivilegeRequest)(nil),             // 4: bonus.GetPrivilegeRequest
	(*GetPrivilegeResponse)(nil),            // 5: bonus.GetPrivilegeResponse
	(*CreatePrivilegeRequest)(nil),          // 6: bonus.CreatePrivilegeRequest
	(*CreatePrivilegeResponse)(nil),         // 7: bonus.CreatePrivilegeResponse
	(*UpdatePrivilegeRequest)(nil),          // 8: bonus.UpdatePrivilegeRequest
	(*UpdatePrivilegeResponse)(nil),         // 9: bonus.UpdatePrivilegeResponse
	(*CreateOperationRequest)(nil),          // 10: bonus.CreateOperationRequest
	(*CreateOperationResponse)(nil),         // 11: bonus.CreateOperationResponse
	(*timestamppb.Timestamp)(nil),           // 12: google.protobuf.Timestamp
}
var file_bonus_service_proto_bonus_proto_depIdxs = []int32{
	2,  // 0: bonus.GetPrivilegeWithHistoryResponse.privilege:type_name -> bonus.Privilege
	3,  // 1: bonus.GetPrivilegeWithHistoryResponse.history:type_name -> bonus.Operation
	12, // 2: bonus.Operation.date:type_name -> google.protobuf.Timestamp
	2,  // 3: bonus.GetPrivilegeResponse.privilege:type_name -> bonus.Privilege
	2,  // 4: bonus.CreatePrivilegeRequest.privilege:type_name -> bonus.Privilege
	2,  // 5: bonus.UpdatePrivilegeRequest.privilege:type_name -> bonus.Privilege
	3,  // 6: bonus.CreateOperationRequest.operation:type_name -> bonus.Operation
	0,  // 7: bonus.BonusService.GetPrivilegeWithHistory:input_type -> bonus.GetPrivilegeWithHistoryRequest
	4,  // 8: bonus.BonusService.GetPrivilege:input_type -> bonus.GetPrivilegeRequest
	6,  // 9: bonus.BonusService.CreatePrivilege:input_type -> bonus.CreatePrivilegeRequest
	8,  // 10: bonus.BonusService.UpdatePrivilege:input_type -> bonus.UpdatePrivilegeRequest
	10, // 11: bonus.BonusService.CreateOperation:input_type -> bonus.CreateOperationRequest
	1,  // 12: bonus.BonusService.GetPrivilegeWithHistory:output_type -> bonus.GetPrivilegeWithHistoryResponse
	5,  // 13: bonus.BonusService.GetPrivilege:output_type -> bonus.GetPrivilegeResponse
	7,  // 14: bonus.BonusService.CreatePrivilege:output_type -> bonus.CreatePrivilegeResponse
	9,  // 15: bonus.BonusService.UpdatePrivilege:output_type -> bonus.UpdatePrivilegeResponse
	11, // 16: bonus.BonusService.CreateOperation:output_type -> bonus.CreateOperationResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_bonus_service_proto_bonus_proto_init() }
func file_bonus_service_proto_bonus_proto_init() {
	if File_bonus_service_proto_bonus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bonus_service_proto_bonus_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPrivilegeWithHistoryRequest); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetPrivilegeWithHistoryResponse); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Privilege); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Operation); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPrivilegeRequest); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetPrivilegeResponse); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePrivilegeRequest); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePrivilegeResponse); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePrivilegeRequest); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePrivilegeResponse); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOperationRequest); i {
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
		file_bonus_service_proto_bonus_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOperationResponse); i {
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
			RawDescriptor: file_bonus_service_proto_bonus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bonus_service_proto_bonus_proto_goTypes,
		DependencyIndexes: file_bonus_service_proto_bonus_proto_depIdxs,
		MessageInfos:      file_bonus_service_proto_bonus_proto_msgTypes,
	}.Build()
	File_bonus_service_proto_bonus_proto = out.File
	file_bonus_service_proto_bonus_proto_rawDesc = nil
	file_bonus_service_proto_bonus_proto_goTypes = nil
	file_bonus_service_proto_bonus_proto_depIdxs = nil
}
