// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: protobuf/equation.proto

package protobuf

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

type BalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Equation      string                 `protobuf:"bytes,1,opt,name=equation,proto3" json:"equation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	mi := &file_protobuf_equation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_equation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_equation_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceRequest) GetEquation() string {
	if x != nil {
		return x.Equation
	}
	return ""
}

type BalanceResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	BalancedEquation string                 `protobuf:"bytes,1,opt,name=balanced_equation,json=balancedEquation,proto3" json:"balanced_equation,omitempty"`
	ErrorMessage     string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	mi := &file_protobuf_equation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_equation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_equation_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceResponse) GetBalancedEquation() string {
	if x != nil {
		return x.BalancedEquation
	}
	return ""
}

func (x *BalanceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_protobuf_equation_proto protoreflect.FileDescriptor

var file_protobuf_equation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x71, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x22, 0x2c, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x63, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5a, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x6d, 0x69, 0x63,
	0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_equation_proto_rawDescOnce sync.Once
	file_protobuf_equation_proto_rawDescData = file_protobuf_equation_proto_rawDesc
)

func file_protobuf_equation_proto_rawDescGZIP() []byte {
	file_protobuf_equation_proto_rawDescOnce.Do(func() {
		file_protobuf_equation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_equation_proto_rawDescData)
	})
	return file_protobuf_equation_proto_rawDescData
}

var file_protobuf_equation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_equation_proto_goTypes = []any{
	(*BalanceRequest)(nil),  // 0: protobuf.BalanceRequest
	(*BalanceResponse)(nil), // 1: protobuf.BalanceResponse
}
var file_protobuf_equation_proto_depIdxs = []int32{
	0, // 0: protobuf.ChemicalBalancer.BalanceEquation:input_type -> protobuf.BalanceRequest
	1, // 1: protobuf.ChemicalBalancer.BalanceEquation:output_type -> protobuf.BalanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_equation_proto_init() }
func file_protobuf_equation_proto_init() {
	if File_protobuf_equation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_equation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_equation_proto_goTypes,
		DependencyIndexes: file_protobuf_equation_proto_depIdxs,
		MessageInfos:      file_protobuf_equation_proto_msgTypes,
	}.Build()
	File_protobuf_equation_proto = out.File
	file_protobuf_equation_proto_rawDesc = nil
	file_protobuf_equation_proto_goTypes = nil
	file_protobuf_equation_proto_depIdxs = nil
}
