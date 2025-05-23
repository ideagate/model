// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: client/controller/v1/command.proto

package v1

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

type GetListPodRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListPodRequest) Reset() {
	*x = GetListPodRequest{}
	mi := &file_client_controller_v1_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListPodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPodRequest) ProtoMessage() {}

func (x *GetListPodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_v1_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPodRequest.ProtoReflect.Descriptor instead.
func (*GetListPodRequest) Descriptor() ([]byte, []int) {
	return file_client_controller_v1_command_proto_rawDescGZIP(), []int{0}
}

type GetListPodResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pods          []*Pod                 `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListPodResponse) Reset() {
	*x = GetListPodResponse{}
	mi := &file_client_controller_v1_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListPodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPodResponse) ProtoMessage() {}

func (x *GetListPodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_v1_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPodResponse.ProtoReflect.Descriptor instead.
func (*GetListPodResponse) Descriptor() ([]byte, []int) {
	return file_client_controller_v1_command_proto_rawDescGZIP(), []int{1}
}

func (x *GetListPodResponse) GetPods() []*Pod {
	if x != nil {
		return x.Pods
	}
	return nil
}

var File_client_controller_v1_command_proto protoreflect.FileDescriptor

const file_client_controller_v1_command_proto_rawDesc = "" +
	"\n" +
	"\"client/controller/v1/command.proto\x12\x02v1\x1a\x1eclient/controller/v1/pod.proto\"\x13\n" +
	"\x11GetListPodRequest\"1\n" +
	"\x12GetListPodResponse\x12\x1b\n" +
	"\x04pods\x18\x01 \x03(\v2\a.v1.PodR\x04podsBu\n" +
	"\x06com.v1B\fCommandProtoP\x01Z5github.com/ideagate/model/gen-go/client/controller/v1\xa2\x02\x03VXX\xaa\x02\x02V1\xca\x02\x02V1\xe2\x02\x0eV1\\GPBMetadata\xea\x02\x02V1b\x06proto3"

var (
	file_client_controller_v1_command_proto_rawDescOnce sync.Once
	file_client_controller_v1_command_proto_rawDescData []byte
)

func file_client_controller_v1_command_proto_rawDescGZIP() []byte {
	file_client_controller_v1_command_proto_rawDescOnce.Do(func() {
		file_client_controller_v1_command_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_client_controller_v1_command_proto_rawDesc), len(file_client_controller_v1_command_proto_rawDesc)))
	})
	return file_client_controller_v1_command_proto_rawDescData
}

var file_client_controller_v1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_controller_v1_command_proto_goTypes = []any{
	(*GetListPodRequest)(nil),  // 0: v1.GetListPodRequest
	(*GetListPodResponse)(nil), // 1: v1.GetListPodResponse
	(*Pod)(nil),                // 2: v1.Pod
}
var file_client_controller_v1_command_proto_depIdxs = []int32{
	2, // 0: v1.GetListPodResponse.pods:type_name -> v1.Pod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_controller_v1_command_proto_init() }
func file_client_controller_v1_command_proto_init() {
	if File_client_controller_v1_command_proto != nil {
		return
	}
	file_client_controller_v1_pod_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_client_controller_v1_command_proto_rawDesc), len(file_client_controller_v1_command_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_controller_v1_command_proto_goTypes,
		DependencyIndexes: file_client_controller_v1_command_proto_depIdxs,
		MessageInfos:      file_client_controller_v1_command_proto_msgTypes,
	}.Build()
	File_client_controller_v1_command_proto = out.File
	file_client_controller_v1_command_proto_goTypes = nil
	file_client_controller_v1_command_proto_depIdxs = nil
}
