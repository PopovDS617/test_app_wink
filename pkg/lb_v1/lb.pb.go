// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: lb.proto

package lb_v1

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

type GetVideoURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video string `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoURLRequest) Reset() {
	*x = GetVideoURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoURLRequest) ProtoMessage() {}

func (x *GetVideoURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoURLRequest.ProtoReflect.Descriptor instead.
func (*GetVideoURLRequest) Descriptor() ([]byte, []int) {
	return file_lb_proto_rawDescGZIP(), []int{0}
}

func (x *GetVideoURLRequest) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

type GetVideoURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video string `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *GetVideoURLResponse) Reset() {
	*x = GetVideoURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoURLResponse) ProtoMessage() {}

func (x *GetVideoURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoURLResponse.ProtoReflect.Descriptor instead.
func (*GetVideoURLResponse) Descriptor() ([]byte, []int) {
	return file_lb_proto_rawDescGZIP(), []int{1}
}

func (x *GetVideoURLResponse) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

var File_lb_proto protoreflect.FileDescriptor

var file_lb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6c, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x62, 0x5f, 0x76,
	0x31, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x2b, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x32, 0x4d, 0x0a, 0x05, 0x4c, 0x42,
	0x5f, 0x76, 0x31, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x52, 0x4c, 0x12, 0x19, 0x2e, 0x6c, 0x62, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6c, 0x62, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x6b, 0x67,
	0x2f, 0x6c, 0x62, 0x5f, 0x76, 0x31, 0x3b, 0x6c, 0x62, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lb_proto_rawDescOnce sync.Once
	file_lb_proto_rawDescData = file_lb_proto_rawDesc
)

func file_lb_proto_rawDescGZIP() []byte {
	file_lb_proto_rawDescOnce.Do(func() {
		file_lb_proto_rawDescData = protoimpl.X.CompressGZIP(file_lb_proto_rawDescData)
	})
	return file_lb_proto_rawDescData
}

var file_lb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lb_proto_goTypes = []interface{}{
	(*GetVideoURLRequest)(nil),  // 0: lb_v1.GetVideoURLRequest
	(*GetVideoURLResponse)(nil), // 1: lb_v1.GetVideoURLResponse
}
var file_lb_proto_depIdxs = []int32{
	0, // 0: lb_v1.LB_v1.GetVideoURL:input_type -> lb_v1.GetVideoURLRequest
	1, // 1: lb_v1.LB_v1.GetVideoURL:output_type -> lb_v1.GetVideoURLResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lb_proto_init() }
func file_lb_proto_init() {
	if File_lb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoURLRequest); i {
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
		file_lb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoURLResponse); i {
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
			RawDescriptor: file_lb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lb_proto_goTypes,
		DependencyIndexes: file_lb_proto_depIdxs,
		MessageInfos:      file_lb_proto_msgTypes,
	}.Build()
	File_lb_proto = out.File
	file_lb_proto_rawDesc = nil
	file_lb_proto_goTypes = nil
	file_lb_proto_depIdxs = nil
}
