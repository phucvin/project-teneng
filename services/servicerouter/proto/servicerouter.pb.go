// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: servicerouter.proto

package servicerouter

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

// The request message containing the user's name.
type InvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicerouter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicerouter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeRequest.ProtoReflect.Descriptor instead.
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return file_servicerouter_proto_rawDescGZIP(), []int{0}
}

func (x *InvokeRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// The response message containing the greetings
type InvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InvokeResponse) Reset() {
	*x = InvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicerouter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResponse) ProtoMessage() {}

func (x *InvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servicerouter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeResponse.ProtoReflect.Descriptor instead.
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return file_servicerouter_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_servicerouter_proto protoreflect.FileDescriptor

var file_servicerouter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x58, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x06,
	0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x75, 0x63, 0x76, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x74, 0x65, 0x6e, 0x65, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicerouter_proto_rawDescOnce sync.Once
	file_servicerouter_proto_rawDescData = file_servicerouter_proto_rawDesc
)

func file_servicerouter_proto_rawDescGZIP() []byte {
	file_servicerouter_proto_rawDescOnce.Do(func() {
		file_servicerouter_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicerouter_proto_rawDescData)
	})
	return file_servicerouter_proto_rawDescData
}

var file_servicerouter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_servicerouter_proto_goTypes = []interface{}{
	(*InvokeRequest)(nil),  // 0: servicerouter.InvokeRequest
	(*InvokeResponse)(nil), // 1: servicerouter.InvokeResponse
}
var file_servicerouter_proto_depIdxs = []int32{
	0, // 0: servicerouter.ServiceRouter.Invoke:input_type -> servicerouter.InvokeRequest
	1, // 1: servicerouter.ServiceRouter.Invoke:output_type -> servicerouter.InvokeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_servicerouter_proto_init() }
func file_servicerouter_proto_init() {
	if File_servicerouter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servicerouter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeRequest); i {
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
		file_servicerouter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeResponse); i {
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
			RawDescriptor: file_servicerouter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servicerouter_proto_goTypes,
		DependencyIndexes: file_servicerouter_proto_depIdxs,
		MessageInfos:      file_servicerouter_proto_msgTypes,
	}.Build()
	File_servicerouter_proto = out.File
	file_servicerouter_proto_rawDesc = nil
	file_servicerouter_proto_goTypes = nil
	file_servicerouter_proto_depIdxs = nil
}
