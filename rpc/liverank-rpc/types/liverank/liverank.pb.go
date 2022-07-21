// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: liverank.proto

package liverank

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

type SubmitRankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// B站uid
	Buid int64 `protobuf:"varint,1,opt,name=buid,proto3" json:"buid,omitempty"`
	// 房间号
	Roomid int64 `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
	// 排名
	Rank int64 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	// 礼物价值
	Giftvalue int64 `protobuf:"varint,4,opt,name=giftvalue,proto3" json:"giftvalue,omitempty"`
	// 是否关注
	Isconcern bool `protobuf:"varint,5,opt,name=isconcern,proto3" json:"isconcern,omitempty"`
}

func (x *SubmitRankRequest) Reset() {
	*x = SubmitRankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRankRequest) ProtoMessage() {}

func (x *SubmitRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRankRequest.ProtoReflect.Descriptor instead.
func (*SubmitRankRequest) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitRankRequest) GetBuid() int64 {
	if x != nil {
		return x.Buid
	}
	return 0
}

func (x *SubmitRankRequest) GetRoomid() int64 {
	if x != nil {
		return x.Roomid
	}
	return 0
}

func (x *SubmitRankRequest) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *SubmitRankRequest) GetGiftvalue() int64 {
	if x != nil {
		return x.Giftvalue
	}
	return 0
}

func (x *SubmitRankRequest) GetIsconcern() bool {
	if x != nil {
		return x.Isconcern
	}
	return false
}

type SubmitRankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 返回信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitRankResponse) Reset() {
	*x = SubmitRankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRankResponse) ProtoMessage() {}

func (x *SubmitRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRankResponse.ProtoReflect.Descriptor instead.
func (*SubmitRankResponse) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitRankResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitRankResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetRankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 时间戳
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 房间号
	Roomid int64 `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
}

func (x *GetRankRequest) Reset() {
	*x = GetRankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankRequest) ProtoMessage() {}

func (x *GetRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankRequest.ProtoReflect.Descriptor instead.
func (*GetRankRequest) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{2}
}

func (x *GetRankRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetRankRequest) GetRoomid() int64 {
	if x != nil {
		return x.Roomid
	}
	return 0
}

// 排名列表
type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// B站uid
	Buid int64 `protobuf:"varint,1,opt,name=buid,proto3" json:"buid,omitempty"`
	// 排名
	Rank int64 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{3}
}

func (x *Rank) GetBuid() int64 {
	if x != nil {
		return x.Buid
	}
	return 0
}

func (x *Rank) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type GetRankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 返回信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 时间戳
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 排名列表
	Ranks []*Rank `protobuf:"bytes,4,rep,name=ranks,proto3" json:"ranks,omitempty"`
}

func (x *GetRankResponse) Reset() {
	*x = GetRankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankResponse) ProtoMessage() {}

func (x *GetRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankResponse.ProtoReflect.Descriptor instead.
func (*GetRankResponse) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{4}
}

func (x *GetRankResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetRankResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRankResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetRankResponse) GetRanks() []*Rank {
	if x != nil {
		return x.Ranks
	}
	return nil
}

type MarkConcernRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// B站uid
	Buid int64 `protobuf:"varint,1,opt,name=buid,proto3" json:"buid,omitempty"`
	// 房间号
	Roomid int64 `protobuf:"varint,2,opt,name=roomid,proto3" json:"roomid,omitempty"`
	// 是否关注
	Isconcern bool `protobuf:"varint,3,opt,name=isconcern,proto3" json:"isconcern,omitempty"`
}

func (x *MarkConcernRequest) Reset() {
	*x = MarkConcernRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkConcernRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkConcernRequest) ProtoMessage() {}

func (x *MarkConcernRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkConcernRequest.ProtoReflect.Descriptor instead.
func (*MarkConcernRequest) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{5}
}

func (x *MarkConcernRequest) GetBuid() int64 {
	if x != nil {
		return x.Buid
	}
	return 0
}

func (x *MarkConcernRequest) GetRoomid() int64 {
	if x != nil {
		return x.Roomid
	}
	return 0
}

func (x *MarkConcernRequest) GetIsconcern() bool {
	if x != nil {
		return x.Isconcern
	}
	return false
}

type MarkConcernResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 返回信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MarkConcernResponse) Reset() {
	*x = MarkConcernResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liverank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkConcernResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkConcernResponse) ProtoMessage() {}

func (x *MarkConcernResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liverank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkConcernResponse.ProtoReflect.Descriptor instead.
func (*MarkConcernResponse) Descriptor() ([]byte, []int) {
	return file_liverank_proto_rawDescGZIP(), []int{6}
}

func (x *MarkConcernResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MarkConcernResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_liverank_proto protoreflect.FileDescriptor

var file_liverank_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x62, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x1c, 0x0a, 0x09, 0x67, 0x69, 0x66, 0x74, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x69, 0x66, 0x74, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x22, 0x42, 0x0a, 0x12,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x62, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x6b,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61,
	0x6e, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x22, 0x5e,
	0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x62, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x22, 0x43,
	0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xe5, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x49, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1b,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x6e, 0x12, 0x1c, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x63,
	0x65, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_liverank_proto_rawDescOnce sync.Once
	file_liverank_proto_rawDescData = file_liverank_proto_rawDesc
)

func file_liverank_proto_rawDescGZIP() []byte {
	file_liverank_proto_rawDescOnce.Do(func() {
		file_liverank_proto_rawDescData = protoimpl.X.CompressGZIP(file_liverank_proto_rawDescData)
	})
	return file_liverank_proto_rawDescData
}

var file_liverank_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_liverank_proto_goTypes = []interface{}{
	(*SubmitRankRequest)(nil),   // 0: liverank.SubmitRankRequest
	(*SubmitRankResponse)(nil),  // 1: liverank.SubmitRankResponse
	(*GetRankRequest)(nil),      // 2: liverank.GetRankRequest
	(*Rank)(nil),                // 3: liverank.Rank
	(*GetRankResponse)(nil),     // 4: liverank.GetRankResponse
	(*MarkConcernRequest)(nil),  // 5: liverank.MarkConcernRequest
	(*MarkConcernResponse)(nil), // 6: liverank.MarkConcernResponse
}
var file_liverank_proto_depIdxs = []int32{
	3, // 0: liverank.GetRankResponse.ranks:type_name -> liverank.Rank
	0, // 1: liverank.Liverank.SubmitRank:input_type -> liverank.SubmitRankRequest
	2, // 2: liverank.Liverank.GetRank:input_type -> liverank.GetRankRequest
	5, // 3: liverank.Liverank.MarkConcern:input_type -> liverank.MarkConcernRequest
	1, // 4: liverank.Liverank.SubmitRank:output_type -> liverank.SubmitRankResponse
	4, // 5: liverank.Liverank.GetRank:output_type -> liverank.GetRankResponse
	6, // 6: liverank.Liverank.MarkConcern:output_type -> liverank.MarkConcernResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_liverank_proto_init() }
func file_liverank_proto_init() {
	if File_liverank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liverank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRankRequest); i {
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
		file_liverank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRankResponse); i {
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
		file_liverank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankRequest); i {
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
		file_liverank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rank); i {
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
		file_liverank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankResponse); i {
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
		file_liverank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkConcernRequest); i {
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
		file_liverank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkConcernResponse); i {
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
			RawDescriptor: file_liverank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liverank_proto_goTypes,
		DependencyIndexes: file_liverank_proto_depIdxs,
		MessageInfos:      file_liverank_proto_msgTypes,
	}.Build()
	File_liverank_proto = out.File
	file_liverank_proto_rawDesc = nil
	file_liverank_proto_goTypes = nil
	file_liverank_proto_depIdxs = nil
}