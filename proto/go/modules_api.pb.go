// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: modules_api.proto

package protoapi

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

type FindAllModulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontendOnly bool `protobuf:"varint,1,opt,name=frontend_only,json=frontendOnly,proto3" json:"frontend_only,omitempty"`
	BackendOnly  bool `protobuf:"varint,2,opt,name=backend_only,json=backendOnly,proto3" json:"backend_only,omitempty"`
}

func (x *FindAllModulesRequest) Reset() {
	*x = FindAllModulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllModulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllModulesRequest) ProtoMessage() {}

func (x *FindAllModulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllModulesRequest.ProtoReflect.Descriptor instead.
func (*FindAllModulesRequest) Descriptor() ([]byte, []int) {
	return file_modules_api_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllModulesRequest) GetFrontendOnly() bool {
	if x != nil {
		return x.FrontendOnly
	}
	return false
}

func (x *FindAllModulesRequest) GetBackendOnly() bool {
	if x != nil {
		return x.BackendOnly
	}
	return false
}

type FindAllModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *FindAllModulesResponse) Reset() {
	*x = FindAllModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllModulesResponse) ProtoMessage() {}

func (x *FindAllModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllModulesResponse.ProtoReflect.Descriptor instead.
func (*FindAllModulesResponse) Descriptor() ([]byte, []int) {
	return file_modules_api_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllModulesResponse) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

var File_modules_api_proto protoreflect.FileDescriptor

var file_modules_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x22, 0x3b, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_api_proto_rawDescOnce sync.Once
	file_modules_api_proto_rawDescData = file_modules_api_proto_rawDesc
)

func file_modules_api_proto_rawDescGZIP() []byte {
	file_modules_api_proto_rawDescOnce.Do(func() {
		file_modules_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_api_proto_rawDescData)
	})
	return file_modules_api_proto_rawDescData
}

var file_modules_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_modules_api_proto_goTypes = []interface{}{
	(*FindAllModulesRequest)(nil),  // 0: FindAllModulesRequest
	(*FindAllModulesResponse)(nil), // 1: FindAllModulesResponse
	(*Module)(nil),                 // 2: Module
}
var file_modules_api_proto_depIdxs = []int32{
	2, // 0: FindAllModulesResponse.modules:type_name -> Module
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modules_api_proto_init() }
func file_modules_api_proto_init() {
	if File_modules_api_proto != nil {
		return
	}
	file_generic_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_modules_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllModulesRequest); i {
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
		file_modules_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllModulesResponse); i {
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
			RawDescriptor: file_modules_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_api_proto_goTypes,
		DependencyIndexes: file_modules_api_proto_depIdxs,
		MessageInfos:      file_modules_api_proto_msgTypes,
	}.Build()
	File_modules_api_proto = out.File
	file_modules_api_proto_rawDesc = nil
	file_modules_api_proto_goTypes = nil
	file_modules_api_proto_depIdxs = nil
}
