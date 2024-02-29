// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: role_api.proto

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

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{0}
}

func (x *RoleCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleCreateRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RoleCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleCreateResponse) Reset() {
	*x = RoleCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateResponse) ProtoMessage() {}

func (x *RoleCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateResponse.ProtoReflect.Descriptor instead.
func (*RoleCreateResponse) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{1}
}

func (x *RoleCreateResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *UUID    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *RoleUpdateRequest) Reset() {
	*x = RoleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateRequest) ProtoMessage() {}

func (x *RoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{2}
}

func (x *RoleUpdateRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RoleUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleUpdateRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RoleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleUpdateResponse) Reset() {
	*x = RoleUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateResponse) ProtoMessage() {}

func (x *RoleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateResponse.ProtoReflect.Descriptor instead.
func (*RoleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{3}
}

func (x *RoleUpdateResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleDeleteRequest) Reset() {
	*x = RoleDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteRequest) ProtoMessage() {}

func (x *RoleDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteRequest.ProtoReflect.Descriptor instead.
func (*RoleDeleteRequest) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{4}
}

func (x *RoleDeleteRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type RoleDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoleDeleteResponse) Reset() {
	*x = RoleDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDeleteResponse) ProtoMessage() {}

func (x *RoleDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDeleteResponse.ProtoReflect.Descriptor instead.
func (*RoleDeleteResponse) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{5}
}

type RoleFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleFindOneRequest) Reset() {
	*x = RoleFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindOneRequest) ProtoMessage() {}

func (x *RoleFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindOneRequest.ProtoReflect.Descriptor instead.
func (*RoleFindOneRequest) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{6}
}

func (x *RoleFindOneRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type RoleFindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleFindOneResponse) Reset() {
	*x = RoleFindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindOneResponse) ProtoMessage() {}

func (x *RoleFindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindOneResponse.ProtoReflect.Descriptor instead.
func (*RoleFindOneResponse) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{7}
}

func (x *RoleFindOneResponse) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type RoleFindManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *RoleFindManyResponse) Reset() {
	*x = RoleFindManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleFindManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleFindManyResponse) ProtoMessage() {}

func (x *RoleFindManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleFindManyResponse.ProtoReflect.Descriptor instead.
func (*RoleFindManyResponse) Descriptor() ([]byte, []int) {
	return file_role_api_proto_rawDescGZIP(), []int{8}
}

func (x *RoleFindManyResponse) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_role_api_proto protoreflect.FileDescriptor

var file_role_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x11, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x11,
	0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b,
	0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x52,
	0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x33, 0x0a,
	0x14, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_api_proto_rawDescOnce sync.Once
	file_role_api_proto_rawDescData = file_role_api_proto_rawDesc
)

func file_role_api_proto_rawDescGZIP() []byte {
	file_role_api_proto_rawDescOnce.Do(func() {
		file_role_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_api_proto_rawDescData)
	})
	return file_role_api_proto_rawDescData
}

var file_role_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_role_api_proto_goTypes = []interface{}{
	(*RoleCreateRequest)(nil),    // 0: RoleCreateRequest
	(*RoleCreateResponse)(nil),   // 1: RoleCreateResponse
	(*RoleUpdateRequest)(nil),    // 2: RoleUpdateRequest
	(*RoleUpdateResponse)(nil),   // 3: RoleUpdateResponse
	(*RoleDeleteRequest)(nil),    // 4: RoleDeleteRequest
	(*RoleDeleteResponse)(nil),   // 5: RoleDeleteResponse
	(*RoleFindOneRequest)(nil),   // 6: RoleFindOneRequest
	(*RoleFindOneResponse)(nil),  // 7: RoleFindOneResponse
	(*RoleFindManyResponse)(nil), // 8: RoleFindManyResponse
	(*Role)(nil),                 // 9: Role
	(*UUID)(nil),                 // 10: UUID
}
var file_role_api_proto_depIdxs = []int32{
	9,  // 0: RoleCreateResponse.role:type_name -> Role
	10, // 1: RoleUpdateRequest.id:type_name -> UUID
	9,  // 2: RoleUpdateResponse.role:type_name -> Role
	10, // 3: RoleDeleteRequest.id:type_name -> UUID
	10, // 4: RoleFindOneRequest.id:type_name -> UUID
	9,  // 5: RoleFindOneResponse.role:type_name -> Role
	9,  // 6: RoleFindManyResponse.roles:type_name -> Role
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_role_api_proto_init() }
func file_role_api_proto_init() {
	if File_role_api_proto != nil {
		return
	}
	file_common_proto_init()
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateRequest); i {
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
		file_role_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateResponse); i {
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
		file_role_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateRequest); i {
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
		file_role_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateResponse); i {
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
		file_role_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteRequest); i {
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
		file_role_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDeleteResponse); i {
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
		file_role_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindOneRequest); i {
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
		file_role_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindOneResponse); i {
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
		file_role_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleFindManyResponse); i {
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
			RawDescriptor: file_role_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_api_proto_goTypes,
		DependencyIndexes: file_role_api_proto_depIdxs,
		MessageInfos:      file_role_api_proto_msgTypes,
	}.Build()
	File_role_api_proto = out.File
	file_role_api_proto_rawDesc = nil
	file_role_api_proto_goTypes = nil
	file_role_api_proto_depIdxs = nil
}
