// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: info_api.proto

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

type HardwareInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HardwareInfoRequest) Reset() {
	*x = HardwareInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareInfoRequest) ProtoMessage() {}

func (x *HardwareInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareInfoRequest.ProtoReflect.Descriptor instead.
func (*HardwareInfoRequest) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{0}
}

type HardwareInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os            string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Cpu           string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	CpuClockSpeed uint32 `protobuf:"varint,3,opt,name=cpu_clock_speed,json=cpuClockSpeed,proto3" json:"cpu_clock_speed,omitempty"`
	Cores         uint32 `protobuf:"varint,4,opt,name=cores,proto3" json:"cores,omitempty"`
	LogicalCores  uint32 `protobuf:"varint,5,opt,name=logical_cores,json=logicalCores,proto3" json:"logical_cores,omitempty"`
	TotalMemory   uint64 `protobuf:"varint,6,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	TotalDisk     uint64 `protobuf:"varint,7,opt,name=total_disk,json=totalDisk,proto3" json:"total_disk,omitempty"`
	MemoryUsage   uint64 `protobuf:"varint,8,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	DiskUsage     uint64 `protobuf:"varint,9,opt,name=disk_usage,json=diskUsage,proto3" json:"disk_usage,omitempty"`
}

func (x *HardwareInfoResponse) Reset() {
	*x = HardwareInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareInfoResponse) ProtoMessage() {}

func (x *HardwareInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareInfoResponse.ProtoReflect.Descriptor instead.
func (*HardwareInfoResponse) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{1}
}

func (x *HardwareInfoResponse) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *HardwareInfoResponse) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *HardwareInfoResponse) GetCpuClockSpeed() uint32 {
	if x != nil {
		return x.CpuClockSpeed
	}
	return 0
}

func (x *HardwareInfoResponse) GetCores() uint32 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *HardwareInfoResponse) GetLogicalCores() uint32 {
	if x != nil {
		return x.LogicalCores
	}
	return 0
}

func (x *HardwareInfoResponse) GetTotalMemory() uint64 {
	if x != nil {
		return x.TotalMemory
	}
	return 0
}

func (x *HardwareInfoResponse) GetTotalDisk() uint64 {
	if x != nil {
		return x.TotalDisk
	}
	return 0
}

func (x *HardwareInfoResponse) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *HardwareInfoResponse) GetDiskUsage() uint64 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

type GetFreePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFreePortRequest) Reset() {
	*x = GetFreePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreePortRequest) ProtoMessage() {}

func (x *GetFreePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreePortRequest.ProtoReflect.Descriptor instead.
func (*GetFreePortRequest) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{2}
}

type GetFreePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FreePort *Port `protobuf:"bytes,1,opt,name=free_port,json=freePort,proto3" json:"free_port,omitempty"`
}

func (x *GetFreePortResponse) Reset() {
	*x = GetFreePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFreePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFreePortResponse) ProtoMessage() {}

func (x *GetFreePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFreePortResponse.ProtoReflect.Descriptor instead.
func (*GetFreePortResponse) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetFreePortResponse) GetFreePort() *Port {
	if x != nil {
		return x.FreePort
	}
	return nil
}

type CreateDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CreateDirectoryRequest) Reset() {
	*x = CreateDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryRequest) ProtoMessage() {}

func (x *CreateDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryRequest.ProtoReflect.Descriptor instead.
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDirectoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CreateDirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDirectoryResponse) Reset() {
	*x = CreateDirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryResponse) ProtoMessage() {}

func (x *CreateDirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_info_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryResponse.ProtoReflect.Descriptor instead.
func (*CreateDirectoryResponse) Descriptor() ([]byte, []int) {
	return file_info_api_proto_rawDescGZIP(), []int{5}
}

var File_info_api_proto protoreflect.FileDescriptor

var file_info_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x14, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x70, 0x75, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x69, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x73,
	0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x08, 0x66,
	0x72, 0x65, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xce, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x48, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_api_proto_rawDescOnce sync.Once
	file_info_api_proto_rawDescData = file_info_api_proto_rawDesc
)

func file_info_api_proto_rawDescGZIP() []byte {
	file_info_api_proto_rawDescOnce.Do(func() {
		file_info_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_api_proto_rawDescData)
	})
	return file_info_api_proto_rawDescData
}

var file_info_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_info_api_proto_goTypes = []interface{}{
	(*HardwareInfoRequest)(nil),     // 0: HardwareInfoRequest
	(*HardwareInfoResponse)(nil),    // 1: HardwareInfoResponse
	(*GetFreePortRequest)(nil),      // 2: GetFreePortRequest
	(*GetFreePortResponse)(nil),     // 3: GetFreePortResponse
	(*CreateDirectoryRequest)(nil),  // 4: CreateDirectoryRequest
	(*CreateDirectoryResponse)(nil), // 5: CreateDirectoryResponse
	(*Port)(nil),                    // 6: Port
}
var file_info_api_proto_depIdxs = []int32{
	6, // 0: GetFreePortResponse.free_port:type_name -> Port
	0, // 1: NodeInfo.GetNodeHardwareInfo:input_type -> HardwareInfoRequest
	2, // 2: NodeInfo.GetFreePort:input_type -> GetFreePortRequest
	4, // 3: NodeInfo.CreateDirectory:input_type -> CreateDirectoryRequest
	1, // 4: NodeInfo.GetNodeHardwareInfo:output_type -> HardwareInfoResponse
	3, // 5: NodeInfo.GetFreePort:output_type -> GetFreePortResponse
	5, // 6: NodeInfo.CreateDirectory:output_type -> CreateDirectoryResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_info_api_proto_init() }
func file_info_api_proto_init() {
	if File_info_api_proto != nil {
		return
	}
	file_common_proto_init()
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_info_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareInfoRequest); i {
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
		file_info_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareInfoResponse); i {
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
		file_info_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreePortRequest); i {
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
		file_info_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFreePortResponse); i {
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
		file_info_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryRequest); i {
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
		file_info_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryResponse); i {
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
			RawDescriptor: file_info_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_api_proto_goTypes,
		DependencyIndexes: file_info_api_proto_depIdxs,
		MessageInfos:      file_info_api_proto_msgTypes,
	}.Build()
	File_info_api_proto = out.File
	file_info_api_proto_rawDesc = nil
	file_info_api_proto_goTypes = nil
	file_info_api_proto_depIdxs = nil
}
