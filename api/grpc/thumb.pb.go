// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/grpc/thumb.proto

package thumbnailapi

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

type ThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoUrl string `protobuf:"bytes,1,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
}

func (x *ThumbnailRequest) Reset() {
	*x = ThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_thumb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailRequest) ProtoMessage() {}

func (x *ThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_thumb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailRequest.ProtoReflect.Descriptor instead.
func (*ThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_thumb_proto_rawDescGZIP(), []int{0}
}

func (x *ThumbnailRequest) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

type ThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThumbnailData []byte `protobuf:"bytes,1,opt,name=thumbnail_data,json=thumbnailData,proto3" json:"thumbnail_data,omitempty"`
}

func (x *ThumbnailResponse) Reset() {
	*x = ThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_thumb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailResponse) ProtoMessage() {}

func (x *ThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_thumb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailResponse.ProtoReflect.Descriptor instead.
func (*ThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_thumb_proto_rawDescGZIP(), []int{1}
}

func (x *ThumbnailResponse) GetThumbnailData() []byte {
	if x != nil {
		return x.ThumbnailData
	}
	return nil
}

var File_api_grpc_thumb_proto protoreflect.FileDescriptor

var file_api_grpc_thumb_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2f, 0x0a, 0x10, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x3a, 0x0a, 0x11,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0x51, 0x0a, 0x10, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x68, 0x65, 0x6c, 0x6f,
	0x6e, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_thumb_proto_rawDescOnce sync.Once
	file_api_grpc_thumb_proto_rawDescData = file_api_grpc_thumb_proto_rawDesc
)

func file_api_grpc_thumb_proto_rawDescGZIP() []byte {
	file_api_grpc_thumb_proto_rawDescOnce.Do(func() {
		file_api_grpc_thumb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_thumb_proto_rawDescData)
	})
	return file_api_grpc_thumb_proto_rawDescData
}

var file_api_grpc_thumb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_thumb_proto_goTypes = []any{
	(*ThumbnailRequest)(nil),  // 0: api.ThumbnailRequest
	(*ThumbnailResponse)(nil), // 1: api.ThumbnailResponse
}
var file_api_grpc_thumb_proto_depIdxs = []int32{
	0, // 0: api.ThumbnailService.GetThumbnail:input_type -> api.ThumbnailRequest
	1, // 1: api.ThumbnailService.GetThumbnail:output_type -> api.ThumbnailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_thumb_proto_init() }
func file_api_grpc_thumb_proto_init() {
	if File_api_grpc_thumb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_thumb_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbnailRequest); i {
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
		file_api_grpc_thumb_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbnailResponse); i {
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
			RawDescriptor: file_api_grpc_thumb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_thumb_proto_goTypes,
		DependencyIndexes: file_api_grpc_thumb_proto_depIdxs,
		MessageInfos:      file_api_grpc_thumb_proto_msgTypes,
	}.Build()
	File_api_grpc_thumb_proto = out.File
	file_api_grpc_thumb_proto_rawDesc = nil
	file_api_grpc_thumb_proto_goTypes = nil
	file_api_grpc_thumb_proto_depIdxs = nil
}
