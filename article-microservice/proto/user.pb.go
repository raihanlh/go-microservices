// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user.proto

package proto

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day   int32 `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Year  int32 `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdAccount   int64                `protobuf:"varint,2,opt,name=id_account,json=idAccount,proto3" json:"id_account,omitempty"`
	Fullname    string               `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	IdGender    int64                `protobuf:"varint,4,opt,name=id_gender,json=idGender,proto3" json:"id_gender,omitempty"`
	Gender      string               `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Phone       string               `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	DateOfBirth *Date                `protobuf:"bytes,7,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDetail) GetIdAccount() int64 {
	if x != nil {
		return x.IdAccount
	}
	return 0
}

func (x *UserDetail) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UserDetail) GetIdGender() int64 {
	if x != nil {
		return x.IdGender
	}
	return 0
}

func (x *UserDetail) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserDetail) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserDetail) GetDateOfBirth() *Date {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *UserDetail) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserDetail) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateUpdateUserDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname    string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	IdGender    int64  `protobuf:"varint,2,opt,name=id_gender,json=idGender,proto3" json:"id_gender,omitempty"`
	Phone       string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	DateOfBirth *Date  `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Token       string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateUpdateUserDetailRequest) Reset() {
	*x = CreateUpdateUserDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUpdateUserDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUpdateUserDetailRequest) ProtoMessage() {}

func (x *CreateUpdateUserDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUpdateUserDetailRequest.ProtoReflect.Descriptor instead.
func (*CreateUpdateUserDetailRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUpdateUserDetailRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateUpdateUserDetailRequest) GetIdGender() int64 {
	if x != nil {
		return x.IdGender
	}
	return 0
}

func (x *CreateUpdateUserDetailRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUpdateUserDetailRequest) GetDateOfBirth() *Date {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *CreateUpdateUserDetailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetUserDetailByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetUserDetailByUserRequest) Reset() {
	*x = GetUserDetailByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailByUserRequest) ProtoMessage() {}

func (x *GetUserDetailByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailByUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserDetailByUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserDetailByUserRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetAllUserDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDetails []*UserDetail `protobuf:"bytes,1,rep,name=user_details,json=userDetails,proto3" json:"user_details,omitempty"`
}

func (x *GetAllUserDetailResponse) Reset() {
	*x = GetAllUserDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserDetailResponse) ProtoMessage() {}

func (x *GetAllUserDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserDetailResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserDetailResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllUserDetailResponse) GetUserDetails() []*UserDetail {
	if x != nil {
		return x.UserDetails
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x22, 0xc9, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x64, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb5, 0x01,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x64, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2f,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0xfc, 0x01, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x69, 0x68, 0x61, 0x6e, 0x6c,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(*Empty)(nil),                         // 0: proto.Empty
	(*Date)(nil),                          // 1: proto.Date
	(*UserDetail)(nil),                    // 2: proto.UserDetail
	(*CreateUpdateUserDetailRequest)(nil), // 3: proto.CreateUpdateUserDetailRequest
	(*GetUserDetailByUserRequest)(nil),    // 4: proto.GetUserDetailByUserRequest
	(*GetAllUserDetailResponse)(nil),      // 5: proto.GetAllUserDetailResponse
	(*timestamp.Timestamp)(nil),           // 6: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	1, // 0: proto.UserDetail.date_of_birth:type_name -> proto.Date
	6, // 1: proto.UserDetail.created_at:type_name -> google.protobuf.Timestamp
	6, // 2: proto.UserDetail.updated_at:type_name -> google.protobuf.Timestamp
	1, // 3: proto.CreateUpdateUserDetailRequest.date_of_birth:type_name -> proto.Date
	2, // 4: proto.GetAllUserDetailResponse.user_details:type_name -> proto.UserDetail
	3, // 5: proto.UserDetailService.CreateUpdateUserDetail:input_type -> proto.CreateUpdateUserDetailRequest
	4, // 6: proto.UserDetailService.GetUserDetailByUser:input_type -> proto.GetUserDetailByUserRequest
	0, // 7: proto.UserDetailService.GetAllUserDetail:input_type -> proto.Empty
	2, // 8: proto.UserDetailService.CreateUpdateUserDetail:output_type -> proto.UserDetail
	2, // 9: proto.UserDetailService.GetUserDetailByUser:output_type -> proto.UserDetail
	5, // 10: proto.UserDetailService.GetAllUserDetail:output_type -> proto.GetAllUserDetailResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUpdateUserDetailRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailByUserRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserDetailResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
