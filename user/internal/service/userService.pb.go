// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: userService.proto

package service

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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDetail *UserModel `protobuf:"bytes,1,opt,name=UserDetail,proto3" json:"UserDetail,omitempty"`
	Code       uint32     `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *UserDetailResponse) Reset() {
	*x = UserDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailResponse) ProtoMessage() {}

func (x *UserDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailResponse.ProtoReflect.Descriptor instead.
func (*UserDetailResponse) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetailResponse) GetUserDetail() *UserModel {
	if x != nil {
		return x.UserDetail
	}
	return nil
}

func (x *UserDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_userService_proto protoreflect.FileDescriptor

var file_userService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x60, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xb2, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_userService_proto_rawDescOnce sync.Once
	file_userService_proto_rawDescData = file_userService_proto_rawDesc
)

func file_userService_proto_rawDescGZIP() []byte {
	file_userService_proto_rawDescOnce.Do(func() {
		file_userService_proto_rawDescData = protoimpl.X.CompressGZIP(file_userService_proto_rawDescData)
	})
	return file_userService_proto_rawDescData
}

var file_userService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userService_proto_goTypes = []interface{}{
	(*UserRequest)(nil),        // 0: userService.UserRequest
	(*UserDetailResponse)(nil), // 1: userService.UserDetailResponse
	(*UserModel)(nil),          // 2: userService.UserModel
}
var file_userService_proto_depIdxs = []int32{
	2, // 0: userService.UserDetailResponse.UserDetail:type_name -> userService.UserModel
	0, // 1: userService.UserService.UserLogin:input_type -> userService.UserRequest
	0, // 2: userService.UserService.UserRegister:input_type -> userService.UserRequest
	0, // 3: userService.UserService.UserLogout:input_type -> userService.UserRequest
	0, // 4: userService.UserService.UserDelete:input_type -> userService.UserRequest
	1, // 5: userService.UserService.UserLogin:output_type -> userService.UserDetailResponse
	1, // 6: userService.UserService.UserRegister:output_type -> userService.UserDetailResponse
	1, // 7: userService.UserService.UserLogout:output_type -> userService.UserDetailResponse
	1, // 8: userService.UserService.UserDelete:output_type -> userService.UserDetailResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userService_proto_init() }
func file_userService_proto_init() {
	if File_userService_proto != nil {
		return
	}
	file_userModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_userService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_userService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailResponse); i {
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
			RawDescriptor: file_userService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userService_proto_goTypes,
		DependencyIndexes: file_userService_proto_depIdxs,
		MessageInfos:      file_userService_proto_msgTypes,
	}.Build()
	File_userService_proto = out.File
	file_userService_proto_rawDesc = nil
	file_userService_proto_goTypes = nil
	file_userService_proto_depIdxs = nil
}
