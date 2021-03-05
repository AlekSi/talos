// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: time/time.proto

package time

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/talos-systems/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The response message containing the ntp server
type TimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *TimeRequest) Reset() {
	*x = TimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRequest) ProtoMessage() {}

func (x *TimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRequest.ProtoReflect.Descriptor instead.
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata   *common.Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Server     string                 `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Localtime  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=localtime,proto3" json:"localtime,omitempty"`
	Remotetime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=remotetime,proto3" json:"remotetime,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{1}
}

func (x *Time) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Time) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *Time) GetLocaltime() *timestamppb.Timestamp {
	if x != nil {
		return x.Localtime
	}
	return nil
}

func (x *Time) GetRemotetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Remotetime
	}
	return nil
}

// The response message containing the ntp server, time, and offset
type TimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Time `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *TimeResponse) Reset() {
	*x = TimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeResponse) ProtoMessage() {}

func (x *TimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_time_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeResponse.ProtoReflect.Descriptor instead.
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return file_time_time_proto_rawDescGZIP(), []int{2}
}

func (x *TimeResponse) GetMessages() []*Time {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_time_time_proto protoreflect.FileDescriptor

var file_time_time_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0xc2, 0x01, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x38, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x75,
	0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x11,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x70, 0x69, 0x50, 0x01,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_time_proto_rawDescOnce sync.Once
	file_time_time_proto_rawDescData = file_time_time_proto_rawDesc
)

func file_time_time_proto_rawDescGZIP() []byte {
	file_time_time_proto_rawDescOnce.Do(func() {
		file_time_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_time_proto_rawDescData)
	})
	return file_time_time_proto_rawDescData
}

var (
	file_time_time_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_time_time_proto_goTypes  = []interface{}{
		(*TimeRequest)(nil),           // 0: time.TimeRequest
		(*Time)(nil),                  // 1: time.Time
		(*TimeResponse)(nil),          // 2: time.TimeResponse
		(*common.Metadata)(nil),       // 3: common.Metadata
		(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
		(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
	}
)

var file_time_time_proto_depIdxs = []int32{
	3, // 0: time.Time.metadata:type_name -> common.Metadata
	4, // 1: time.Time.localtime:type_name -> google.protobuf.Timestamp
	4, // 2: time.Time.remotetime:type_name -> google.protobuf.Timestamp
	1, // 3: time.TimeResponse.messages:type_name -> time.Time
	5, // 4: time.TimeService.Time:input_type -> google.protobuf.Empty
	0, // 5: time.TimeService.TimeCheck:input_type -> time.TimeRequest
	2, // 6: time.TimeService.Time:output_type -> time.TimeResponse
	2, // 7: time.TimeService.TimeCheck:output_type -> time.TimeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_time_time_proto_init() }
func file_time_time_proto_init() {
	if File_time_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRequest); i {
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
		file_time_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_time_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeResponse); i {
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
			RawDescriptor: file_time_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_time_proto_goTypes,
		DependencyIndexes: file_time_time_proto_depIdxs,
		MessageInfos:      file_time_time_proto_msgTypes,
	}.Build()
	File_time_time_proto = out.File
	file_time_time_proto_rawDesc = nil
	file_time_time_proto_goTypes = nil
	file_time_time_proto_depIdxs = nil
}
