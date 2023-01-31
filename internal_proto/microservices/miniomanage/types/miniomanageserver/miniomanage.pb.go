// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/miniomanage.proto

package miniomanageserver

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

type PutFileUploaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectPre  string `protobuf:"bytes,2,opt,name=objectPre,proto3" json:"objectPre,omitempty"`
	FilePath   string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *PutFileUploaderRequest) Reset() {
	*x = PutFileUploaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderRequest) ProtoMessage() {}

func (x *PutFileUploaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderRequest.ProtoReflect.Descriptor instead.
func (*PutFileUploaderRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{0}
}

func (x *PutFileUploaderRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutFileUploaderRequest) GetObjectPre() string {
	if x != nil {
		return x.ObjectPre
	}
	return ""
}

func (x *PutFileUploaderRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type PutFileUploaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 用户性别
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *PutFileUploaderResponse) Reset() {
	*x = PutFileUploaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderResponse) ProtoMessage() {}

func (x *PutFileUploaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderResponse.ProtoReflect.Descriptor instead.
func (*PutFileUploaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{1}
}

func (x *PutFileUploaderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutFileUploaderResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutFileUploaderResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type PutFileUploaderByteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectPre  string `protobuf:"bytes,2,opt,name=objectPre,proto3" json:"objectPre,omitempty"`
	FilePath   string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *PutFileUploaderByteRequest) Reset() {
	*x = PutFileUploaderByteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderByteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderByteRequest) ProtoMessage() {}

func (x *PutFileUploaderByteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderByteRequest.ProtoReflect.Descriptor instead.
func (*PutFileUploaderByteRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{2}
}

func (x *PutFileUploaderByteRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *PutFileUploaderByteRequest) GetObjectPre() string {
	if x != nil {
		return x.ObjectPre
	}
	return ""
}

func (x *PutFileUploaderByteRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type PutFileUploaderByteponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 用户性别
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *PutFileUploaderByteponse) Reset() {
	*x = PutFileUploaderByteponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileUploaderByteponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileUploaderByteponse) ProtoMessage() {}

func (x *PutFileUploaderByteponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileUploaderByteponse.ProtoReflect.Descriptor instead.
func (*PutFileUploaderByteponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{3}
}

func (x *PutFileUploaderByteponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutFileUploaderByteponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutFileUploaderByteponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type GetMinioConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 用户性别
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *GetMinioConnectRequest) Reset() {
	*x = GetMinioConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinioConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinioConnectRequest) ProtoMessage() {}

func (x *GetMinioConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinioConnectRequest.ProtoReflect.Descriptor instead.
func (*GetMinioConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{4}
}

func (x *GetMinioConnectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMinioConnectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMinioConnectRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type GetMinioConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	ObjectPre  string `protobuf:"bytes,2,opt,name=objectPre,proto3" json:"objectPre,omitempty"`
	FilePath   string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *GetMinioConnectResponse) Reset() {
	*x = GetMinioConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miniomanage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinioConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinioConnectResponse) ProtoMessage() {}

func (x *GetMinioConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miniomanage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinioConnectResponse.ProtoReflect.Descriptor instead.
func (*GetMinioConnectResponse) Descriptor() ([]byte, []int) {
	return file_proto_miniomanage_proto_rawDescGZIP(), []int{5}
}

func (x *GetMinioConnectResponse) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *GetMinioConnectResponse) GetObjectPre() string {
	if x != nil {
		return x.ObjectPre
	}
	return ""
}

func (x *GetMinioConnectResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

var File_proto_miniomanage_proto protoreflect.FileDescriptor

var file_proto_miniomanage_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x16,
	0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x55, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x1a, 0x50, 0x75, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x56, 0x0a, 0x18, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x69,
	0x6e, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x73, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x32, 0xda, 0x02, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x71, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x69, 0x6e, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_miniomanage_proto_rawDescOnce sync.Once
	file_proto_miniomanage_proto_rawDescData = file_proto_miniomanage_proto_rawDesc
)

func file_proto_miniomanage_proto_rawDescGZIP() []byte {
	file_proto_miniomanage_proto_rawDescOnce.Do(func() {
		file_proto_miniomanage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_miniomanage_proto_rawDescData)
	})
	return file_proto_miniomanage_proto_rawDescData
}

var file_proto_miniomanage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_miniomanage_proto_goTypes = []interface{}{
	(*PutFileUploaderRequest)(nil),     // 0: miniomanageserver.PutFileUploaderRequest
	(*PutFileUploaderResponse)(nil),    // 1: miniomanageserver.PutFileUploaderResponse
	(*PutFileUploaderByteRequest)(nil), // 2: miniomanageserver.PutFileUploaderByteRequest
	(*PutFileUploaderByteponse)(nil),   // 3: miniomanageserver.PutFileUploaderByteponse
	(*GetMinioConnectRequest)(nil),     // 4: miniomanageserver.GetMinioConnectRequest
	(*GetMinioConnectResponse)(nil),    // 5: miniomanageserver.GetMinioConnectResponse
}
var file_proto_miniomanage_proto_depIdxs = []int32{
	0, // 0: miniomanageserver.MinioManageServer.PutFileUploader:input_type -> miniomanageserver.PutFileUploaderRequest
	2, // 1: miniomanageserver.MinioManageServer.PutFileUploaderByte:input_type -> miniomanageserver.PutFileUploaderByteRequest
	4, // 2: miniomanageserver.MinioManageServer.GetFileUploader:input_type -> miniomanageserver.GetMinioConnectRequest
	1, // 3: miniomanageserver.MinioManageServer.PutFileUploader:output_type -> miniomanageserver.PutFileUploaderResponse
	3, // 4: miniomanageserver.MinioManageServer.PutFileUploaderByte:output_type -> miniomanageserver.PutFileUploaderByteponse
	5, // 5: miniomanageserver.MinioManageServer.GetFileUploader:output_type -> miniomanageserver.GetMinioConnectResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_miniomanage_proto_init() }
func file_proto_miniomanage_proto_init() {
	if File_proto_miniomanage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_miniomanage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderResponse); i {
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
		file_proto_miniomanage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderByteRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileUploaderByteponse); i {
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
		file_proto_miniomanage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinioConnectRequest); i {
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
		file_proto_miniomanage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinioConnectResponse); i {
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
			RawDescriptor: file_proto_miniomanage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_miniomanage_proto_goTypes,
		DependencyIndexes: file_proto_miniomanage_proto_depIdxs,
		MessageInfos:      file_proto_miniomanage_proto_msgTypes,
	}.Build()
	File_proto_miniomanage_proto = out.File
	file_proto_miniomanage_proto_rawDesc = nil
	file_proto_miniomanage_proto_goTypes = nil
	file_proto_miniomanage_proto_depIdxs = nil
}
