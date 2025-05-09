// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: pipe.proto

package protos

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

// 검증·빌드 상태
type BuildStatus int32

const (
	BuildStatus_OK   BuildStatus = 0 // 문제 없음
	BuildStatus_WARN BuildStatus = 1 // 경고만 있음(진행 가능)
	BuildStatus_FAIL BuildStatus = 2 // 치명적 오류(반려)
)

// Enum value maps for BuildStatus.
var (
	BuildStatus_name = map[int32]string{
		0: "OK",
		1: "WARN",
		2: "FAIL",
	}
	BuildStatus_value = map[string]int32{
		"OK":   0,
		"WARN": 1,
		"FAIL": 2,
	}
)

func (x BuildStatus) Enum() *BuildStatus {
	p := new(BuildStatus)
	*p = x
	return p
}

func (x BuildStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pipe_proto_enumTypes[0].Descriptor()
}

func (BuildStatus) Type() protoreflect.EnumType {
	return &file_pipe_proto_enumTypes[0]
}

func (x BuildStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildStatus.Descriptor instead.
func (BuildStatus) EnumDescriptor() ([]byte, []int) {
	return file_pipe_proto_rawDescGZIP(), []int{0}
}

// 요청 메시지
type DockerfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dockerfile 전체 텍스트
	DockerfileContent string `protobuf:"bytes,1,opt,name=dockerfile_content,json=dockerfileContent,proto3" json:"dockerfile_content,omitempty"`
	// 이미 서버에 저장된 Dockerfile 을 다시 참조할 때 사용, 처음 보낼때는 값을 empty 로 보낸다.
	DockerfileId string `protobuf:"bytes,2,opt,name=dockerfile_id,json=dockerfileId,proto3" json:"dockerfile_id,omitempty"`
}

func (x *DockerfileRequest) Reset() {
	*x = DockerfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerfileRequest) ProtoMessage() {}

func (x *DockerfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerfileRequest.ProtoReflect.Descriptor instead.
func (*DockerfileRequest) Descriptor() ([]byte, []int) {
	return file_pipe_proto_rawDescGZIP(), []int{0}
}

func (x *DockerfileRequest) GetDockerfileContent() string {
	if x != nil {
		return x.DockerfileContent
	}
	return ""
}

func (x *DockerfileRequest) GetDockerfileId() string {
	if x != nil {
		return x.DockerfileId
	}
	return ""
}

// 응답 메시지
type DockerfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 처리된 Dockerfile 의 고유 ID (서버가 생성 또는 재사용)
	DockerfileId string `protobuf:"bytes,1,opt,name=dockerfile_id,json=dockerfileId,proto3" json:"dockerfile_id,omitempty"`
	// 검증 또는 빌드 결과 상태
	Status BuildStatus `protobuf:"varint,2,opt,name=status,proto3,enum=protos.BuildStatus" json:"status,omitempty"`
	// 위험 구문·경고·오류 메시지 목록
	Messages []string `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *DockerfileResponse) Reset() {
	*x = DockerfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerfileResponse) ProtoMessage() {}

func (x *DockerfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerfileResponse.ProtoReflect.Descriptor instead.
func (*DockerfileResponse) Descriptor() ([]byte, []int) {
	return file_pipe_proto_rawDescGZIP(), []int{1}
}

func (x *DockerfileResponse) GetDockerfileId() string {
	if x != nil {
		return x.DockerfileId
	}
	return ""
}

func (x *DockerfileResponse) GetStatus() BuildStatus {
	if x != nil {
		return x.Status
	}
	return BuildStatus_OK
}

func (x *DockerfileResponse) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_pipe_proto protoreflect.FileDescriptor

var file_pipe_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x67, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x82, 0x01,
	0x0a, 0x12, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2a, 0x29, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52,
	0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x32, 0x5f, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x67, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6f,
	0x79, 0x68, 0x61, 0x65, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipe_proto_rawDescOnce sync.Once
	file_pipe_proto_rawDescData = file_pipe_proto_rawDesc
)

func file_pipe_proto_rawDescGZIP() []byte {
	file_pipe_proto_rawDescOnce.Do(func() {
		file_pipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipe_proto_rawDescData)
	})
	return file_pipe_proto_rawDescData
}

var file_pipe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pipe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pipe_proto_goTypes = []interface{}{
	(BuildStatus)(0),           // 0: protos.BuildStatus
	(*DockerfileRequest)(nil),  // 1: protos.DockerfileRequest
	(*DockerfileResponse)(nil), // 2: protos.DockerfileResponse
}
var file_pipe_proto_depIdxs = []int32{
	0, // 0: protos.DockerfileResponse.status:type_name -> protos.BuildStatus
	1, // 1: protos.StageBuilderService.BuildStageImage:input_type -> protos.DockerfileRequest
	2, // 2: protos.StageBuilderService.BuildStageImage:output_type -> protos.DockerfileResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pipe_proto_init() }
func file_pipe_proto_init() {
	if File_pipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerfileRequest); i {
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
		file_pipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerfileResponse); i {
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
			RawDescriptor: file_pipe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pipe_proto_goTypes,
		DependencyIndexes: file_pipe_proto_depIdxs,
		EnumInfos:         file_pipe_proto_enumTypes,
		MessageInfos:      file_pipe_proto_msgTypes,
	}.Build()
	File_pipe_proto = out.File
	file_pipe_proto_rawDesc = nil
	file_pipe_proto_goTypes = nil
	file_pipe_proto_depIdxs = nil
}
