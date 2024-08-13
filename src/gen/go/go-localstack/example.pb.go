// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: example.proto

package go_localstack

import (
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

// リクエストメッセージを定義します
type GetExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // IDを指定するフィールド
}

func (x *GetExampleRequest) Reset() {
	*x = GetExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleRequest) ProtoMessage() {}

func (x *GetExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleRequest.ProtoReflect.Descriptor instead.
func (*GetExampleRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *GetExampleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetExampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // メッセージを含むフィールド
}

func (x *GetExampleResponse) Reset() {
	*x = GetExampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleResponse) ProtoMessage() {}

func (x *GetExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleResponse.ProtoReflect.Descriptor instead.
func (*GetExampleResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *GetExampleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Insert用のリクエストメッセージを定義します
type InsertExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 挿入するメッセージを指定するフィールド
}

func (x *InsertExampleRequest) Reset() {
	*x = InsertExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertExampleRequest) ProtoMessage() {}

func (x *InsertExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertExampleRequest.ProtoReflect.Descriptor instead.
func (*InsertExampleRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *InsertExampleRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Insert用のレスポンスメッセージを定義します
type InsertExampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 挿入が成功したかどうかを示すフィールド
}

func (x *InsertExampleResponse) Reset() {
	*x = InsertExampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertExampleResponse) ProtoMessage() {}

func (x *InsertExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertExampleResponse.ProtoReflect.Descriptor instead.
func (*InsertExampleResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{3}
}

func (x *InsertExampleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a,
	0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xd9, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x10, 0x5a, 0x0e,
	0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_proto_goTypes = []any{
	(*GetExampleRequest)(nil),     // 0: example.GetExampleRequest
	(*GetExampleResponse)(nil),    // 1: example.GetExampleResponse
	(*InsertExampleRequest)(nil),  // 2: example.InsertExampleRequest
	(*InsertExampleResponse)(nil), // 3: example.InsertExampleResponse
}
var file_example_proto_depIdxs = []int32{
	0, // 0: example.ExampleService.GetExample:input_type -> example.GetExampleRequest
	2, // 1: example.ExampleService.InsertExample:input_type -> example.InsertExampleRequest
	1, // 2: example.ExampleService.GetExample:output_type -> example.GetExampleResponse
	3, // 3: example.ExampleService.InsertExample:output_type -> example.InsertExampleResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetExampleRequest); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetExampleResponse); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InsertExampleRequest); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*InsertExampleResponse); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
