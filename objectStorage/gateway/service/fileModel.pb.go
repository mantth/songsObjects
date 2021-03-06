// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.0
// source: fileModel.proto

package service

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

type FileModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	HashCode string `protobuf:"bytes,2,opt,name=HashCode,proto3" json:"HashCode,omitempty"`
	FileSize uint64 `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	IsExist  bool   `protobuf:"varint,4,opt,name=IsExist,proto3" json:"IsExist,omitempty"`
	ShardID  uint32 `protobuf:"varint,5,opt,name=ShardID,proto3" json:"ShardID,omitempty"`
}

func (x *FileModel) Reset() {
	*x = FileModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileModel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileModel) ProtoMessage() {}

func (x *FileModel) ProtoReflect() protoreflect.Message {
	mi := &file_fileModel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileModel.ProtoReflect.Descriptor instead.
func (*FileModel) Descriptor() ([]byte, []int) {
	return file_fileModel_proto_rawDescGZIP(), []int{0}
}

func (x *FileModel) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileModel) GetHashCode() string {
	if x != nil {
		return x.HashCode
	}
	return ""
}

func (x *FileModel) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileModel) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

func (x *FileModel) GetShardID() uint32 {
	if x != nil {
		return x.ShardID
	}
	return 0
}

var File_fileModel_proto protoreflect.FileDescriptor

var file_fileModel_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49,
	0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileModel_proto_rawDescOnce sync.Once
	file_fileModel_proto_rawDescData = file_fileModel_proto_rawDesc
)

func file_fileModel_proto_rawDescGZIP() []byte {
	file_fileModel_proto_rawDescOnce.Do(func() {
		file_fileModel_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileModel_proto_rawDescData)
	})
	return file_fileModel_proto_rawDescData
}

var file_fileModel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fileModel_proto_goTypes = []interface{}{
	(*FileModel)(nil), // 0: service.FileModel
}
var file_fileModel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fileModel_proto_init() }
func file_fileModel_proto_init() {
	if File_fileModel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileModel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileModel); i {
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
			RawDescriptor: file_fileModel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fileModel_proto_goTypes,
		DependencyIndexes: file_fileModel_proto_depIdxs,
		MessageInfos:      file_fileModel_proto_msgTypes,
	}.Build()
	File_fileModel_proto = out.File
	file_fileModel_proto_rawDesc = nil
	file_fileModel_proto_goTypes = nil
	file_fileModel_proto_depIdxs = nil
}
