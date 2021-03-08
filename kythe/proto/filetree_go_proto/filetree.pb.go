// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: kythe/proto/filetree.proto

package filetree_go_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DirectoryReply_Kind int32

const (
	DirectoryReply_UNKNOWN   DirectoryReply_Kind = 0
	DirectoryReply_FILE      DirectoryReply_Kind = 1
	DirectoryReply_DIRECTORY DirectoryReply_Kind = 2
)

// Enum value maps for DirectoryReply_Kind.
var (
	DirectoryReply_Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "FILE",
		2: "DIRECTORY",
	}
	DirectoryReply_Kind_value = map[string]int32{
		"UNKNOWN":   0,
		"FILE":      1,
		"DIRECTORY": 2,
	}
)

func (x DirectoryReply_Kind) Enum() *DirectoryReply_Kind {
	p := new(DirectoryReply_Kind)
	*p = x
	return p
}

func (x DirectoryReply_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DirectoryReply_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_kythe_proto_filetree_proto_enumTypes[0].Descriptor()
}

func (DirectoryReply_Kind) Type() protoreflect.EnumType {
	return &file_kythe_proto_filetree_proto_enumTypes[0]
}

func (x DirectoryReply_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DirectoryReply_Kind.Descriptor instead.
func (DirectoryReply_Kind) EnumDescriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{3, 0}
}

type CorpusRootsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CorpusRootsRequest) Reset() {
	*x = CorpusRootsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorpusRootsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorpusRootsRequest) ProtoMessage() {}

func (x *CorpusRootsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorpusRootsRequest.ProtoReflect.Descriptor instead.
func (*CorpusRootsRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{0}
}

type CorpusRootsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Corpus []*CorpusRootsReply_Corpus `protobuf:"bytes,1,rep,name=corpus,proto3" json:"corpus,omitempty"`
}

func (x *CorpusRootsReply) Reset() {
	*x = CorpusRootsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorpusRootsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorpusRootsReply) ProtoMessage() {}

func (x *CorpusRootsReply) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorpusRootsReply.ProtoReflect.Descriptor instead.
func (*CorpusRootsReply) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{1}
}

func (x *CorpusRootsReply) GetCorpus() []*CorpusRootsReply_Corpus {
	if x != nil {
		return x.Corpus
	}
	return nil
}

type DirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Corpus string `protobuf:"bytes,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Root   string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DirectoryRequest) Reset() {
	*x = DirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryRequest) ProtoMessage() {}

func (x *DirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryRequest.ProtoReflect.Descriptor instead.
func (*DirectoryRequest) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryRequest) GetCorpus() string {
	if x != nil {
		return x.Corpus
	}
	return ""
}

func (x *DirectoryRequest) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *DirectoryRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirectoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Corpus string                  `protobuf:"bytes,3,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Root   string                  `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	Path   string                  `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Entry  []*DirectoryReply_Entry `protobuf:"bytes,6,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (x *DirectoryReply) Reset() {
	*x = DirectoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryReply) ProtoMessage() {}

func (x *DirectoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryReply.ProtoReflect.Descriptor instead.
func (*DirectoryReply) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{3}
}

func (x *DirectoryReply) GetCorpus() string {
	if x != nil {
		return x.Corpus
	}
	return ""
}

func (x *DirectoryReply) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *DirectoryReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DirectoryReply) GetEntry() []*DirectoryReply_Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type CorpusRootsReply_Corpus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Root        []string `protobuf:"bytes,2,rep,name=root,proto3" json:"root,omitempty"`
	BuildConfig []string `protobuf:"bytes,3,rep,name=build_config,json=buildConfig,proto3" json:"build_config,omitempty"`
}

func (x *CorpusRootsReply_Corpus) Reset() {
	*x = CorpusRootsReply_Corpus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorpusRootsReply_Corpus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorpusRootsReply_Corpus) ProtoMessage() {}

func (x *CorpusRootsReply_Corpus) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorpusRootsReply_Corpus.ProtoReflect.Descriptor instead.
func (*CorpusRootsReply_Corpus) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CorpusRootsReply_Corpus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CorpusRootsReply_Corpus) GetRoot() []string {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *CorpusRootsReply_Corpus) GetBuildConfig() []string {
	if x != nil {
		return x.BuildConfig
	}
	return nil
}

type DirectoryReply_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        DirectoryReply_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=kythe.proto.DirectoryReply_Kind" json:"kind,omitempty"`
	Name        string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BuildConfig []string            `protobuf:"bytes,3,rep,name=build_config,json=buildConfig,proto3" json:"build_config,omitempty"`
}

func (x *DirectoryReply_Entry) Reset() {
	*x = DirectoryReply_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_filetree_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryReply_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryReply_Entry) ProtoMessage() {}

func (x *DirectoryReply_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_filetree_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryReply_Entry.ProtoReflect.Descriptor instead.
func (*DirectoryReply_Entry) Descriptor() ([]byte, []int) {
	return file_kythe_proto_filetree_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DirectoryReply_Entry) GetKind() DirectoryReply_Kind {
	if x != nil {
		return x.Kind
	}
	return DirectoryReply_UNKNOWN
}

func (x *DirectoryReply_Entry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirectoryReply_Entry) GetBuildConfig() []string {
	if x != nil {
		return x.BuildConfig
	}
	return nil
}

var File_kythe_proto_filetree_proto protoreflect.FileDescriptor

var file_kythe_proto_filetree_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xa5, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x1a, 0x53, 0x0a, 0x06, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x52, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xb9, 0x02, 0x0a, 0x0e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x74, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2c, 0x0a,
	0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x01, 0x10,
	0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x32, 0xad, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x6b, 0x79, 0x74,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52,
	0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x52, 0x6f, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x09,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x34, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x6b,
	0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x11, 0x66, 0x69, 0x6c, 0x65,
	0x74, 0x72, 0x65, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kythe_proto_filetree_proto_rawDescOnce sync.Once
	file_kythe_proto_filetree_proto_rawDescData = file_kythe_proto_filetree_proto_rawDesc
)

func file_kythe_proto_filetree_proto_rawDescGZIP() []byte {
	file_kythe_proto_filetree_proto_rawDescOnce.Do(func() {
		file_kythe_proto_filetree_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_filetree_proto_rawDescData)
	})
	return file_kythe_proto_filetree_proto_rawDescData
}

var file_kythe_proto_filetree_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kythe_proto_filetree_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kythe_proto_filetree_proto_goTypes = []interface{}{
	(DirectoryReply_Kind)(0),        // 0: kythe.proto.DirectoryReply.Kind
	(*CorpusRootsRequest)(nil),      // 1: kythe.proto.CorpusRootsRequest
	(*CorpusRootsReply)(nil),        // 2: kythe.proto.CorpusRootsReply
	(*DirectoryRequest)(nil),        // 3: kythe.proto.DirectoryRequest
	(*DirectoryReply)(nil),          // 4: kythe.proto.DirectoryReply
	(*CorpusRootsReply_Corpus)(nil), // 5: kythe.proto.CorpusRootsReply.Corpus
	(*DirectoryReply_Entry)(nil),    // 6: kythe.proto.DirectoryReply.Entry
}
var file_kythe_proto_filetree_proto_depIdxs = []int32{
	5, // 0: kythe.proto.CorpusRootsReply.corpus:type_name -> kythe.proto.CorpusRootsReply.Corpus
	6, // 1: kythe.proto.DirectoryReply.entry:type_name -> kythe.proto.DirectoryReply.Entry
	0, // 2: kythe.proto.DirectoryReply.Entry.kind:type_name -> kythe.proto.DirectoryReply.Kind
	1, // 3: kythe.proto.FileTreeService.CorpusRoots:input_type -> kythe.proto.CorpusRootsRequest
	3, // 4: kythe.proto.FileTreeService.Directory:input_type -> kythe.proto.DirectoryRequest
	2, // 5: kythe.proto.FileTreeService.CorpusRoots:output_type -> kythe.proto.CorpusRootsReply
	4, // 6: kythe.proto.FileTreeService.Directory:output_type -> kythe.proto.DirectoryReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kythe_proto_filetree_proto_init() }
func file_kythe_proto_filetree_proto_init() {
	if File_kythe_proto_filetree_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_filetree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorpusRootsRequest); i {
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
		file_kythe_proto_filetree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorpusRootsReply); i {
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
		file_kythe_proto_filetree_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryRequest); i {
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
		file_kythe_proto_filetree_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryReply); i {
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
		file_kythe_proto_filetree_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorpusRootsReply_Corpus); i {
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
		file_kythe_proto_filetree_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryReply_Entry); i {
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
			RawDescriptor: file_kythe_proto_filetree_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kythe_proto_filetree_proto_goTypes,
		DependencyIndexes: file_kythe_proto_filetree_proto_depIdxs,
		EnumInfos:         file_kythe_proto_filetree_proto_enumTypes,
		MessageInfos:      file_kythe_proto_filetree_proto_msgTypes,
	}.Build()
	File_kythe_proto_filetree_proto = out.File
	file_kythe_proto_filetree_proto_rawDesc = nil
	file_kythe_proto_filetree_proto_goTypes = nil
	file_kythe_proto_filetree_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileTreeServiceClient is the client API for FileTreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileTreeServiceClient interface {
	CorpusRoots(ctx context.Context, in *CorpusRootsRequest, opts ...grpc.CallOption) (*CorpusRootsReply, error)
	Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryReply, error)
}

type fileTreeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTreeServiceClient(cc grpc.ClientConnInterface) FileTreeServiceClient {
	return &fileTreeServiceClient{cc}
}

func (c *fileTreeServiceClient) CorpusRoots(ctx context.Context, in *CorpusRootsRequest, opts ...grpc.CallOption) (*CorpusRootsReply, error) {
	out := new(CorpusRootsReply)
	err := c.cc.Invoke(ctx, "/kythe.proto.FileTreeService/CorpusRoots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTreeServiceClient) Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryReply, error) {
	out := new(DirectoryReply)
	err := c.cc.Invoke(ctx, "/kythe.proto.FileTreeService/Directory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTreeServiceServer is the server API for FileTreeService service.
type FileTreeServiceServer interface {
	CorpusRoots(context.Context, *CorpusRootsRequest) (*CorpusRootsReply, error)
	Directory(context.Context, *DirectoryRequest) (*DirectoryReply, error)
}

// UnimplementedFileTreeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileTreeServiceServer struct {
}

func (*UnimplementedFileTreeServiceServer) CorpusRoots(context.Context, *CorpusRootsRequest) (*CorpusRootsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CorpusRoots not implemented")
}
func (*UnimplementedFileTreeServiceServer) Directory(context.Context, *DirectoryRequest) (*DirectoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Directory not implemented")
}

func RegisterFileTreeServiceServer(s *grpc.Server, srv FileTreeServiceServer) {
	s.RegisterService(&_FileTreeService_serviceDesc, srv)
}

func _FileTreeService_CorpusRoots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CorpusRootsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTreeServiceServer).CorpusRoots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.FileTreeService/CorpusRoots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTreeServiceServer).CorpusRoots(ctx, req.(*CorpusRootsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTreeService_Directory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTreeServiceServer).Directory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.FileTreeService/Directory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTreeServiceServer).Directory(ctx, req.(*DirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileTreeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.FileTreeService",
	HandlerType: (*FileTreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CorpusRoots",
			Handler:    _FileTreeService_CorpusRoots_Handler,
		},
		{
			MethodName: "Directory",
			Handler:    _FileTreeService_Directory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kythe/proto/filetree.proto",
}
