// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: ticker.proto

package gogen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TickerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbol        string                 `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TickerRequest) Reset() {
	*x = TickerRequest{}
	mi := &file_ticker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TickerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickerRequest) ProtoMessage() {}

func (x *TickerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickerRequest.ProtoReflect.Descriptor instead.
func (*TickerRequest) Descriptor() ([]byte, []int) {
	return file_ticker_proto_rawDescGZIP(), []int{0}
}

func (x *TickerRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type TickerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbol        string                 `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Ltp           float64                `protobuf:"fixed64,2,opt,name=ltp,proto3" json:"ltp,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TickerResponse) Reset() {
	*x = TickerResponse{}
	mi := &file_ticker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TickerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickerResponse) ProtoMessage() {}

func (x *TickerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickerResponse.ProtoReflect.Descriptor instead.
func (*TickerResponse) Descriptor() ([]byte, []int) {
	return file_ticker_proto_rawDescGZIP(), []int{1}
}

func (x *TickerResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TickerResponse) GetLtp() float64 {
	if x != nil {
		return x.Ltp
	}
	return 0
}

func (x *TickerResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_ticker_proto protoreflect.FileDescriptor

var file_ticker_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x67, 0x6f, 0x67, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x58,
	0x0a, 0x0e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x74, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x74, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x56, 0x0a, 0x13, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x14, 0x2e, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_ticker_proto_rawDescOnce sync.Once
	file_ticker_proto_rawDescData []byte
)

func file_ticker_proto_rawDescGZIP() []byte {
	file_ticker_proto_rawDescOnce.Do(func() {
		file_ticker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ticker_proto_rawDesc), len(file_ticker_proto_rawDesc)))
	})
	return file_ticker_proto_rawDescData
}

var file_ticker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ticker_proto_goTypes = []any{
	(*TickerRequest)(nil),  // 0: gogen.TickerRequest
	(*TickerResponse)(nil), // 1: gogen.TickerResponse
}
var file_ticker_proto_depIdxs = []int32{
	0, // 0: gogen.TickerStreamService.TickerStream:input_type -> gogen.TickerRequest
	1, // 1: gogen.TickerStreamService.TickerStream:output_type -> gogen.TickerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ticker_proto_init() }
func file_ticker_proto_init() {
	if File_ticker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ticker_proto_rawDesc), len(file_ticker_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticker_proto_goTypes,
		DependencyIndexes: file_ticker_proto_depIdxs,
		MessageInfos:      file_ticker_proto_msgTypes,
	}.Build()
	File_ticker_proto = out.File
	file_ticker_proto_goTypes = nil
	file_ticker_proto_depIdxs = nil
}
