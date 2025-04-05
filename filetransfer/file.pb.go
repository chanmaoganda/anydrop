// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: file.proto

package filetransfer

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

type FileChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []byte                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ChunkIndex    int32                  `protobuf:"varint,2,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
	Planned       int32                  `protobuf:"varint,3,opt,name=planned,proto3" json:"planned,omitempty"`
	FileHash      string                 `protobuf:"bytes,4,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	FileName      string                 `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	mi := &file_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *FileChunk) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

func (x *FileChunk) GetPlanned() int32 {
	if x != nil {
		return x.Planned
	}
	return 0
}

func (x *FileChunk) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *FileChunk) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadStatus struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Success        bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message        string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ReceivedChunks int32                  `protobuf:"varint,3,opt,name=received_chunks,json=receivedChunks,proto3" json:"received_chunks,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	mi := &file_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UploadStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadStatus) GetReceivedChunks() int32 {
	if x != nil {
		return x.ReceivedChunks
	}
	return 0
}

type FileMeta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileHash      string                 `protobuf:"bytes,2,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	PlannedChunks int32                  `protobuf:"varint,3,opt,name=planned_chunks,json=plannedChunks,proto3" json:"planned_chunks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	mi := &file_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *FileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMeta) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *FileMeta) GetPlannedChunks() int32 {
	if x != nil {
		return x.PlannedChunks
	}
	return 0
}

type UploadPlan struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileHash      string                 `protobuf:"bytes,1,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	NeededChunks  []int32                `protobuf:"varint,2,rep,packed,name=needed_chunks,json=neededChunks,proto3" json:"needed_chunks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadPlan) Reset() {
	*x = UploadPlan{}
	mi := &file_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPlan) ProtoMessage() {}

func (x *UploadPlan) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPlan.ProtoReflect.Descriptor instead.
func (*UploadPlan) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *UploadPlan) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *UploadPlan) GetNeededChunks() []int32 {
	if x != nil {
		return x.NeededChunks
	}
	return nil
}

var File_file_proto protoreflect.FileDescriptor

const file_file_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"file.proto\x12\ffiletransfer\"\x9a\x01\n" +
	"\tFileChunk\x12\x18\n" +
	"\acontent\x18\x01 \x01(\fR\acontent\x12\x1f\n" +
	"\vchunk_index\x18\x02 \x01(\x05R\n" +
	"chunkIndex\x12\x18\n" +
	"\aplanned\x18\x03 \x01(\x05R\aplanned\x12\x1b\n" +
	"\tfile_hash\x18\x04 \x01(\tR\bfileHash\x12\x1b\n" +
	"\tfile_name\x18\x05 \x01(\tR\bfileName\"k\n" +
	"\fUploadStatus\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12'\n" +
	"\x0freceived_chunks\x18\x03 \x01(\x05R\x0ereceivedChunks\"k\n" +
	"\bFileMeta\x12\x1b\n" +
	"\tfile_name\x18\x01 \x01(\tR\bfileName\x12\x1b\n" +
	"\tfile_hash\x18\x02 \x01(\tR\bfileHash\x12%\n" +
	"\x0eplanned_chunks\x18\x03 \x01(\x05R\rplannedChunks\"N\n" +
	"\n" +
	"UploadPlan\x12\x1b\n" +
	"\tfile_hash\x18\x01 \x01(\tR\bfileHash\x12#\n" +
	"\rneeded_chunks\x18\x02 \x03(\x05R\fneededChunks2\xcd\x01\n" +
	"\vFileService\x12=\n" +
	"\tQueryPlan\x12\x16.filetransfer.FileMeta\x1a\x18.filetransfer.UploadPlan\x12?\n" +
	"\x06Upload\x12\x17.filetransfer.FileChunk\x1a\x1a.filetransfer.UploadStatus(\x01\x12>\n" +
	"\bMakeFile\x12\x16.filetransfer.FileMeta\x1a\x1a.filetransfer.UploadStatusB\x11Z\x0f../filetransferb\x06proto3"

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData []byte
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_file_proto_rawDesc), len(file_file_proto_rawDesc)))
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_proto_goTypes = []any{
	(*FileChunk)(nil),    // 0: filetransfer.FileChunk
	(*UploadStatus)(nil), // 1: filetransfer.UploadStatus
	(*FileMeta)(nil),     // 2: filetransfer.FileMeta
	(*UploadPlan)(nil),   // 3: filetransfer.UploadPlan
}
var file_file_proto_depIdxs = []int32{
	2, // 0: filetransfer.FileService.QueryPlan:input_type -> filetransfer.FileMeta
	0, // 1: filetransfer.FileService.Upload:input_type -> filetransfer.FileChunk
	2, // 2: filetransfer.FileService.MakeFile:input_type -> filetransfer.FileMeta
	3, // 3: filetransfer.FileService.QueryPlan:output_type -> filetransfer.UploadPlan
	1, // 4: filetransfer.FileService.Upload:output_type -> filetransfer.UploadStatus
	1, // 5: filetransfer.FileService.MakeFile:output_type -> filetransfer.UploadStatus
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_proto_rawDesc), len(file_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
