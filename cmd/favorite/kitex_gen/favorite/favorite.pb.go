// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.3
// source: favorite.proto

package favorite

import (
	context "context"
	feed "github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
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

type BloomblogFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户id
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                              // 用户鉴权token
	PostId     int64  `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`             // 视频id
	ActionType int32  `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (x *BloomblogFavoriteActionRequest) Reset() {
	*x = BloomblogFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogFavoriteActionRequest) ProtoMessage() {}

func (x *BloomblogFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*BloomblogFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *BloomblogFavoriteActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BloomblogFavoriteActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *BloomblogFavoriteActionRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *BloomblogFavoriteActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type BloomblogFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *BloomblogFavoriteActionResponse) Reset() {
	*x = BloomblogFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogFavoriteActionResponse) ProtoMessage() {}

func (x *BloomblogFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*BloomblogFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *BloomblogFavoriteActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BloomblogFavoriteActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type BloomblogFavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *BloomblogFavoriteListRequest) Reset() {
	*x = BloomblogFavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogFavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogFavoriteListRequest) ProtoMessage() {}

func (x *BloomblogFavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogFavoriteListRequest.ProtoReflect.Descriptor instead.
func (*BloomblogFavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *BloomblogFavoriteListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BloomblogFavoriteListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type BloomblogFavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string      `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	PostList   []*feed.Post `protobuf:"bytes,3,rep,name=post_list,json=postList,proto3" json:"post_list,omitempty"`          // 用户点赞视频列表
}

func (x *BloomblogFavoriteListResponse) Reset() {
	*x = BloomblogFavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomblogFavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomblogFavoriteListResponse) ProtoMessage() {}

func (x *BloomblogFavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomblogFavoriteListResponse.ProtoReflect.Descriptor instead.
func (*BloomblogFavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *BloomblogFavoriteListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BloomblogFavoriteListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *BloomblogFavoriteListResponse) GetPostList() []*feed.Post {
	if x != nil {
		return x.PostList
	}
	return nil
}

var File_favorite_proto protoreflect.FileDescriptor

var file_favorite_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x0a, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x6f, 0x6d,
	0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x22, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c,
	0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22,
	0x50, 0x0a, 0x1f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x20, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x32, 0xe1, 0x01, 0x0a, 0x0b, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x72, 0x76, 0x12, 0x6b, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x62, 0x6c,
	0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x65, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x29, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x6f,
	0x6d, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x62, 0x6c, 0x6f, 0x67,
	0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x67, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x61, 0x62,
	0x79, 0x2f, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6d, 0x64, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_favorite_proto_rawDescOnce sync.Once
	file_favorite_proto_rawDescData = file_favorite_proto_rawDesc
)

func file_favorite_proto_rawDescGZIP() []byte {
	file_favorite_proto_rawDescOnce.Do(func() {
		file_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorite_proto_rawDescData)
	})
	return file_favorite_proto_rawDescData
}

var file_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_favorite_proto_goTypes = []interface{}{
	(*BloomblogFavoriteActionRequest)(nil),  // 0: favorite.bloomblog_favorite_action_request
	(*BloomblogFavoriteActionResponse)(nil), // 1: favorite.bloomblog_favorite_action_response
	(*BloomblogFavoriteListRequest)(nil),    // 2: favorite.bloomblog_favorite_list_request
	(*BloomblogFavoriteListResponse)(nil),   // 3: favorite.bloomblog_favorite_list_response
	(*feed.Post)(nil),                       // 4: feed.Post
}
var file_favorite_proto_depIdxs = []int32{
	4, // 0: favorite.bloomblog_favorite_list_response.post_list:type_name -> feed.Post
	0, // 1: favorite.FavoriteSrv.FavoriteAction:input_type -> favorite.bloomblog_favorite_action_request
	2, // 2: favorite.FavoriteSrv.FavoriteList:input_type -> favorite.bloomblog_favorite_list_request
	1, // 3: favorite.FavoriteSrv.FavoriteAction:output_type -> favorite.bloomblog_favorite_action_response
	3, // 4: favorite.FavoriteSrv.FavoriteList:output_type -> favorite.bloomblog_favorite_list_response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_favorite_proto_init() }
func file_favorite_proto_init() {
	if File_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogFavoriteActionRequest); i {
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
		file_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogFavoriteActionResponse); i {
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
		file_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogFavoriteListRequest); i {
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
		file_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomblogFavoriteListResponse); i {
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
	file_favorite_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_favorite_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_favorite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorite_proto_goTypes,
		DependencyIndexes: file_favorite_proto_depIdxs,
		MessageInfos:      file_favorite_proto_msgTypes,
	}.Build()
	File_favorite_proto = out.File
	file_favorite_proto_rawDesc = nil
	file_favorite_proto_goTypes = nil
	file_favorite_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.11.3. DO NOT EDIT.

type FavoriteSrv interface {
	FavoriteAction(ctx context.Context, req *BloomblogFavoriteActionRequest) (res *BloomblogFavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, req *BloomblogFavoriteListRequest) (res *BloomblogFavoriteListResponse, err error)
}
