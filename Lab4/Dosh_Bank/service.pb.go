// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service.proto

package main

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

type GetTotalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTotalRequest) Reset() {
	*x = GetTotalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalRequest) ProtoMessage() {}

func (x *GetTotalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalRequest.ProtoReflect.Descriptor instead.
func (*GetTotalRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type GetTotalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetTotalResponse) Reset() {
	*x = GetTotalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalResponse) ProtoMessage() {}

func (x *GetTotalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalResponse.ProtoReflect.Descriptor instead.
func (*GetTotalResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTotalResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x47, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x3f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5e,
	0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x70,
	0x65, 0x62, 0x65, 0x61, 0x74, 0x73, 0x2d, 0x66, 0x6c, 0x70, 0x2f, 0x4c, 0x61, 0x62, 0x6f, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x73, 0x2d, 0x53, 0x69, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x73,
	0x2d, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x2f, 0x74, 0x72,
	0x65, 0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x61, 0x62, 0x34, 0x2f, 0x44,
	0x6f, 0x73, 0x68, 0x5f, 0x42, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_proto_goTypes = []interface{}{
	(*GetTotalRequest)(nil),  // 0: service.GetTotalRequest
	(*GetTotalResponse)(nil), // 1: service.GetTotalResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: service.Bank.GetTotal:input_type -> service.GetTotalRequest
	1, // 1: service.Bank.GetTotal:output_type -> service.GetTotalResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
