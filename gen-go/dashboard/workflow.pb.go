// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: dashboard/workflow.proto

package dashboard

import (
	endpoint "github.com/ideagate/model/gen-go/core/endpoint"
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

type CreateWorkflowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ApplicationId string                 `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	EntrypointId  string                 `protobuf:"bytes,3,opt,name=entrypoint_id,json=entrypointId,proto3" json:"entrypoint_id,omitempty"`
	// New workflow will be copied from this version.
	// If not specified, then new workflow will be created from scratch.
	FromVersion   int64 `protobuf:"varint,4,opt,name=from_version,json=fromVersion,proto3" json:"from_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWorkflowRequest) Reset() {
	*x = CreateWorkflowRequest{}
	mi := &file_dashboard_workflow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkflowRequest) ProtoMessage() {}

func (x *CreateWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkflowRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateWorkflowRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *CreateWorkflowRequest) GetEntrypointId() string {
	if x != nil {
		return x.EntrypointId
	}
	return ""
}

func (x *CreateWorkflowRequest) GetFromVersion() int64 {
	if x != nil {
		return x.FromVersion
	}
	return 0
}

type CreateWorkflowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       int64                  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWorkflowResponse) Reset() {
	*x = CreateWorkflowResponse{}
	mi := &file_dashboard_workflow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkflowResponse) ProtoMessage() {}

func (x *CreateWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkflowResponse.ProtoReflect.Descriptor instead.
func (*CreateWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkflowResponse) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type GetWorkflowsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ApplicationId string                 `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	EntrypointId  string                 `protobuf:"bytes,3,opt,name=entrypoint_id,json=entrypointId,proto3" json:"entrypoint_id,omitempty"`
	Version       int64                  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkflowsRequest) Reset() {
	*x = GetWorkflowsRequest{}
	mi := &file_dashboard_workflow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowsRequest) ProtoMessage() {}

func (x *GetWorkflowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowsRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowsRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *GetWorkflowsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetWorkflowsRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *GetWorkflowsRequest) GetEntrypointId() string {
	if x != nil {
		return x.EntrypointId
	}
	return ""
}

func (x *GetWorkflowsRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type GetWorkflowsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Workflows     []*endpoint.Workflow   `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkflowsResponse) Reset() {
	*x = GetWorkflowsResponse{}
	mi := &file_dashboard_workflow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowsResponse) ProtoMessage() {}

func (x *GetWorkflowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowsResponse.ProtoReflect.Descriptor instead.
func (*GetWorkflowsResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorkflowsResponse) GetWorkflows() []*endpoint.Workflow {
	if x != nil {
		return x.Workflows
	}
	return nil
}

type UpdateWorkflowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Workflow      *endpoint.Workflow     `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWorkflowRequest) Reset() {
	*x = UpdateWorkflowRequest{}
	mi := &file_dashboard_workflow_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowRequest) ProtoMessage() {}

func (x *UpdateWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateWorkflowRequest) GetWorkflow() *endpoint.Workflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type UpdateWorkflowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWorkflowResponse) Reset() {
	*x = UpdateWorkflowResponse{}
	mi := &file_dashboard_workflow_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkflowResponse) ProtoMessage() {}

func (x *UpdateWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkflowResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{5}
}

type DeleteWorkflowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ApplicationId string                 `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	EntrypointId  string                 `protobuf:"bytes,3,opt,name=entrypoint_id,json=entrypointId,proto3" json:"entrypoint_id,omitempty"`
	Version       int64                  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWorkflowRequest) Reset() {
	*x = DeleteWorkflowRequest{}
	mi := &file_dashboard_workflow_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkflowRequest) ProtoMessage() {}

func (x *DeleteWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkflowRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteWorkflowRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *DeleteWorkflowRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *DeleteWorkflowRequest) GetEntrypointId() string {
	if x != nil {
		return x.EntrypointId
	}
	return ""
}

func (x *DeleteWorkflowRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DeleteWorkflowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWorkflowResponse) Reset() {
	*x = DeleteWorkflowResponse{}
	mi := &file_dashboard_workflow_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkflowResponse) ProtoMessage() {}

func (x *DeleteWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_workflow_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkflowResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_workflow_proto_rawDescGZIP(), []int{7}
}

var File_dashboard_workflow_proto protoreflect.FileDescriptor

const file_dashboard_workflow_proto_rawDesc = "" +
	"\n" +
	"\x18dashboard/workflow.proto\x12\tdashboard\x1a\x1ccore/endpoint/workflow.proto\"\xa5\x01\n" +
	"\x15CreateWorkflowRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x12%\n" +
	"\x0eapplication_id\x18\x02 \x01(\tR\rapplicationId\x12#\n" +
	"\rentrypoint_id\x18\x03 \x01(\tR\fentrypointId\x12!\n" +
	"\ffrom_version\x18\x04 \x01(\x03R\vfromVersion\"2\n" +
	"\x16CreateWorkflowResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\x03R\aversion\"\x9a\x01\n" +
	"\x13GetWorkflowsRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x12%\n" +
	"\x0eapplication_id\x18\x02 \x01(\tR\rapplicationId\x12#\n" +
	"\rentrypoint_id\x18\x03 \x01(\tR\fentrypointId\x12\x18\n" +
	"\aversion\x18\x04 \x01(\x03R\aversion\"H\n" +
	"\x14GetWorkflowsResponse\x120\n" +
	"\tworkflows\x18\x01 \x03(\v2\x12.endpoint.WorkflowR\tworkflows\"G\n" +
	"\x15UpdateWorkflowRequest\x12.\n" +
	"\bworkflow\x18\x01 \x01(\v2\x12.endpoint.WorkflowR\bworkflow\"\x18\n" +
	"\x16UpdateWorkflowResponse\"\x9c\x01\n" +
	"\x15DeleteWorkflowRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x12%\n" +
	"\x0eapplication_id\x18\x02 \x01(\tR\rapplicationId\x12#\n" +
	"\rentrypoint_id\x18\x03 \x01(\tR\fentrypointId\x12\x18\n" +
	"\aversion\x18\x04 \x01(\x03R\aversion\"\x18\n" +
	"\x16DeleteWorkflowResponseB\x8e\x01\n" +
	"\rcom.dashboardB\rWorkflowProtoP\x01Z*github.com/ideagate/model/gen-go/dashboard\xa2\x02\x03DXX\xaa\x02\tDashboard\xca\x02\tDashboard\xe2\x02\x15Dashboard\\GPBMetadata\xea\x02\tDashboardb\x06proto3"

var (
	file_dashboard_workflow_proto_rawDescOnce sync.Once
	file_dashboard_workflow_proto_rawDescData []byte
)

func file_dashboard_workflow_proto_rawDescGZIP() []byte {
	file_dashboard_workflow_proto_rawDescOnce.Do(func() {
		file_dashboard_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dashboard_workflow_proto_rawDesc), len(file_dashboard_workflow_proto_rawDesc)))
	})
	return file_dashboard_workflow_proto_rawDescData
}

var file_dashboard_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dashboard_workflow_proto_goTypes = []any{
	(*CreateWorkflowRequest)(nil),  // 0: dashboard.CreateWorkflowRequest
	(*CreateWorkflowResponse)(nil), // 1: dashboard.CreateWorkflowResponse
	(*GetWorkflowsRequest)(nil),    // 2: dashboard.GetWorkflowsRequest
	(*GetWorkflowsResponse)(nil),   // 3: dashboard.GetWorkflowsResponse
	(*UpdateWorkflowRequest)(nil),  // 4: dashboard.UpdateWorkflowRequest
	(*UpdateWorkflowResponse)(nil), // 5: dashboard.UpdateWorkflowResponse
	(*DeleteWorkflowRequest)(nil),  // 6: dashboard.DeleteWorkflowRequest
	(*DeleteWorkflowResponse)(nil), // 7: dashboard.DeleteWorkflowResponse
	(*endpoint.Workflow)(nil),      // 8: endpoint.Workflow
}
var file_dashboard_workflow_proto_depIdxs = []int32{
	8, // 0: dashboard.GetWorkflowsResponse.workflows:type_name -> endpoint.Workflow
	8, // 1: dashboard.UpdateWorkflowRequest.workflow:type_name -> endpoint.Workflow
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dashboard_workflow_proto_init() }
func file_dashboard_workflow_proto_init() {
	if File_dashboard_workflow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dashboard_workflow_proto_rawDesc), len(file_dashboard_workflow_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dashboard_workflow_proto_goTypes,
		DependencyIndexes: file_dashboard_workflow_proto_depIdxs,
		MessageInfos:      file_dashboard_workflow_proto_msgTypes,
	}.Build()
	File_dashboard_workflow_proto = out.File
	file_dashboard_workflow_proto_goTypes = nil
	file_dashboard_workflow_proto_depIdxs = nil
}
