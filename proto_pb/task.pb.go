// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto_pb/task.proto

package proto_pb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone int64  `protobuf:"varint,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *UserReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *UserRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityType string `protobuf:"bytes,1,opt,name=ActivityType,proto3" json:"ActivityType,omitempty"`
	TimeStamp    string `protobuf:"bytes,2,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	Duration     int32  `protobuf:"varint,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Label        string `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
	Name         string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{4}
}

func (x *Activity) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *Activity) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *Activity) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Activity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Activity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ActivityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity *Activity `protobuf:"bytes,1,opt,name=Activity,proto3" json:"Activity,omitempty"`
}

func (x *ActivityReq) Reset() {
	*x = ActivityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityReq) ProtoMessage() {}

func (x *ActivityReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityReq.ProtoReflect.Descriptor instead.
func (*ActivityReq) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{5}
}

func (x *ActivityReq) GetActivity() *Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type ActivityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ActivityRes) Reset() {
	*x = ActivityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityRes) ProtoMessage() {}

func (x *ActivityRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityRes.ProtoReflect.Descriptor instead.
func (*ActivityRes) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{6}
}

func (x *ActivityRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type IsValidReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ActivityType string `protobuf:"bytes,2,opt,name=ActivityType,proto3" json:"ActivityType,omitempty"`
}

func (x *IsValidReq) Reset() {
	*x = IsValidReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidReq) ProtoMessage() {}

func (x *IsValidReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidReq.ProtoReflect.Descriptor instead.
func (*IsValidReq) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{7}
}

func (x *IsValidReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IsValidReq) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

type IsValidRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *IsValidRes) Reset() {
	*x = IsValidRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidRes) ProtoMessage() {}

func (x *IsValidRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidRes.ProtoReflect.Descriptor instead.
func (*IsValidRes) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{8}
}

func (x *IsValidRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type IsDoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ActivityType string `protobuf:"bytes,2,opt,name=ActivityType,proto3" json:"ActivityType,omitempty"`
}

func (x *IsDoneReq) Reset() {
	*x = IsDoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsDoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsDoneReq) ProtoMessage() {}

func (x *IsDoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsDoneReq.ProtoReflect.Descriptor instead.
func (*IsDoneReq) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{9}
}

func (x *IsDoneReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IsDoneReq) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

type IsDoneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *IsDoneRes) Reset() {
	*x = IsDoneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsDoneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsDoneRes) ProtoMessage() {}

func (x *IsDoneRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsDoneRes.ProtoReflect.Descriptor instead.
func (*IsDoneRes) Descriptor() ([]byte, []int) {
	return file_proto_pb_task_proto_rawDescGZIP(), []int{10}
}

func (x *IsDoneRes) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_proto_pb_task_proto protoreflect.FileDescriptor

var file_proto_pb_task_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x22,
	0x46, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x3d, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x2e,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x25,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x44, 0x0a, 0x0a, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x49,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x43, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa6, 0x02, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70,
	0x62, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x06, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pb_task_proto_rawDescOnce sync.Once
	file_proto_pb_task_proto_rawDescData = file_proto_pb_task_proto_rawDesc
)

func file_proto_pb_task_proto_rawDescGZIP() []byte {
	file_proto_pb_task_proto_rawDescOnce.Do(func() {
		file_proto_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pb_task_proto_rawDescData)
	})
	return file_proto_pb_task_proto_rawDescData
}

var file_proto_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_pb_task_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: proto_pb.User
	(*UserReq)(nil),     // 1: proto_pb.UserReq
	(*UserRes)(nil),     // 2: proto_pb.UserRes
	(*UpdateReq)(nil),   // 3: proto_pb.UpdateReq
	(*Activity)(nil),    // 4: proto_pb.Activity
	(*ActivityReq)(nil), // 5: proto_pb.ActivityReq
	(*ActivityRes)(nil), // 6: proto_pb.ActivityRes
	(*IsValidReq)(nil),  // 7: proto_pb.IsValidReq
	(*IsValidRes)(nil),  // 8: proto_pb.IsValidRes
	(*IsDoneReq)(nil),   // 9: proto_pb.IsDoneReq
	(*IsDoneRes)(nil),   // 10: proto_pb.IsDoneRes
}
var file_proto_pb_task_proto_depIdxs = []int32{
	0,  // 0: proto_pb.UserReq.User:type_name -> proto_pb.User
	0,  // 1: proto_pb.UpdateReq.user:type_name -> proto_pb.User
	4,  // 2: proto_pb.ActivityReq.Activity:type_name -> proto_pb.Activity
	1,  // 3: proto_pb.UserService.AddUser:input_type -> proto_pb.UserReq
	3,  // 4: proto_pb.UserService.UpdateUser:input_type -> proto_pb.UpdateReq
	5,  // 5: proto_pb.UserService.AddActivity:input_type -> proto_pb.ActivityReq
	7,  // 6: proto_pb.UserService.IsValid:input_type -> proto_pb.IsValidReq
	9,  // 7: proto_pb.UserService.IsDone:input_type -> proto_pb.IsDoneReq
	2,  // 8: proto_pb.UserService.AddUser:output_type -> proto_pb.UserRes
	2,  // 9: proto_pb.UserService.UpdateUser:output_type -> proto_pb.UserRes
	6,  // 10: proto_pb.UserService.AddActivity:output_type -> proto_pb.ActivityRes
	8,  // 11: proto_pb.UserService.IsValid:output_type -> proto_pb.IsValidRes
	10, // 12: proto_pb.UserService.IsDone:output_type -> proto_pb.IsDoneRes
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_pb_task_proto_init() }
func file_proto_pb_task_proto_init() {
	if File_proto_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_proto_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRes); i {
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
		file_proto_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_proto_pb_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_proto_pb_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityReq); i {
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
		file_proto_pb_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityRes); i {
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
		file_proto_pb_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidReq); i {
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
		file_proto_pb_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidRes); i {
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
		file_proto_pb_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsDoneReq); i {
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
		file_proto_pb_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsDoneRes); i {
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
			RawDescriptor: file_proto_pb_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pb_task_proto_goTypes,
		DependencyIndexes: file_proto_pb_task_proto_depIdxs,
		MessageInfos:      file_proto_pb_task_proto_msgTypes,
	}.Build()
	File_proto_pb_task_proto = out.File
	file_proto_pb_task_proto_rawDesc = nil
	file_proto_pb_task_proto_goTypes = nil
	file_proto_pb_task_proto_depIdxs = nil
}